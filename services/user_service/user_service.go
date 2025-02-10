package userservice

import (
	"errors"

	"test_case_putri/constants"
	"test_case_putri/models"
	userrespository "test_case_putri/repositories/user_respository"
	"test_case_putri/utilities"
)

func GetUsersService() ([]models.UserResponse, error) {
	dataUser, err := userrespository.GetUsersRepository()
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
	tokenString, _ := utilities.HashPassword(request.Password)
	reqUser := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Token:    tokenString,
	}

	insertUser, err := userrespository.InsertUserRepository(reqUser)
	if err != nil {
		return insertUser, err
	}

	return insertUser, nil
}

func UpdateUserService(Id int, request models.UserRequest) (int64, error) {
	user, err := userrespository.GetUserByIdRepository(Id)
	if err != nil || user.Id == 0 {
		return 0, errors.New("invalid id")
	}

	updateUser, err := userrespository.UpdateUserRepository(Id, request)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil
}

func DeleteUserService(Id int) (int64, error) {
	user, err := userrespository.GetUserByIdRepository(Id)
	if err != nil || user.Id == 0 {
		return 0, errors.New("invalid id")
	}

	deleteUser, err := userrespository.DeleteUserRepository(Id)
	if err != nil {
		return deleteUser, err
	}

	return deleteUser, nil
}

func Login(request models.Login) (*models.GenerateJWT, error) {
	checkUser, err := userrespository.GetUserByEmailOrPasswordRepository(request.Email, request.Password)
	if err != nil || checkUser.Id == 0 {
		return nil, errors.New(constants.ErrFailedAuthentication)
	}

	if checkUser.IsLogin {
		return nil, errors.New(constants.ErrLogin)
	}	

	login := models.GenerateJWT{
		UserId: checkUser.Id,
		Name:   checkUser.Name,
		Email:  checkUser.Email,
	}

	tokenString, _, err := utilities.GenerateJWT(&login)
	if err != nil {
		return nil, err
	}

	login.Token = tokenString
	
	_, err = userrespository.UpdateUserLoginRepository(checkUser.Id, true, &tokenString)
	if err != nil {
		return nil, err
	}

	return &login, nil
}

func Logout(request models.Login) (int64, error) {
	checkUser, err := userrespository.GetUserByEmailOrPasswordRepository(request.Email, request.Password)
	if err != nil || checkUser.Id == 0 {
		return 0, errors.New(constants.ErrLogout)
	}

	if !checkUser.IsLogin {
		return 0, errors.New(constants.ErrLogout)
	}	

	_, err = userrespository.UpdateUserLoginRepository(checkUser.Id, false, nil)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
