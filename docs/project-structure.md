# Project Structure 

/SmartSpend
│── /cmd               # Punto de entrada de la aplicación (main.go)
│── /config            # Archivos de configuración (variables de entorno, etc.)
│── /internal          # Código interno del backend (no accesible desde otros módulos)
│   │── /handlers      # Controladores HTTP (manejan las peticiones)
│   │── /services      # Lógica de negocio
│   │── /repositories  # Acceso a la base de datos
│   │── /models        # Definiciones de estructuras de datos
│   │── /middlewares   # Middlewares (autenticación, logging, etc.)
│── /pkg               # Código reutilizable (puede ser usado por otros proyectos)
│── /db                # Migraciones y scripts para la base de datos
│── /scripts           # Scripts útiles (ej. inicializar datos)
│── /test              # Pruebas unitarias e integración
│── .env               # Variables de entorno (no subir a git)
│── go.mod             # Archivo de módulos de Go
│── go.sum             # Checksums de dependencias
│── Makefile           # Comandos de automatización
│── README.md          # Documentación del proyecto
