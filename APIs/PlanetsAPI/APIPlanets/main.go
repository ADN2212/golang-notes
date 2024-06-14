package main

import (
	"fmt"
	"example.com/functions"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main(){
	defer fmt.Println("Fin")
	router := gin.Default()
	//Planets Endpoints:
	router.GET("/planets", functions.GetAllPlanets)
	router.GET("/planets/:id", functions.GetPlanetByID)
	router.POST("/addPlanet", functions.AddPlanet)
	router.PATCH("/upDatePlanet/:id", functions.UpDatePlanet)
	
	//Stars Endpoints:
	router.GET("/stars", functions.GetAllStars)
	router.POST("/addStar", functions.AddStar)
	router.GET("/stars/:id", functions.GetStarByID)
	router.POST("/addPlanetToStar/:id_star/:id_planet", functions.AddPlanetToStar)
	router.GET("/getPlanetsArroundStar/:id_star", functions.GetPlanetsArroundStar)

	//Asinar un puerto a la API.
	router.Run("localhost:8080")
}
