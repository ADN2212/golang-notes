package functions


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	//"reflect"
)

// func Foo(){
// 	fmt.Println("...Working...")
// }


type star struct {
	ID string `json:"id_star"`
	StarName string `json:"star_name"`
}

type planet struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
	Vol string `json:"vol"`
	Mass string `json:"mass"`
	IDStar string `json:"id_star"`//el id de la strella a la que pertenece este planeta.
}

//Mas adelante esta info debera salir de una BD:

var stars = []star{
	{
		ID: "1",
		StarName: "Sol",
	},
}


var planets = []planet{
	{
		ID: "1", 
		Name: "Mercury",
		Desc: "Mercury is the smallest planet in the Solar System and the closest to the Sun. Its orbit around the Sun takes 87.97 Earth days, the shortest of all the Sun's planets.",
		Vol: "6.083 x 10^10 km^3",
		Mass: "3.3011 x 10^23 kg",
	},
	{
		ID: "2", 
		Name: "Venus",
		Desc: "Venus is the second planet from the Sun and is named after the Roman goddess of love and beauty. As the brightest natural object in Earth's night sky after the Moon, Venus can cast shadows and can be visible to the naked eye in broad daylight.",
		Vol: "9.2843 x 10^11 km^3",
		Mass: "4.8675 x 10^24 kg",
	},
	{
		ID: "3",
		Name: "Earth",
		Desc: "Earth is the third planet from the Sun and the only astronomical object known to harbor life. While large amounts of water can be found throughout the Solar System, only Earth sustains liquid surface water. About 71% of Earth's surface is made up of the ocean, dwarfing Earth's polar ice, lakes, and rivers.",
		Vol: "1.08321 x 1012 km^3",
		Mass: "5.97237 x 1024 kg",
	},
	{
		ID: "4",
		Name: "Mars",
		Desc: "Mars is the fourth planet from the Sun and the second-smallest planet in the Solar System, being larger than only Mercury. In English, Mars carries the name of the Roman god of war and is often called the \\\"Red Planet\\\".\"",
		Vol: "1.63118 x 10^11 km^3",
		Mass: "6.4171 x 10^23 kg",
	},
	{
		ID: "5",
		Name: "Jupiter",
		Desc: "Jupiter is the fifth planet from the Sun and the largest in the Solar System. It is a gas giant with a mass more than two and a half times that of all the other planets in the Solar System combined, but slightly less than one-thousandth the mass of the Sun.",
		Vol: "1.4313 x 10^15 km^3",
		Mass: "1.8982 x 10^27 kg",
	},
	{
		ID: "6",
		Name: "Saturn",
		Desc: "Saturn is the sixth planet from the Sun and the second-largest in the Solar System, after Jupiter. It is a gas giant with an average radius of about nine and a half times that of Earth. It has only one-eighth the average density of Earth; however, with its larger volume, Saturn is over 95 times more massive.",
		Vol: "8.2713 x 10^14 km^3",
		Mass: "8.2713 x 10^14 km^3",
	},
	{
		ID: "7",
		Name: "Uranus",
		Desc: "Uranus is the seventh planet from the Sun. Its name is a reference to the Greek god of the sky, Uranus, who, according to Greek mythology, was the great-grandfather of Ares, grandfather of Zeus and father of Cronus. It has the third-largest planetary radius and fourth-largest planetary mass in the Solar System.",
		Vol: "6.833 x 1013 km^3",
		Mass: "(8.6810±0.0013) x 1025 kg",
	},
	{
		ID: "8",
		Name: "Neptune",
		Desc: "Neptune is the eighth and farthest-known Solar planet from the Sun. In the Solar System, it is the fourth-largest planet by diameter, the third-most-massive planet, and the densest giant planet. It is 17 times the mass of Earth, and slightly more massive than its near-twin Uranus.",
		Vol: "6.253 x 1013 km^3",
		Mass: "1.02413 x 1026 kg",
	},
}


//Planets Endpoints:

//Notece como el arg es la struct context de gin, ver https://pkg.go.dev/github.com/gin-gonic/gin#Context
func GetAllPlanets(context *gin.Context) {
	//El metodo JSON transforma las structs de tipo planet en JSONs y los envia
	//como response de la peticion de la API
	//el primer argumento es el estado de la response.
	context.JSON(http.StatusOK, planets)
}

//Ojo: para que una func pueda ser callback de un endpoint
//debe ser metodo del objeto context. 
func GetPlanetByID(context *gin.Context) {

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

//Notece como las funciones que son exportables empiezan con una letra mayuscula.
func AddPlanet(context *gin.Context) {
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


func UpDatePlanet(context *gin.Context) {

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
		//podremos actualizarlo de manera directa:
		var newDataPlanet planet
		uptadeErr := context.BindJSON(&newDataPlanet)
		//fmt.Println(uptadeErr)
		if uptadeErr == nil {
			//fmt.Println(newDataPlanet)
			//fmt.Println(planetPointer)
			//Notece como no es nesesario actualizar el ID:
			// (*planetPointer).Name = newDataPlanet.Name
			// (*planetPointer).Desc = newDataPlanet.Desc
			// (*planetPointer).Vol = newDataPlanet.Vol
			// (*planetPointer).Mass = newDataPlanet.Mass

			upDateFieldIfNotBlank(planetPointer, &newDataPlanet)
			
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

func upDateFieldIfNotBlank(planetPointer, newDataPlanetPointer *planet){
		//planetValues := reflect.valueOf(*newDataPlanetPointer)
		//keys.planetValues.Type()
		
		//Esto deberia hacerce de manera iterativa.

		if (*newDataPlanetPointer).Name != "" {
			(*planetPointer).Name = (*newDataPlanetPointer).Name
		}

		if (*newDataPlanetPointer).Desc != "" {
			(*planetPointer).Desc = (*newDataPlanetPointer).Desc
		}

		if (*newDataPlanetPointer).Vol != "" {
			(*planetPointer).Vol = (*newDataPlanetPointer).Vol
		}

		if (*newDataPlanetPointer).Mass != "" {
			(*planetPointer).Mass = (*newDataPlanetPointer).Mass
		}
}

//Stars endpoints:
func GetAllStars(context *gin.Context) {
	context.JSON(http.StatusOK, stars)
}

func findStarByID(id string) (*star, error) {
	for i, star := range stars {
		if star.ID == id {
			return &stars[i], nil
		}
	}
	return nil, errors.New("Star not found")
}

func GetStarByID(context *gin.Context) {

	id := context.Param("id")
	s, err := findStarByID(id)

	if err != nil {
		context.JSON(
			http.StatusNotFound, 
			gin.H{"message": fmt.Sprint("Star whit id = ", id, " not found")})
		return
	}

	context.JSON(http.StatusOK, s)
}

//Esta funcion se encarga de agregar un planeta a la orbita de una estrella.
func AddPlanetToStar(context *gin.Context){

	idStar := context.Param("id_star")
	idPlanet := context.Param("id_planet")

	//fmt.Printf("Adding planet id = %v to star id = %v.\n", idPlanet, idStar)

	starPointer, starErr := findStarByID(idStar)
	planetPointer, planetErr := findPlanetByID(idPlanet)
	
	if starErr != nil || planetErr != nil {
			context.JSON(
	 		http.StatusNotFound, 
	 		//gin.H{"message": fmt.Sprintf("One of them (planet of id = %v or star of id = %v) does not exist.", idPlanet, idStar)})
	 		gin.H{"message": fmt.Sprint("One of them (planet of id = ",  idPlanet," or star of id = ",idStar, ") does not exist.")})
	 	return
	}

	(*planetPointer).IDStar = idStar

	context.JSON(
		http.StatusOK, 
		gin.H{"message": fmt.Sprintf("done, the planet %v is now orbiting around the star %v.\n", (*planetPointer).Name, (*starPointer).StarName)})

}

//Esta funcion retorna un array de los planetas que orvitan al rededor de una strella.
func GetPlanetsArroundStar(context *gin.Context){

	idStar := context.Param("id_star")

	_, starErr := findStarByID(idStar)

	if starErr != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": fmt.Sprintf("there is no star whit id = %v.\n", idStar)})
		return
	}

	var planetsAroundStar []planet

	for _, p := range planets {
		if p.IDStar == idStar {
			planetsAroundStar = append(planetsAroundStar, p)
		}
	}

	//fmt.Println(planetsAroundStar)
	context.JSON(http.StatusOK, planetsAroundStar)
}

func AddStar(context *gin.Context){
	var newStar star
	err := context.BindJSON(&newStar)	
	if err != nil {
			context.JSON(
				http.StatusBadRequest, 
				gin.H{"message": fmt.Sprint(err)})
		return
	}

	//El proceso de comprobar el ID de la estrella no sera nesesario cuando se este usando una BD real.
	//newStarId := newStart.ID;

	stars = append(stars, newStar)

	context.JSON(http.StatusOK, fmt.Sprint("New star added successfully.")) 

}


