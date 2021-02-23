package service

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph"
	"github.com/magicalcosmos/goblogssr/graph/dao"
	"github.com/magicalcosmos/goblogssr/graph/model"
	"math/rand"
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

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return dao.GetUserInfo(), nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
