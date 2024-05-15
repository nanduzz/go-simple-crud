package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nanduzz/go-simple-crud/entity"
	createuser "github.com/nanduzz/go-simple-crud/use_case/create_user"
	finduser "github.com/nanduzz/go-simple-crud/use_case/find_user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setUpTestFindUserHandler(t testing.TB) {
	t.Helper()
	findUserByIdFn = func(id string) (*entity.User, error) {
		return &entity.User{
			ID:       primitive.NewObjectID(),
			Username: "user1",
		}, nil
	}

}

// https://blog.canopas.com/golang-unit-tests-with-test-gin-context-80e1ac04adcd
func TestFindUserHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	t.Run("should return 500 status code when findUserByIdFn is not set", func(t *testing.T) {
		//Given
		setUpTestFindUserHandler(t)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		findUserByIdFn = nil

		//When
		FindUserHandlerById(ctx)

		//Then
		assert.Equal(t, 500, w.Code)
	})

	t.Run("should return 200 status code and user data", func(t *testing.T) {
		//Given
		setUpTestFindUserHandler(t)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		//When
		FindUserHandlerById(ctx)

		//Then
		assert.Equal(t, 200, w.Code)
		// convert the response body to object and compare
		var user finduser.Output
		err := json.Unmarshal(w.Body.Bytes(), &user)

		assert.NoError(t, err)
		assert.Equal(t, "user1", user.Name)

	})

	t.Run("should return 400 status code when id is not provided", func(t *testing.T) {
		//Given
		setUpTestFindUserHandler(t)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		//When
		FindUserHandlerById(ctx)

		//Then
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 500 status code when findUserByIdFn returns error", func(t *testing.T) {
		//Given
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})
		findUserByIdFn = func(id string) (*entity.User, error) {
			return nil, assert.AnError
		}

		//When
		FindUserHandlerById(ctx)

		//Then
		assert.Equal(t, 500, w.Code)
	})

}

func setUpTestCreateUserHandler(t testing.TB) {
	t.Helper()
	saveUserFn = func(u *entity.User) (*entity.User, error) {
		return u, nil
	}
}

func TestCreateUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return 200 status code when user is created", func(t *testing.T) {
		//Given
		setUpTestCreateUserHandler(t)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		inputStruct := &createuser.CreateUserInput{
			Username: "user1",
			Password: "1234",
		}
		jsonData, _ := json.Marshal(inputStruct)
		ctx.Request = httptest.NewRequest("POST", "/create", bytes.NewBuffer(jsonData))

		//When
		CreateUserHandler(ctx)

		//Then
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("should return 400 status code when invalid input is provided", func(t *testing.T) {
		//Given
		setUpTestCreateUserHandler(t)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest("POST", "/create", bytes.NewBuffer([]byte{}))

		//When
		CreateUserHandler(ctx)

		//Then
		assert.Equal(t, 400, w.Code)
	})

	t.Run("should return 500 status code when saveUserFn returns error", func(t *testing.T) {
		//Given
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		inputStruct := &createuser.CreateUserInput{
			Username: "user1",
			Password: "1234",
		}
		jsonData, _ := json.Marshal(inputStruct)
		ctx.Request = httptest.NewRequest("POST", "/create", bytes.NewBuffer(jsonData))
		saveUserFn = func(u *entity.User) (*entity.User, error) {
			return nil, assert.AnError
		}

		//When
		CreateUserHandler(ctx)

		//Then
		assert.Equal(t, 500, w.Code)
	})

}
