package userhandler

import (
	"net/http"
	"strconv"
	"test_case_putri/constants"
	"test_case_putri/models"
	userservice "test_case_putri/services/user_service"
	"test_case_putri/utilities"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(c *gin.Context) {
	data, err := userservice.GetUsersService()

	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, nil)
}

func GetUserByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	data, err := userservice.GetUserByIdService(id)

	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, nil)
}

func InsertUserHandler(c *gin.Context) {
	var request models.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	data, err := userservice.InsertUserService(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
}

func UpdateUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	}

	var request models.UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	}

	data, err := userservice.UpdateUserService(id, request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData, err)
}

func DeleteUserHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDeleteData, err)
		return
	}

	data, err := userservice.DeleteUserService(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDeleteData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDeleteData, err)
}

func LoginHandler(c *gin.Context) {
	var request models.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrFailedAuthentication, err)
		return
	}

	if request.Email == "" || request.Password == "" {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrFailedAuthentication, nil)
		return
	}

	data, err := userservice.Login(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrFailedAuthentication, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessLogin, err)
}

func LogoutHandler(c *gin.Context) {
	var request models.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrLogout, err)
		return
	}

	if request.Email == "" || request.Password == "" {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrFailedAuthentication, nil)
		return
	}

	data, err := userservice.Logout(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.ErrFailedAuthentication, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessLogout, err)
}