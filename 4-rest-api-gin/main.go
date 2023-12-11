package main

import "net/http"; 
import "github.com/gin-gonic/gin"; 

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Grace", Artist: "Jeff Buckley", Price: 10.88},
	{ID: "2", Title: "Gently Disturbed", Artist: "Avashi Cohen Trio", Price: 9.99},
	{ID: "3", Title: "24k", Artist: "Bruno Mars", Price: 15.07},
	{ID: "4", Title: "Invisible Cinema", Artist: "Aaron Parks", Price: 21.99},
}

func main() {
  router := gin.Default(); 
  router.GET("/albums", getAlbums); 
  router.Run("localhost:8080");
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums); 
}


