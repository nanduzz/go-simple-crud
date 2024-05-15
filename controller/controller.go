package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createuser "github.com/nanduzz/go-simple-crud/use_case/create_user"
	findall "github.com/nanduzz/go-simple-crud/use_case/find_all"
	finduser "github.com/nanduzz/go-simple-crud/use_case/find_user"
	"github.com/nanduzz/go-simple-crud/util"
)

var findAllUsersFn findall.FindAllUsersFn
var findUserByIdFn finduser.FindUserByIdFn
var saveUserFn createuser.SaveFunc

func Initialize(
	findAllUsersFnInjected findall.FindAllUsersFn,
	findUserByIdFnInjected finduser.FindUserByIdFn,
	saveUserFnInjected createuser.SaveFunc,
) {
	findAllUsersFn = findAllUsersFnInjected
	findUserByIdFn = findUserByIdFnInjected
	saveUserFn = saveUserFnInjected
}

func CreateUserHandler(c *gin.Context) {
	var input createuser.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	output, err := createuser.Execute(&input, saveUserFn)
	if err != nil {
		// log error
		util.InfoLog.Printf("Encountered error in main: %+v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unexpected error while creating user, try again later",
		})
		return
	}

	util.InfoLog.Printf("creating user with username: %s\n", input.Username)
	c.Header("Location", "/users/"+output.ID)
	c.JSON(http.StatusCreated, output)
}

func FindAll(c *gin.Context) {
	util.InfoLog.Println("Finding All Users")
	users, err := findall.Execute(findAllUsersFn)
	if err != nil {
		// log error
		util.WarnrLog.Printf("Encountered error: %+v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unexpected error while fetching users, try again later",
		})
		return
	}

	c.JSON(http.StatusOK, users)

}

func FindUserHandlerById(c *gin.Context) {
	if findUserByIdFn == nil {
		util.ErrorLog.Println("findUserByIdFn not set")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while fetching user, try again later",
		})
		return
	}
	id := c.Param("id")

	user, err := finduser.Execute(finduser.Input{Id: id}, findUserByIdFn)
	if err != nil {
		if finduser.ErrFindUserByIdInputIdRequired.Is(err) {
			util.WarnrLog.Println("id is required")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}
