package service

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/magicalcosmos/goblogssr/graph"
	"github.com/magicalcosmos/goblogssr/graph/dao"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return dao.GetUserInfo(), nil
}

func (r *queryResolver) Category(ctx context.Context) ([]*model.Category, error) {
	category, err := dao.GetCategory()
	return category, err
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
