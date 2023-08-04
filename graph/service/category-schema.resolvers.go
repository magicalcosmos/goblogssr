package service

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"time"

	"github.com/magicalcosmos/goblogssr/graph/generated"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

// Sort is the resolver for the sort field.
func (r *categoryResolver) Sort(ctx context.Context, obj *model.Category) (int, error) {
	panic(fmt.Errorf("not implemented: Sort - sort"))
}

// CreateAt is the resolver for the createAt field.
func (r *categoryResolver) CreateAt(ctx context.Context, obj *model.Category) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: CreateAt - createAt"))
}

// UpdateAt is the resolver for the updateAt field.
func (r *categoryResolver) UpdateAt(ctx context.Context, obj *model.Category) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: UpdateAt - updateAt"))
}

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// CategoryList is the resolver for the categoryList field.
func (r *queryResolver) CategoryList(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: CategoryList - categoryList"))
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

type categoryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: Categories - categories"))
}
