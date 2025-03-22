package users

import (
	"fmt"
	"net/http"

	"github.com/PabloPei/SmartSpend-backend/internal/errors"
	"github.com/PabloPei/SmartSpend-backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	// User routes
	router.HandleFunc("/user/register", h.handleUserRegister).Methods("POST")
	router.HandleFunc("/user/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/user/photo/{email}", h.handleUserPhoto).Methods("POST", "PUT")

	// Admin Routes
	router.HandleFunc("/user/{email}", h.handleUser).Methods("GET")
}

func (h *Handler) handleUserRegister(w http.ResponseWriter, r *http.Request) {

	var user RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user: %v", errors))
		return
	}

	err := h.service.RegisterUser(user)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	var user LogInUserPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	token, err := h.service.LogInUser(user)
	if err == errors.ErrInvalidCredentials || err == errors.ErrUserNotFound(user.Email) {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	} else if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) handleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, ok := vars["email"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing email"))
		return
	}

	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) handleUserPhoto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	email, ok := vars["email"]

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing email"))
		return
	}

	var payload UploadPhotoPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	err := h.service.UploadPhoto(payload, email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Photo uploaded successfully",
	})
}
