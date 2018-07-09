package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Person struct
type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	ID        string `json:"id"`
}

// People array for storing persons
type People []Person

// intitalize a peeps People (slice-type)
var peeps People

func main() {

	// create a literal peeps slice
	peeps = append(peeps, Person{Firstname: "colin", Lastname: "Bell", ID: "1"})

	// initalize Gin
	r := gin.Default()

	fmt.Println("gin type is: ", reflect.TypeOf(r))

	// load all html files into this package
	r.LoadHTMLGlob("./*.html")

	// serve the homepage
	r.GET("/", homepage)

	// serve the list of people
	r.GET("/peeps", getPeeps)

	// create new person
	r.POST("/person", createPerson)

	// update a person
	r.PUT("/person/:id", updatePerson)

	r.DELETE("/person/:id", removePerson)

	r.Run(":3000") // listen and serve on localhost:3000

}

// loads the homepage to the base url
func homepage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// serves up a list of all the people
func getPeeps(c *gin.Context) {
	c.JSON(200, gin.H{
		"peeps": peeps,
	})
}

// creates a new person
func createPerson(c *gin.Context) {
	newPerson := &Person{}
	newPerson.ID = strconv.Itoa(rand.Intn(999999))
	c.Bind(newPerson)
	peeps = append(peeps, *newPerson)
	c.JSON(http.StatusOK, peeps)
}

// updates a person
func updatePerson(c *gin.Context) {
	id := c.Param("id")
	index := -1
	// loop trough the peeps slice and try to find the person with the matching id
	for idx, peep := range peeps {
		if peep.ID == id {
			index = idx
		}
	}

	// check if the id exists - if not, return a string message
	if index == -1 {
		c.String(http.StatusOK, "Sorry, we couldn't find anyone with that id in our database")
	} else {

		c.Bind(&peeps[index])
		c.JSON(http.StatusOK, &peeps)
	}
}

// deletes a person from the db
func removePerson(c *gin.Context) {
	id := c.Param("id")
	index := -1
	// loop through the peeps slice and try to find the person with the matching id
	for idx, peep := range peeps {
		if peep.ID == id {
			index = idx
		}
	}

	// check if the id exists - if not, return a string message
	if index == -1 {
		c.String(http.StatusOK, "Sorry, we can't find a person with that id in our database")
	} else {
		peeps = append(peeps[:index], peeps[index+1:]...)
		c.JSON(http.StatusOK, &peeps)
	}
}
