package service

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/magicalcosmos/goblogssr/graph"
	"github.com/magicalcosmos/goblogssr/graph/dao"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.NewUser) (bool, error) {
	user := &model.User{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Username: input.Username,
		Password: input.Password,
	}
	fmt.Println(user.Password)
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	category, err := dao.CreateCategory(input)
	if err != nil {
		fmt.Println(err)
	}
	return category, err
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, input model.NewCategory) (string, error) {
	message, err := dao.DeleteCategory(input)
	return message, err
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
