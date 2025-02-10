package userrespository

import (
	"database/sql"
	"test_case_putri/config"
	"test_case_putri/models"
	"time"
)

func GetUsersRepository() ([]models.UserResponse, error) {
	queryGet, err := config.DbConn.MySql.Prepare("SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer queryGet.Close()

	rows, err := queryGet.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var u models.UserResponse
		err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUserByIdRepository(Id int) (*models.UserResponse, error) {
	queryGet, err := config.DbConn.MySql.Prepare("SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer queryGet.Close()

	row := queryGet.QueryRow(Id)

	var user models.UserResponse
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.UserResponse{}, err
		}
		return nil, err
	}

	return &user, nil
}

func InsertUserRepository(request models.UserRequest) (int64, error) {
	queryInsert, err := config.DbConn.MySql.Prepare("INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer queryInsert.Close()

	res, err := queryInsert.Exec(request.Name, request.Email, time.Now())
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func UpdateUserRepository(Id int, request models.UserRequest) (int64, error) {
	result, err := config.DbConn.MySql.Exec(
		"UPDATE users SET name = ?, email = ?, updated_at = NOW() WHERE id = ?",
		request.Name, request.Email, Id,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func DeleteUserRepository(Id int) (int64, error) {
	result, err := config.DbConn.MySql.Exec(
		"DELETE FROM users WHERE id = ?",
		Id,
	)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
