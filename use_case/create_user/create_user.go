package createuser

import (
	"github.com/nanduzz/go-simple-crud/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/src-d/go-errors.v1"
)

type SaveFunc func(user *entity.User) (*entity.User, error)

var ErrSaveFnNotSet = errors.NewKind("SaveFn is not set")

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func Execute(input *CreateUserInput, saveFn SaveFunc) (*CreateUserOutput, error) {
	if saveFn == nil {
		return nil, ErrSaveFnNotSet.New()
	}
	user := entity.NewUser(primitive.ObjectID{}, input.Username, input.Password)
	savedUser, err := saveFn(user)

	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:       savedUser.ID.Hex(),
		Username: savedUser.Username,
	}, nil
}
