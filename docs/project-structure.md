# Project Structure 

/SmartSpend
│── /cmd # Entry point of the application (main.go)
│── /config # Configuration files (environment variables, etc.)
│── /internal # Internal backend code (not accessible from other modules)
│ │── /handlers # HTTP controllers (handle requests)
│ │── /services # Business logic
│ │── /repositories # Database access
│ │── /models # Data structure definitions
│ │── /middlewares # Middlewares (authentication, logging, etc.)
│── /pkg # Reusable code (can be used by other projects)
│── /db # Database migrations and scripts
│── /scripts # Useful scripts (e.g., initialize data)
│── /test # Unit and integration tests
│── .env # Environment variables (do not upload to git)
│── go.mod # Go module file
│── go.sum # Dependency checksums
│── Makefile # Automation commands
│── README.md # Project documentation