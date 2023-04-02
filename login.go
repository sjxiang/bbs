package bbs

import "context"


type LoginInput struct {
	Email    string
	Username *string
}

type Service struct {

}


func (svc *Service) Login(ctx context.Context, in LoginInput) error {

	return nil
}