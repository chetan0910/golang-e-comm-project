package main

import (
	"log"
	"os"

	"github.com/chetan0910/golang-e-comm-project/controllers"
	"github.com/chetan0910/golang-e-comm-project/database"
	"github.com/chetan0910/golang-e-comm-project/middleware"
	"github.com/chetan0910/golang-e-comm-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT") //os is a built-in Go package that lets you interact with the operating system.
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(
		database.ProductData(database.client, "Products"),
		database.UserData(database.clent, "Users"),
	)

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication()) //for checking the authentication using middleware

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
