// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/ArifulProtik/gograph-notes/ent"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	User         *ent.User `json:"user"`
	Accestoken   *string   `json:"accestoken"`
	RefreshToken *string   `json:"refreshToken"`
}

type MuiltipleNotes struct {
	Notes    []*ent.Notes `json:"notes"`
	Page     *int         `json:"Page"`
	Perpage  *int         `json:"Perpage"`
	Lastpage *int         `json:"lastpage"`
}

type NewNote struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type NewUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResNotes struct {
	Note *ent.Notes `json:"Note"`
	User *ent.User  `json:"User"`
}

type UserRes struct {
	User  *ent.User `json:"user"`
	Error *string   `json:"error"`
}
