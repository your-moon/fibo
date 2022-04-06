package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4"

	"go-backend-template/internal/util/errors"
)

// Service

type Service interface {
	Repos() Repos
	BeginTx(ctx context.Context) (Transaction, error)
}

func NewService(client *Client) Service {
	repos := newRepos(client.pool, goqu.Dialect("postgres"))

	service := service{
		repos:  repos,
		client: client,
	}

	return &service
}

type service struct {
	repos  Repos
	client *Client
}

func (s *service) Repos() Repos {
	return s.repos
}

func (s *service) BeginTx(ctx context.Context) (Transaction, error) {
	conn, err := s.client.pool.Begin(ctx)
	if err != nil {
		return nil, errors.New(errors.DatabaseError, "cannot open transaction").SetInternal(err)
	}

	repos := newRepos(conn, goqu.Dialect("postgres"))
	transaction := newTransaction(conn, repos)

	return transaction, nil
}

// Transaction

type Transaction interface {
	Repos() Repos
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

func newTransaction(conn pgx.Tx, repos Repos) Transaction {
	return &transaction{
		repos: repos,
		conn:  conn,
	}
}

type transaction struct {
	repos Repos
	conn  pgx.Tx
}

func (t *transaction) Repos() Repos {
	return t.repos
}

func (t *transaction) Commit(ctx context.Context) error {
	err := t.conn.Commit(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot commit transaction").SetInternal(err)
	}

	return nil
}

func (t *transaction) Rollback(ctx context.Context) error {
	err := t.conn.Rollback(ctx)
	if err != nil {
		return errors.New(errors.DatabaseError, "cannot rollback transaction").SetInternal(err)
	}

	return nil
}
