package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//album represents data about the record album.
type album struct{
	ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
// albums slice to see record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}
//gerAlbumsByID - loop through the albums and match with the 
//ID passed to parameter and return
func getAlbumsByID(c *gin.Context){
	id := c.Param("id")
	for _, a := range albums{
	if a.ID == id{
		c.IndentedJSON(http.StatusOK, a)
		return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Messge": "Album not found"})
}
// Write a handler to add a new item "postAlbums" in JSON
//  When the client makes a POST request at /albums, you want to add the album described in the 
// request body to the existing albums data.
func postAlbums(c *gin.Context){
	var newAlbum album
	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// add the new albums to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
