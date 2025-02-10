package userservice

import (
	"test_case_putri/models"
	userrespository "test_case_putri/repositories/user_respository"
)

func GetAllUserService() ([]models.UserResponse, error) {
	dataUser, err := userrespository.GetAllUserRepository()
	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

func GetUserByIdService(Id int) (*models.UserResponse, error) {
	dataUser, err := userrespository.GetUserByIdRepository(Id)
	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

func InsertUserService(request models.UserRequest) (int64, error) {
	insertUser, err := userrespository.InsertUserRepository(request)
	if err != nil {
		return insertUser, err
	}

	return insertUser, nil
}

func UpdateUserService(Id int, request models.UserRequest) (int64, error) {
	insertUser, err := userrespository.UpdateUserRepository(Id, request)
	if err != nil {
		return insertUser, err
	}

	return insertUser, nil
}

func DeleteUserService(Id int) (int64, error) {
	insertUser, err := userrespository.DeleteUserRepository(Id)
	if err != nil {
		return insertUser, err
	}

	return insertUser, nil
}
