package repositories

import (
	"github.com/mojtabafarzaneh/blog/internal/modules/models"
	"github.com/mojtabafarzaneh/blog/pkg/database"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connect(),
	}
}

type Article interface {
	List(limit int) *[]models.Article
}

func (r *ArticleRepository) List(limit int) *[]models.Article {
	var articles []models.Article

	r.DB.Limit(limit).Joins("User").Order("id").Find(&articles)

	return &articles

}
