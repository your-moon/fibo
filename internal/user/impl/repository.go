package impl

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"

	databaseImpl "fibo/internal/base/database/impl"
	"fibo/internal/base/errors"
	"fibo/internal/user"
)

type UserRepositoryOpts struct {
	ConnManager databaseImpl.ConnManager
}

func NewUserRepository(opts UserRepositoryOpts) user.UserRepository {
	return &userRepository{
		ConnManager: opts.ConnManager,
	}
}

type userRepository struct {
	databaseImpl.ConnManager
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]user.UserModel, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Select(
			"user_id",
			"firstname",
			"lastname",
			"email",
			"password",
			"reputation",
		).
		From("users").
		ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	rows, err := r.Conn(ctx).Query(ctx, sql)
	if err != nil {
		return nil, errors.Wrap(err, errors.DatabaseError, "get all users failed")
	}
	defer rows.Close()

	var models []user.UserModel

	for rows.Next() {
		var model user.UserModel
		err = rows.Scan(
			&model.Id,
			&model.FirstName,
			&model.LastName,
			&model.Email,
			&model.Password,
			&model.Reputation,
		)
		if err != nil {
			return nil, errors.Wrap(err, errors.DatabaseError, "scan user failed")
		}

		models = append(models, model)
	}

	return models, nil
}

func (r *userRepository) Add(ctx context.Context, model user.UserModel) (int64, error) {
	fmt.Printf("Add user: %+v\n", model)
	sql, _, err := databaseImpl.QueryBuilder.
		Insert("users").
		Rows(databaseImpl.Record{
			"firstname":  model.FirstName,
			"lastname":   model.LastName,
			"email":      model.Email,
			"password":   model.Password,
			"reputation": model.Reputation,
		}).
		Returning("user_id").
		ToSQL()
	if err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&model.Id); err != nil {
		return 0, parseAddUserError(&model, err)
	}

	return model.Id, nil
}

func (r *userRepository) Update(ctx context.Context, model user.UserModel) (int64, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Update("users").
		Set(databaseImpl.Record{
			"firstname":  model.FirstName,
			"lastname":   model.LastName,
			"email":      model.Email,
			"password":   model.Password,
			"reputation": model.Reputation,
		}).
		Where(databaseImpl.Ex{"user_id": model.Id}).
		Returning("user_id").
		ToSQL()
	if err != nil {
		return 0, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	if err := row.Scan(&model.Id); err != nil {
		return 0, parseUpdateUserError(&model, err)
	}

	return model.Id, nil
}

func (r *userRepository) GetById(ctx context.Context, userId int64) (user.UserModel, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Select(
			"firstname",
			"lastname",
			"email",
			"password",
			"reputation",
		).
		From("users").
		Where(databaseImpl.Ex{"user_id": userId}).
		ToSQL()
	if err != nil {
		return user.UserModel{}, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	model := user.UserModel{Id: userId}

	err = row.Scan(
		&model.FirstName,
		&model.LastName,
		&model.Email,
		&model.Password,
		&model.Reputation,
	)
	if err != nil {
		return user.UserModel{}, parseGetUserByIdError(userId, err)
	}

	return model, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (user.UserModel, error) {
	sql, _, err := databaseImpl.QueryBuilder.
		Select(
			"user_id",
			"firstname",
			"lastname",
			"password",
			"reputation",
		).
		From("users").
		Where(databaseImpl.Ex{"email": email}).
		ToSQL()
	if err != nil {
		return user.UserModel{}, errors.Wrap(err, errors.DatabaseError, "syntax error")
	}

	row := r.Conn(ctx).QueryRow(ctx, sql)

	model := user.UserModel{Email: email}

	err = row.Scan(
		&model.Id,
		&model.FirstName,
		&model.LastName,
		&model.Password,
		&model.Reputation,
	)
	if err != nil {
		return user.UserModel{}, parseGetUserByEmailError(email, err)
	}

	return model, nil
}

func parseAddUserError(user *user.UserModel, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		switch pgError.ConstraintName {
		case "users_email_key":
			return errors.Wrapf(
				err,
				errors.AlreadyExistsError,
				"user with email \"%s\" already exists",
				user.Email,
			)
		default:
			return errors.Wrapf(err, errors.DatabaseError, "add user failed")
		}
	}

	return errors.Wrapf(err, errors.DatabaseError, "add user failed")
}

func parseUpdateUserError(user *user.UserModel, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.UniqueViolation {
		return errors.Wrapf(
			err,
			errors.AlreadyExistsError,
			"user with email \"%s\" already exists",
			user.Email,
		)
	}

	return errors.Wrapf(err, errors.DatabaseError, "update user failed")
}

func parseGetUserByIdError(userId int64, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.Wrapf(err, errors.NotFoundError, "user with id \"%d\" not found", userId)
	}
	if err.Error() == "no rows in result set" {
		return errors.Wrapf(err, errors.NotFoundError, "user with id \"%d\" not found", userId)
	}

	return errors.Wrap(err, errors.DatabaseError, "get user by id failed")
}

func parseGetUserByEmailError(email string, err error) error {
	pgError, isPgError := err.(*pgconn.PgError)

	if isPgError && pgError.Code == pgerrcode.NoDataFound {
		return errors.Wrapf(err, errors.NotFoundError, "user with email \"%s\" not found", email)
	}
	if err.Error() == "no rows in result set" {
		return errors.Wrapf(err, errors.NotFoundError, "user with email \"%s\" not found", email)
	}

	return errors.Wrap(err, errors.DatabaseError, "get user by email failed")
}
