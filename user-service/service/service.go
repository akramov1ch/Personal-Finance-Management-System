package service

import (
	"context"
	"log"
	"user-service/internal/repo"
	"user-service/protos/genuser"
)

type UserServicee struct {
	genuser.UnimplementedUserServiceServer
	serv *repo.UserServer
	// Lg   *log.Logger
}

func ConnServer() *UserServicee {
	r := repo.Connect()

	return &UserServicee{serv: r}
}

func (s *UserServicee) RegisterUser(ctx context.Context, req *genuser.RegisterReq) (*genuser.RegisterResp, error) {
	log.Println(req)
	resp, err := s.serv.RegisterUserr(ctx, req)
	log.Println("RegisterUser Servicdaaa>>>>>>>>>>>>")
	if err != nil {
		log.Println("Xatolik repo registeruserda", err)
		return nil, err
	}

	return resp, nil
}

func (s *UserServicee) VerifyUser(ctx context.Context, req *genuser.VerifyReq) (*genuser.VerifyResp, error) {
	log.Println("req>>>>>", req)
	log.Println("VerifyUserda>>>>>>>>")
	resp, err := s.serv.VerifyUserr(ctx, req)
	if err != nil {
		log.Println("Xatolik verifyda", err)
		return nil, err
	}

	return resp, nil
}

func (u *UserServicee) DeleteUser(ctx context.Context, req *genuser.DeleteUserReq) (*genuser.DeleteUserResp, error) {
	resp, err := u.serv.DeleteUserr(ctx, req)
	if err != nil {
		log.Println("Xatolik deleteuserda", err)
		return nil, err
	}

	return resp, nil
}

func (u *UserServicee) GetUser(ctx context.Context, req *genuser.GetUserReq) (*genuser.GetUserResp, error) {
	resp, err := u.serv.GetUserr(ctx, req)
	if err != nil {
		log.Println("Error on getuser", err)
		return nil, err
	}

	return resp, nil
}

func (u *UserServicee) LoginUser(ctx context.Context, req *genuser.LoginReq) (*genuser.LoginResp, error) {
	resp, err := u.serv.Loginn(ctx, req)
	if err != nil {
		log.Println("Xatolik loginda", err)
		return nil, err
	}

	return resp, nil

}
