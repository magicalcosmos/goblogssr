package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/magicalcosmos/goblogssr/graph/generated"
	"github.com/magicalcosmos/goblogssr/graph/model"
)

func (r *helloResolver) User(ctx context.Context, obj *model.Hello) (*model.User, error) {
	fmt.Println("hello")
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) He(ctx context.Context) ([]*model.Hello, error) {
	panic(fmt.Errorf("not implemented"))
}

// Hello returns generated.HelloResolver implementation.
func (r *Resolver) Hello() generated.HelloResolver { return &helloResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type helloResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
