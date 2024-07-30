package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/blog/internal/modules/repositories"
)

type Controler struct {
	articleRepository repositories.ArticleRepository
}

func NewControler() *Controler {
	return &Controler{
		articleRepository: *repositories.NewArticleRepository(),
	}
}

func (cl *Controler) HandleGetArticles(c *gin.Context) {
	c.JSON(http.StatusOK, cl.articleRepository.List(8))
}
