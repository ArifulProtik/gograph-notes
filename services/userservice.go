package services

import (
	"context"
	"errors"

	"github.com/ArifulProtik/gograph-notes/auth"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/ent/user"
	"github.com/ArifulProtik/gograph-notes/graph/model"
	"github.com/ArifulProtik/gograph-notes/utils"
)

func CreateUser(client *ent.Client, input *model.NewUser) (*ent.User, error) {

	if input.Password == "" {
		return &ent.User{}, errors.New("password field can't be empty")
	}
	newpass, _ := utils.HashBeforeSave(input.Password)
	user, err := client.Debug().User.Create().SetName(input.Name).SetUsername(input.Username).SetEmail(input.Email).SetPassword(string(newpass)).Save(context.Background())
	if err != nil {
		return &ent.User{}, err
	}
	return user, nil
}
func SigninUser(client *ent.Client, input *model.Login) (*model.LoginRes, error) {
	if input.Email == "" || input.Password == "" {
		return &model.LoginRes{}, errors.New("email or password field empty")
	}
	user, err := client.Debug().User.Query().Where(user.EmailEQ(input.Email)).First(context.Background())
	if err != nil {
		return &model.LoginRes{}, err
	}
	if user == nil {
		return &model.LoginRes{}, errors.New("username or password incorrect")
	}
	err = utils.VerifyPass(user.Password, input.Password)
	if err != nil {
		return &model.LoginRes{}, errors.New("email or password incorrect")
	}
	ac_token, _ := auth.CreateAccessToken(user.ID)
	rf_token, _ := auth.CreateRfreshToken(user.ID)
	return &model.LoginRes{
		User:         user,
		Accestoken:   &ac_token,
		RefreshToken: &rf_token,
	}, nil
}
