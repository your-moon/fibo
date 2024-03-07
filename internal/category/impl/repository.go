package impl

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	databaseImpl "fibo/internal/base/database/impl"
	"fibo/internal/base/errors"
	"fibo/internal/category"
)

type CatRepositoryOpts struct {
	ConnManager databaseImpl.ConnManager
}

func NewCatRepository(opts CatRepositoryOpts) category.CatRepository {
	return &catRepository{
		ConnManager: opts.ConnManager,
	}
}

type catRepository struct {
	databaseImpl.ConnManager
}

func (p *catRepository) GetCategories(
	ctx context.Context,
) ([]category.CategoryModel, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Select("id", "name").
		From("categories").
		ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	rows, err := p.Conn(ctx).Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "get categories failed")
	}
	fmt.Println(err)
	defer rows.Close()

	var result []category.CategoryModel
	for rows.Next() {
		var cat category.CategoryModel
		err = rows.Scan(&cat.Id, &cat.Name)
		fmt.Println(err)
		if err != nil {
			return nil, errors.Wrap(err, errors.DatabaseError, "scan category failed")
		}

		result = append(result, cat)
	}

	return result, nil
}

func (p *catRepository) GetById(
	ctx context.Context,
	id int64,
) (*category.CategoryModel, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Select("id", "name").
		From("categories").
		Where(goqu.Ex{"id": id}).
		ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	var result category.CategoryModel
	err = p.Conn(ctx).QueryRow(ctx, sql).Scan(&result.Id, &result.Name)
	if err != nil {
		return nil, parseGetCatError(err)
	}

	return &result, nil
}

func (p *catRepository) Add(
	ctx context.Context,
	cat category.CategoryModel,
) (int64, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Insert("categories").
		Rows(goqu.Record{"name": cat.Name}).
		ToSQL()
	if err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	result, err := p.Conn(ctx).Exec(ctx, sql)
	if err != nil {
		return 0, parseAddCatError(&cat, err)
	}

	return result.RowsAffected(), nil
}

func parseGetCatError(err error) error {
	pgErr, isPgErr := err.(*pgconn.PgError)

	if isPgErr && pgErr.Code == pgerrcode.NoDataFound {
		return errors.Wrapf(err, errors.DatabaseError, "no category found")
	}
	return errors.Wrapf(err, errors.DatabaseError, "get category failed")
}

func parseAddCatError(post *category.CategoryModel, err error) error {
	pgErr, isPgErr := err.(*pgconn.PgError)

	if isPgErr && pgErr.Code == pgerrcode.UniqueViolation {
		return errors.Wrapf(err, errors.DatabaseError, "unique violation")
	}
	return errors.Wrapf(err, errors.DatabaseError, "add post failed")
}
