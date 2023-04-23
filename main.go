package main

import (
	"log"
	"os"

	"github.com/Fadiouf/goprojet/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initialize routes
func initRoutes(e *echo.Echo, db *gorm.DB, jwtSecret []byte) {
	// Initialize handlers
	userHandler := handlers.UserHandler{DB: db}
	roleHandler := handlers.RoleHandler{DB: db}
	groupHandler := handlers.GroupHandler{DB: db}
	authHandler := handlers.AuthHandler{DB: db, JWTSecret: jwtSecret}

	// Initialize middleware
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtSecret,
	})

	// Users
	e.GET("/users", userHandler.GetUsers, jwtMiddleware)
	e.POST("/users", userHandler.CreateUser, jwtMiddleware)
	e.PUT("/users/:id", userHandler.UpdateUser, jwtMiddleware)
	e.DELETE("/users/:id", userHandler.DeleteUser, jwtMiddleware)

	// Roles
	e.GET("/roles", roleHandler.GetRoles, jwtMiddleware)
	e.POST("/roles", roleHandler.CreateRole, jwtMiddleware)
	e.PUT("/roles/:id", roleHandler.UpdateRole, jwtMiddleware)
	e.DELETE("/roles/:id", roleHandler.DeleteRole, jwtMiddleware)

	// Groups
	e.GET("/groups", groupHandler.GetGroups, jwtMiddleware)
	e.POST("/groups", groupHandler.CreateGroup, jwtMiddleware)
	e.PUT("/groups/:id", groupHandler.UpdateGroup, jwtMiddleware)
	e.DELETE("/groups/:id", groupHandler.DeleteGroup, jwtMiddleware)

	// Auth
	e.POST("/auth", authHandler.Authenticate)
}

// Main function
func main() {
	// Open a connection to the database
	db_connexion, err := gorm.Open(postgres.Open("host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create an instance of the Echo framework
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/auth", handleAuth)

	users := e.Group("/user")
	users.GET("", handleGetUsers)
	users.POST("", handleCreateUser)
	users.PUT("/:id", handleUpdateUser)
	users.DELETE("/:id", handleDeleteUser)

	roles := e.Group("/roles")
	roles.GET("", handleGetRoles)
	roles.POST("", handleCreateRole)
	roles.PUT("/:id", handleUpdateRole)
	roles.DELETE("/:id", handleDeleteRole)

	groups := e.Group("/groups")
	groups.GET("", handleGetGroups)
	groups.POST("", handleCreateGroup)
	groups.PUT("/:id", handleUpdateGroup)
	groups.DELETE("/:id", handleDeleteGroup)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))

	// Initialize environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	// db := database.InitDB()
	// defer db.Close()

	// Initialize JWT secret
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Initialize routes
	initRoutes(e, db_connexion, jwtSecret)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
