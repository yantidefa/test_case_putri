package taskhandler

import (
	"net/http"
	"strconv"
	"test_case_putri/constants"
	"test_case_putri/models"
	taskservice "test_case_putri/services/task_service"
	"test_case_putri/utilities"

	"github.com/gin-gonic/gin"
)

func GetTasksHandler(c *gin.Context) {
	data, err := taskservice.GetTasksService()

	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, nil)
}

func GetTaskByIdHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	data, err := taskservice.GetTaskByIdService(id)

	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, nil)
}

func GetTaskByUserIdHandler(c *gin.Context) {
	idParam := c.Param("user_id")
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	data, err := taskservice.GetTaskByUserIdService(userId)

	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDisplayedData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDisplayedData, nil)
}

func InsertTaskHandler(c *gin.Context) {
	var request models.TaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	if request.Status == "" {
		request.Status = "pending"
	}

	validStatus := map[string]bool{
		"pending":     true,
		"in_progress": true,
		"completed":   true,
	}

	if !validStatus[request.Status] {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.InvalidStatusValue, nil)
		return
	}

	data, err := taskservice.InsertTaskService(request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedAddData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessAddData, err)
}

func UpdateTaskHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	}

	var request models.TaskRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	}

	data, err := taskservice.UpdateTaskService(id, request)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedUpdateData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessUpdateData, err)
}

func DeleteTaskHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDeleteData, err)
		return
	}

	data, err := taskservice.DeleteTaskService(id)
	if err != nil {
		utilities.SetResponseJSON(c, http.StatusBadRequest, nil, constants.FailedDeleteData, err)
		return
	}

	utilities.SetResponseJSON(c, http.StatusOK, data, constants.SuccessDeleteData, err)
}