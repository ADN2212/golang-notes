package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	//   "PlanetsAPI/Internal/funcs"
)

type planet struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Vol string `json:"vol"`
	Mass string `json:"mass"`
}

func main(){
	//fmt.Println("Running...")
	router := gin.Default()
	router.GET("/planets", getAllPlanets)
	router.GET("/planets/:id", getPlanetByID)
	router.POST("/addPlanet", addPlanet)
	router.PATCH("/upDatePlanet/:id", upDatePlanet)
	//router.PATCH("/return", returnBook)
	funcs.Foo()
	router.Run("localhost:8080")

}

func upDatePlanet (context *gin.Context) {

	id := context.Param("id")

	//Es porbable que esta primera condicion no ocurra, porque si no hay id es como si el endpoint fuera otro.
	if id == "" {
		context.JSON(
				http.StatusBadRequest, 
				gin.H{"message": fmt.Sprint("ID is a required file.")})
		return
	} else {
		//Buscar el planeta que se pretende actualizar:
		planetPointer, err := findPlanetByID(id)
		if err != nil {
			context.JSON(
				http.StatusBadRequest, 
				gin.H{"message": fmt.Sprint("There is not planet whit ID = ", id, ".")})
		return

		}
		//Dado que findPlanet retorna el puntero del planeta en el array,  
		//prodremos actualizarlo de manera directa:
		var newDataPlanet planet
		uptadeErr := context.BindJSON(&newDataPlanet)
		//fmt.Println(uptadeErr)
		if uptadeErr == nil {
			//fmt.Println(newDataPlanet)
			//fmt.Println(planetPointer)
			//Notece como no es nesesario actualizar el ID:
			(*planetPointer).Name = newDataPlanet.Name
			(*planetPointer).Desc = newDataPlanet.Desc
			(*planetPointer).Vol = newDataPlanet.Vol
			(*planetPointer).Mass = newDataPlanet.Mass
			//*planetPointer = newDataPlanet//Asignamos la nueva data al planeta.
			context.JSON(
				http.StatusResetContent,
				gin.H{"message": fmt.Sprint("Planet whit id = ", id, " has been updated.")})	
			return
		} else {  
			fmt.Println("Bobo!!")
				context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Something went wrong : (."})	
		}
	}
}

//fmt.Sprint("Planet whit id = ", id, " has been updated.")
//gin.H{"message": fmt.Sprint("Planet whit id = ", id, " not found")})


//Notece como el arg es la struct context de gin, ver https://pkg.go.dev/github.com/gin-gonic/gin#Context
func getAllPlanets(context *gin.Context) {
	//El metodo JSON transforma las structs de tipo planet en JSONs y los envia
	//como response de la peticion de la API
	//el primer argumento es el estado de la response.
	context.JSON(http.StatusOK, planets)

}

//Ojo: para que una func pueda ser callback de un endpoint
//debe ser metodo del objeto context. 
func getPlanetByID(context *gin.Context){

	id := context.Param("id")
	p, err := findPlanetByID(id) 

	if err != nil {
		context.JSON(
			http.StatusNotFound, 
			gin.H{"message": fmt.Sprint("Planet whit id = ", id, " not found")})
		return
	}

	//p, err := findPlanetByID(context.Param("id"))

	//fmt.Println(p, err)

	context.JSON(http.StatusOK, p)
}

func findPlanetByID(id string) (*planet, error) {

	for i, planet := range planets {
		if planet.ID == id{
			//Retornar el puntero al planeta y el error como nulo.
			return &planets[i], nil
		}
	}

	return nil, errors.New("Planet not found")
}


func addPlanet(context *gin.Context) {
	var newPlanet planet//Una struct del tipo planet.
	fmt.Printf("Actual new planet %v.\n", newPlanet)

	//Toma el puntero de la variable como arg y pone el JSON trasnformado en struct
	err := context.BindJSON(&newPlanet)	
	if err != nil {
			context.JSON(
				http.StatusBadRequest, 
				gin.H{"message": fmt.Sprint(" Something went wrong :( ")})
		return
	}

	newId := newPlanet.ID; 

	//Si el JSON que se envio no tiene ID habra un str vacio en la struct:
	//Este proceso se puede repetir pata todos los campos requeridos, so it can be another function.  
	if newId == "" {
		context.JSON(
				http.StatusBadRequest, 
				gin.H{"message": fmt.Sprint("ID is a required file.")})
		return
	} else {
		//Si hay ID provar que no sea repetido:
		for _, p := range planets {
			if p.ID == newId {
				context.JSON(
					http.StatusBadRequest, 
					gin.H{"message": fmt.Sprint("A planet whit this ID already exits.")})
				return
			}
		}
	}

	planets = append(planets, newPlanet)
	//fmt.Printf("New planet actualizado a %v.\n", newPlanet) 
	context.IndentedJSON(
				http.StatusCreated, 
				gin.H{"message": fmt.Sprint("Planet Created.")})
}
