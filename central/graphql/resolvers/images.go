package resolvers

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/utils"
)

func init() {
	schema := getBuilder()
	utils.Must(
		schema.AddQuery("images(query:String): [Image!]!"),
		schema.AddQuery("image(sha:ID!): Image"),
	)
}

// Images returns GraphQL resolvers for all images
func (resolver *Resolver) Images(ctx context.Context, args rawQuery) ([]*imageResolver, error) {
	if err := readImages(ctx); err != nil {
		return nil, err
	}
	q, err := args.AsV1Query()
	if err != nil {
		return nil, err
	}
	if q == nil {
		return resolver.wrapImages(resolver.ImageDataStore.GetImages(ctx))
	}
	return resolver.wrapListImages(
		resolver.ImageDataStore.SearchListImages(ctx, q))
}

// Image returns a graphql resolver for the identified image, if it exists
func (resolver *Resolver) Image(ctx context.Context, args struct{ Sha graphql.ID }) (*imageResolver, error) {
	if err := readImages(ctx); err != nil {
		return nil, err
	}
	return resolver.wrapImage(
		resolver.ImageDataStore.GetImage(ctx, string(args.Sha)))
}

func (resolver *Resolver) getImage(ctx context.Context, id string) *storage.Image {
	alert, ok, err := resolver.ImageDataStore.GetImage(ctx, id)
	if err != nil || !ok {
		return nil
	}
	return alert
}
