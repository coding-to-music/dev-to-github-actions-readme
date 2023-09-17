package collector

import (
	"context"

	"github.com/coding-to-music/dev-to-github-actions-readme/model"
)

type ArticleService interface {
	GetArticles(ctx context.Context, username string, page, perPage int) ([]model.Article, error)
}
