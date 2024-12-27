package schema

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"

	"github.com/nmarsollier/ordersgo/internal/graph/model"
	"github.com/nmarsollier/ordersgo/internal/graph/resolvers"
)

// FindOrderByID is the resolver for the findOrderByID field.
func (r *entityResolver) FindOrderByID(ctx context.Context, id string) (*model.Order, error) {
	return resolvers.FindByOrderId(ctx, id)
}

// Entity returns model.EntityResolver implementation.
func (r *Resolver) Entity() model.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }