//go:generate go run github.com/99designs/gqlgen

package graph

import "github.com/magicalcosmos/goblogssr/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver resolver
type Resolver struct {
	users []*model.User
}
