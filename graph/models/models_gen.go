// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"app-api/ent"
	"fmt"
	"io"
	"strconv"
)

type DevSession struct {
	ID         uint64               `json:"id"`
	PreviewURL string               `json:"previewURL"`
	Component  *ent.CustomComponent `json:"component"`
}

type PrepareDevSessionInput struct {
	ComponentID *uint64 `json:"componentID"`
}

// Define an input type for the mutation below.
// https://graphql.org/learn/schema/#input-types
//
// Note that, this type is mapped to the generated
// input type in mutation_input.go.
type TemplateSectionInput struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Cid       string  `json:"cid"`
	Area      string  `json:"area"`
	Component *string `json:"component"`
	Display   *bool   `json:"display"`
	DeletedAt *string `json:"deletedAt"`
	Position  *int    `json:"position"`
}

type User struct {
	ID       uint64    `json:"id"`
	UserName *string   `json:"userName"`
	Email    *string   `json:"email"`
	Roles    []*string `json:"roles"`
}

type SaveType string

const (
	SaveTypeAutoSave   SaveType = "AUTO_SAVE"
	SaveTypeNormalSave SaveType = "NORMAL_SAVE"
)

var AllSaveType = []SaveType{
	SaveTypeAutoSave,
	SaveTypeNormalSave,
}

func (e SaveType) IsValid() bool {
	switch e {
	case SaveTypeAutoSave, SaveTypeNormalSave:
		return true
	}
	return false
}

func (e SaveType) String() string {
	return string(e)
}

func (e *SaveType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SaveType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SaveType", str)
	}
	return nil
}

func (e SaveType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}