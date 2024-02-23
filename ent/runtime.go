// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/FournyP/go-orm-benchmarks/ent/comment"
	"github.com/FournyP/go-orm-benchmarks/ent/post"
	"github.com/FournyP/go-orm-benchmarks/ent/schema"
	"github.com/FournyP/go-orm-benchmarks/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	commentFields := schema.Comment{}.Fields()
	_ = commentFields
	// commentDescText is the schema descriptor for text field.
	commentDescText := commentFields[1].Descriptor()
	// comment.TextValidator is a validator for the "text" field. It is called by the builders before save.
	comment.TextValidator = commentDescText.Validators[0].(func(string) error)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescTitle is the schema descriptor for title field.
	postDescTitle := postFields[1].Descriptor()
	// post.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	post.TitleValidator = postDescTitle.Validators[0].(func(string) error)
	// postDescContent is the schema descriptor for content field.
	postDescContent := postFields[2].Descriptor()
	// post.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	post.ContentValidator = postDescContent.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
}
