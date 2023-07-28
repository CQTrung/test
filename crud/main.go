package main

import (
	boostrap "trunne-crud/boostrap"
	routes "trunne-crud/route"

	"github.com/gin-gonic/gin"
)

func main() {
	var db = boostrap.InitDB()
	router := gin.New()
	routes.Setup(db, router)
	if err := router.Run(); err != nil {
		panic(err)
	}
}