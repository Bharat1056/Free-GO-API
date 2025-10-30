package types

// this will work after you install validator package
// pkg: github.com/go-playground/validator/v10
type Student struct {
	Id    int
	Email string `validate:"required"`
	Name  string `validate:"required"`
	Age   int    `validate:"required"`
}
