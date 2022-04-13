package usecase

import (
	"context"

	"go-backend-template/internal/database"
	"go-backend-template/internal/usecase/dto"
	"go-backend-template/internal/util/crypto"
)

type UserUsecases interface {
	Add(ctx context.Context, in dto.UserAdd) (int64, error)
	UpdateInfo(ctx context.Context, in dto.UserUpdateInfo) error
	ChangePassword(ctx context.Context, in dto.UserChangePassword) error
	GetById(ctx context.Context, userId int64) (dto.User, error)
}

type userUsecases struct {
	db     database.Service
	crypto crypto.Crypto
	config Config
}

func (u *userUsecases) Add(ctx context.Context, in dto.UserAdd) (userId int64, err error) {
	user, err := in.MapTo()
	if err != nil {
		return 0, err
	}
	if err := user.HashPassword(u.crypto); err != nil {
		return 0, err
	}

	// Transaction demonstration
	err = u.db.BeginTx(ctx, func(ctx context.Context) error {
		userId, err = u.db.UserRepo().Add(ctx, user)
		if err != nil {
			return err
		}
		user.Id = userId

		userId, err = u.db.UserRepo().Update(ctx, user)
		if err != nil {
			return err
		}
		return nil
	})

	return userId, err
}

func (u *userUsecases) UpdateInfo(ctx context.Context, in dto.UserUpdateInfo) error {
	user, err := u.db.UserRepo().GetById(ctx, in.Id)
	if err != nil {
		return err
	}
	err = user.UpdateInfo(in.FirstName, in.LastName, in.Email)
	if err != nil {
		return err
	}
	_, err = u.db.UserRepo().Update(ctx, user)

	return err
}

func (u *userUsecases) ChangePassword(ctx context.Context, in dto.UserChangePassword) error {
	user, err := u.db.UserRepo().GetById(ctx, in.Id)
	if err != nil {
		return err
	}
	if err = user.ChangePassword(in.Password, u.crypto); err != nil {
		return err
	}
	_, err = u.db.UserRepo().Update(ctx, user)

	return err
}

func (u *userUsecases) GetById(ctx context.Context, userId int64) (out dto.User, err error) {
	user, err := u.db.UserRepo().GetById(ctx, userId)
	if err != nil {
		return out, err
	}

	return out.MapFrom(user), nil
}
