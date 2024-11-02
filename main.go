package main

import (
	"net/http"

	events "example.com/api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := events.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event events.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"err": "Invalid data sent"})
	}

	event.ID = 1
	event.UserId = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"msg": "Event was created", "event": event})
}
