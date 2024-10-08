package main

import (

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/kCal0/database"

	"github.com/eu-micaeu/kCal0/middlewares"

	"github.com/eu-micaeu/kCal0/routes"

)

func main() {

	r := gin.Default() // Creates a gin route handle

	r.Use(middlewares.CorsMiddleware())        // Middleware CORS

    r.Use(middlewares.CacheCleanerMiddleware()) // Middleware de limpeza de cache

	db, err := database.NewDB() // Connects to the database

	if err != nil { // If there is an error connecting to the database, the program will stop

		panic(err) 

	}

	routes.UserRoutes(r, db) // Calls the UserRoutes function and passes the route handle and the database connection

	routes.MenuRoutes(r, db) // Calls the MenuRoutes function and passes the route handle and the database connection

	routes.MealRoutes(r, db) // Calls the MealRoutes function and passes the route handle and the database connection

	routes.FoodRoutes(r, db) // Calls the FoodRoutes function and passes the route handle and the database connection

	r.LoadHTMLGlob("./views/*.html") // Load the HTML files

	r.GET("/", func(c *gin.Context) { // When accessing the root route, the index.html file will be rendered

		c.HTML(http.StatusOK, "index.html", nil)

	})

	r.GET("/register", func(c *gin.Context) { // When accessing the /register route, the register.html file will be rendered

		c.HTML(http.StatusOK, "register.html", nil)

	})

	r.GET("/home", func(c *gin.Context) { // When accessing the /home route, the home.html file will be rendered

		c.HTML(http.StatusOK, "home.html", nil)

	})

	r.GET("/perfil", func(c *gin.Context) { // When accessing the /perfil route, the perfil.html file will be rendered

		c.HTML(http.StatusOK, "perfil.html", nil)

	})

	r.Static("/static", "./static") // Load the static files

	r.Run() // Run the server

}