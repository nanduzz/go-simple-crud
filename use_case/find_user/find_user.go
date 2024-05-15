package finduser

import (
	"github.com/nanduzz/go-simple-crud/entity"
	"gopkg.in/src-d/go-errors.v1"
)

type Input struct {
	Id string
}

type Output struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var ErrFindUserByIdFnNotSet = errors.NewKind("FindUserByIdFn is not set")
var ErrFindUserByIdInputIdRequired = errors.NewKind("id is required")

type FindUserByIdFn func(id string) (*entity.User, error)

func Execute(input Input, findUser FindUserByIdFn) (*Output, error) {
	if findUser == nil {
		return nil, ErrFindUserByIdFnNotSet.New()
	}

	if input.Id == "" {
		return nil, ErrFindUserByIdInputIdRequired.New()
	}

	user, err := findUser(input.Id)
	if err != nil {
		return nil, err
	}

	output := &Output{
		Id:   user.ID.Hex(),
		Name: user.Username,
	}

	return output, nil
}
