package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// FIX: Capitalized RegisterRoutes to export it (make it public)
// In Go, only capitalized functions are accessible from other packages
func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	repo := NewRepo(db)

	h := NewHandler(repo)

	notesGroup := r.Group("/notes")
	{
		notesGroup.POST("", h.CreateNote)
		notesGroup.GET("/get-notes",h.ListNotes)
		notesGroup.GET("/:id", h.GetNoteById)
		notesGroup.PUT("/:id", h.UpdateNoteById)
		notesGroup.DELETE("/:id", h.DeleteNoteById)
	}
}
