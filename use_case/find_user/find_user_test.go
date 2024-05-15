package finduser

import (
	"errors"
	"testing"

	"github.com/nanduzz/go-simple-crud/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindUserById(t *testing.T) {

	findUserFn := func(id string) (*entity.User, error) {
		return &entity.User{
			ID:       primitive.NewObjectID(),
			Username: "test",
		}, nil
	}

	t.Run("Given a nil findUserFn, should return an error", func(t *testing.T) {
		// Given
		input := Input{
			Id: "1",
		}

		// When
		_, err := Execute(input, nil)

		// Then
		if err == nil {
			t.Error("expected an error, got nil")
		}

		if !ErrFindUserByIdFnNotSet.Is(err) {
			t.Errorf("expected %q, got %q", ErrFindUserByIdFnNotSet, err)
		}
	})

	t.Run("Given am empty id, should return an error", func(t *testing.T) {
		// Given
		input := Input{}
		// When
		_, err := Execute(input, findUserFn)
		// Then
		if err == nil {
			t.Error("expected an error, got nil")
		}

		if !ErrFindUserByIdInputIdRequired.Is(err) {
			t.Errorf("expected %q, got %q", ErrFindUserByIdInputIdRequired, err)
		}
	})

	t.Run("Given a valid id, should return a user", func(t *testing.T) {
		// Given
		input := Input{
			Id: "1",
		}

		// When
		output, err := Execute(input, findUserFn)

		// Then
		if err != nil {
			t.Fatalf("expected nil, got %q", err)
		}

		if output == nil {
			t.Fatal("expected an output, got nil")
		}

		if output.Id == "" {
			t.Error("expected an id, got an empty string")
		}

		if output.Name == "" {
			t.Error("expected a name, got an empty string")
		}
	})

	t.Run("Given a valid id, when findUserFn fails, should return an error", func(t *testing.T) {
		// Given
		input := Input{
			Id: "1",
		}

		findUserFn := func(id string) (*entity.User, error) {
			return nil, errors.New("error")
		}

		// When
		_, err := Execute(input, findUserFn)

		// Then
		if err == nil {
			t.Error("expected an error, got nil")
		}

	})

}
