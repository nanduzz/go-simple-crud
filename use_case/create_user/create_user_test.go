package createuser

import (
	"errors"
	"testing"

	"github.com/nanduzz/go-simple-crud/entity"
	"github.com/nanduzz/go-simple-crud/util"
)

type UserSaverMock struct {
	SaveFn               func(u *entity.User) (*entity.User, error)
	SaveFuncInvokedTimes int
}

func (usm *UserSaverMock) Save(u *entity.User) (*entity.User, error) {
	usm.SaveFuncInvokedTimes++
	return usm.SaveFn(u)
}

func TestCreateUser(t *testing.T) {
	t.Run("should return an error when SaveFn is not set", func(t *testing.T) {
		// Given
		input := &CreateUserInput{}

		// When
		_, err := Execute(input, nil)

		// Then
		util.AssertEqual(t, err.Error(), ErrSaveFnNotSet.Message)
	})
	t.Run("should create a user", func(t *testing.T) {
		// Given
		input := &CreateUserInput{
			Username: "test",
			Password: "1234",
		}

		userSaver := &UserSaverMock{
			SaveFn: func(u *entity.User) (*entity.User, error) {
				return u, nil
			},
		}
		// When
		output, err := Execute(input, userSaver.Save)

		// Then
		util.AssertNoError(t, err)

		util.AssertEqual(t, output.Username, "test")
		util.AssertEqual(t, userSaver.SaveFuncInvokedTimes, 1)
	})

	t.Run("should return an error when user saver fails", func(t *testing.T) {
		// Given
		input := &CreateUserInput{
			Username: "test",
			Password: "1234",
		}

		// When
		saveFn := func(u *entity.User) (*entity.User, error) {
			return nil, errors.New("error")
		}
		_, err := Execute(input, saveFn)
		// Then
		util.AssertEqual(t, err.Error(), "error")

	})
}
