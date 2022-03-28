package utils

import (
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

// GenSlug Generate Unique Slug
func GenSlug(title string) string {
	slug := slug.Make(title)
	uid := uuid.New()
	uidstring := strings.Replace(uid.String(), "-", "", -1)
	return slug + uidstring
}
