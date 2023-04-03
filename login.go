package bbs

import (
	"context"
	"errors"

	uuid "github.com/satori/go.uuid"
)


type LoginInput struct {
	Email    string
	Username *string
}




func (svc *Service) Login(ctx context.Context, in LoginInput) (User, error) {
	var out User

	exists, err := svc.Queries.UserExistsByEmail(ctx, in.Email)
	if err != nil {
		return out, err
	}

	if exists {
		return svc.Queries.UserByEmail(ctx, in.Email)
	}

	if in.Username == nil {
		return out, errors.New("user not found")
	}

	exists, err = svc.Queries.UserExistsByEmail(ctx, *in.Username)
	if err != nil {
		return out, err
	}

	if exists {
		return out, errors.New("username taken")  // 用户名被占用
	}

	userID := genID()
	createdAt, err := svc.Queries.CreateUser(ctx, CreateUserParams{
		UserID:   userID,
		Email:    in.Email,
		Username: *in.Username,
	})
	if err != nil {
		return out, err
	}

	return User{
		ID:        userID,
		Email:     in.Email,
		Username:  *in.Username,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}


func genID() string {
	return uuid.NewV4().String()
}