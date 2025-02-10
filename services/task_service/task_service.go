package taskservice

import (
	"errors"
	"test_case_putri/models"
	tasksrepository "test_case_putri/repositories/tasks_repository"
	userrespository "test_case_putri/repositories/user_respository"
)

func GetTasksService() ([]models.TaskResponse, error) {
	dataTask, err := tasksrepository.GetTasksRepository()
	if err != nil {
		return nil, err
	}

	return dataTask, nil
}

func GetTaskByIdService(Id int) (*models.TaskResponse, error) {
	dataTask, err := tasksrepository.GetTaskByIdRepository(Id)
	if err != nil {
		return nil, err
	}

	return dataTask, nil
}

func GetTaskByUserIdService(UserId int) ([]models.TaskResponse, error) {
	dataTask, err := tasksrepository.GetTaskByUserIdRepository(UserId)
	if err != nil {
		return nil, err
	}

	return dataTask, nil
}

func InsertTaskService(request models.TaskRequest) (int64, error) {
	user, err := userrespository.GetUserByIdRepository(request.UserId)
	if err != nil || user.Id == 0 {
		return 0, errors.New("invalid user id")
	}

	insertTask, err := tasksrepository.InsertTaskRepository(request)
	if err != nil {
		return 0, err
	}

	return insertTask, nil
}

func UpdateTaskService(Id int, request models.TaskRequest) (int64, error) {
	dataTask, err := tasksrepository.GetTaskByIdRepository(Id)
	if err != nil || dataTask.Id == 0 {
		return 0, errors.New("invalid id")
	}

	user, err := userrespository.GetUserByIdRepository(request.UserId)
	if err != nil || user.Id == 0 {
		return 0, errors.New("invalid user id")
	}

	updateTask, err := tasksrepository.UpdateTaskRepository(Id, request)
	if err != nil {
		return 0, err
	}

	return updateTask, nil
}

func DeleteTaskService(Id int) (int64, error) {
	dataTask, err := tasksrepository.GetTaskByIdRepository(Id)
	if err != nil || dataTask.Id == 0 {
		return 0, errors.New("invalid id")
	}
	
	deleteTask, err := tasksrepository.DeleteTaskRepository(Id)
	if err != nil {
		return deleteTask, err
	}

	return deleteTask, nil
}