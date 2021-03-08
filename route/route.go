package route

import (
	"todo-go/api"
	"todo-go/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NoteRoute(connectionDB *mongo.Client, route *gin.Engine) {
	noteRepository := repository.NoteRepository{
		ConnectionDB: connectionDB,
	}
	noteAPI := api.NoteAPI{
		NoteRepository: &noteRepository,
	}
	route.GET("/api/v1/notes", noteAPI.ListNoteHandler)
	route.POST("/api/v1/notes", noteAPI.CreateNoteHandler)
	route.PUT("/api/v1/notes/:id", noteAPI.UpdateNoteHandler)
	route.DELETE("/api/v1/notes/:id", noteAPI.DeleteNoteHandler)
}
