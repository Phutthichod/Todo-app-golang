package api

import (
	"net/http"
	"todo-go/model"
	"todo-go/repository"

	"github.com/gin-gonic/gin"
)

type NoteAPI struct {
	NoteRepository *repository.NoteRepository
}

func (noteAPI *NoteAPI) ListNoteHandler(context *gin.Context) {
	notes, err := noteAPI.NoteRepository.ListNote()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, notes)
}
func (noteAPI *NoteAPI) UpdateNoteHandler(context *gin.Context) {
	id := context.Param("id")
	var note model.Note
	err := context.ShouldBindJSON(&note)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = noteAPI.NoteRepository.UpdateNote(id, note)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, nil)
}
func (noteAPI *NoteAPI) DeleteNoteHandler(context *gin.Context) {
	id := context.Param("id")
	err := noteAPI.NoteRepository.DeleteNote(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, nil)
}
func (noteAPI *NoteAPI) CreateNoteHandler(context *gin.Context) {
	var note model.Note
	err := context.ShouldBindJSON(&note)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	id, err := noteAPI.NoteRepository.AddNote(note)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, id)
}
