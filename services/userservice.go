package services

import (
	"context"

	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/graph/model"
)

func CreateUser(client *ent.Client, input *model.NewUser) (*ent.User, error) {
	user, err := client.Debug().User.Create().SetName(input.Name).SetUsername(input.Username).SetEmail(input.Email).SetPassword(input.Password).Save(context.Background())
	if err != nil {
		return &ent.User{}, err
	}
	return user, nil
}
