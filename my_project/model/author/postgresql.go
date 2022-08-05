package author

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"golang.org/x/net/context"
	"my_project/database"
	"my_project/logger"
	"my_project/model/Book"
	"strings"
)

type DatB struct {
	client database.Client
	logger *logger.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (d *DatB) Create(ctx context.Context, author *Author) error {
	var q = `
INSERT INTO author (name, born, death) 
VALUES ($1, $2, $3) 
RETURNING id`
	d.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	row := d.client.QueryRow(
		ctx,
		q,
		author.Name,
		author.Born,
		author.Death,
		123,
	)
	if err := row.Scan(&author.AuthorId); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			d.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

type SubQery struct {

}
// как записать данные из разных таблиц?
func (d *DatB FindAll(ctx context.Context) (a []SubQery, err error) {
	q := `
		SELECT public.author.name, public.book.name 
	FROM author, book WHERE public.author.id = public.book.author_id
	`
	d.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := d.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	type SubQery struct {
		a Author
		b Book.Book
	}

	authors := make([]SubQery, 0)

	for rows.Next() {
		var ath SubQery

		err = rows.Scan(&ath.a.Name, &ath.b.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
// не возвращает authors
	return authors, nil
}


func (d *DatB) FindOne(ctx context.Context, id string) ( Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id = $1
	`
	d.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var ath Author
	row := d.client.QueryRow(ctx, q, id)
	err := row.Scan(&ath.AuthorId, &ath.Name)
	if err != nil {
		return Author{}, err
	}

	return ath, nil
}

}

func (d *DatB) Update(ctx context.Context, author Author) error {
	//TODO implement me
	panic("implement me")
}

func (d *DatB) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(client database.Client, logger *logger.Logger) *DatB {
	return &DatB{
		client: client,
		logger: logger,
	}
}
