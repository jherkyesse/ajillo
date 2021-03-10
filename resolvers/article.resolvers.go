package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"ajillo/generated"
	"ajillo/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, article model.CreateArticle) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, id string, update model.UpdateArticle) (*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ArticleList(ctx context.Context) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	// panic(fmt.Errorf("not implemented"))
	article := &model.Article{
		Title:       "蘭潭、番路、嘉義行",
		ID:          "3304",
		Author:      "蘭潭、番路、嘉義行",
		Description: "蘭潭、番路、嘉義行",
		Content:     "蘭潭、番路、嘉義行",
	}
	return article, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
