package findall

import (
	"github.com/nanduzz/go-simple-crud/entity"
	"gopkg.in/src-d/go-errors.v1"
)

type Input struct {
}

type Output struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var ErrFindAllUsersFnNotSet = errors.NewKind("FindAllUsersFn is not set")

type FindAllUsersFn func(query map[string]interface{}) ([]*entity.User, error)

func Execute(findAllUsersFn FindAllUsersFn) ([]Output, error) {
	if findAllUsersFn == nil {
		return nil, ErrFindAllUsersFnNotSet.New()
	}

	users, err := findAllUsersFn(map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var outputs = make([]Output, 0, len(users))

	for _, user := range users {
		outputs = append(outputs, Output{
			Id:   user.ID.Hex(),
			Name: user.Username,
		})
	}

	return outputs, nil

}
