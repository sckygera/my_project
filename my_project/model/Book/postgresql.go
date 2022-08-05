package Book

import (
	"golang.org/x/net/context"
	"my_project/logger"
)

type db struct {
}

func (d *db) Create(ctx context.Context, book Book) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) FindAll(ctx context.Context) (b []Book, err error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) FindOne(ctx context.Context, id string) (Book error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) Update(ctx context.Context, author Book) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(collection string, logger *logger.Logger) Storage {
	return &db{}
}
