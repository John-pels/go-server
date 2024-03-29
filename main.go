package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbums)
	router.DELETE("/albums/:id", deleteAlbums)
	router.Run("localhost:8080")
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for _, a := range albums {

		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// putAlbums replaces an existing album in the albums slice.
func putAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for i, a := range albums {

		if a.ID == id {
			albums[i] = newAlbum
			c.IndentedJSON(http.StatusOK, newAlbum)
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// deleteAlbums deletes an album from the albums slice.
func deleteAlbums(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for i, a := range albums {

		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
