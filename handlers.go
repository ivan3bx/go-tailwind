package kpie

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Handler struct {
	*sqlx.DB
}

func NewHandler() *Handler {
	db := sqlx.MustOpen("sqlite3", ":memory:")
	return &Handler{DB: db}
}

// Index is the handler for the root path
func (h *Handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
