package services

import (
	"context"
	"errors"

	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/ent/notes"
	"github.com/ArifulProtik/gograph-notes/graph/model"
	"github.com/ArifulProtik/gograph-notes/utils"
)

func CreateNote(client *ent.Client, note *model.NewNote, user ent.User) (*ent.Notes, error) {
	if note.Title == "" || note.Body == "" || note.Tags == nil {
		return &ent.Notes{}, errors.New("field cant be empty")
	}
	slug := utils.GenSlug(note.Title)
	newnote, err := client.Debug().Notes.Create().SetAuthor(&user).SetAuthorID(user.ID).SetTitle(note.Title).SetBody(note.Body).SetTags(note.Tags).SetSlug(slug).Save(context.Background())
	if err != nil {
		return &ent.Notes{}, err
	}
	return newnote, nil

}

func Mynotes(client *ent.Client, user ent.User, perpage int, page int) ([]*ent.Notes, int, error) {
	count, err := client.User.QueryNotes(&user).Count(context.Background())
	if err != nil {
		return nil, 0, errors.New(err.Error())
	}
	data, err := client.User.QueryNotes(&user).Order(ent.Desc(notes.FieldCreatedAt)).Limit(perpage).Offset(page).WithAuthor().All(context.Background())
	if err != nil {
		return nil, 0, err
	}
	return data, count, nil

}

func Singlenote(client *ent.Client, slug string) (*ent.Notes, error) {
	note, err := client.Notes.Query().Where(notes.SlugEQ(slug)).Only(context.Background())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return note, nil
}
