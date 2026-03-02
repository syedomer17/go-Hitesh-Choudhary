package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(c *gin.Context){
	// per req obhject from gin 

	var req CreateNoteRequest 

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	now := time.Now().UTC()

	note := Note{
		ID: primitive.NewObjectID(),
		Title: req.Title,
		Content: req.Content,
		Pinned: req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}

	created, err := h.repo.Create(c.Request.Context(), note)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Failed to create the notes",
		})
		return
	}

	c.JSON(http.StatusCreated, created)
}