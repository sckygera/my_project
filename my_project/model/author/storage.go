package author

import "golang.org/x/net/context"

type Storage interface {
	Create(ctx context.Context, author Author) (string, error)
	FindAll(ctx context.Context) (a []Author, err error)
	FindOne(ctx context.Context, id string) (Author error)
	Update(ctx context.Context, author Author) error
	Delete(ctx context.Context, id string) error
}
