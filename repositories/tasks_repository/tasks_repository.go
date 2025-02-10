package tasksrepository

import (
	"database/sql"
	"test_case_putri/config"
	"test_case_putri/models"
	"time"
)

func GetTasksRepository() ([]models.TaskResponse, error) {
	queryGet, err := config.DbConn.MySql.Prepare("SELECT id, title, description, status, user_id, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer queryGet.Close()

	rows, err := queryGet.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.TaskResponse
	for rows.Next() {
		var u models.TaskResponse
		err := rows.Scan(&u.Id, &u.Title, &u.Description, &u.Status, &u.UserId, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetTaskByIdRepository(Id int) (*models.TaskResponse, error) {
	queryGet, err := config.DbConn.MySql.Prepare("SELECT id, title, description, status, user_id, created_at, updated_at FROM tasks WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer queryGet.Close()

	row := queryGet.QueryRow(Id)

	var task models.TaskResponse
	err = row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.UserId, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.TaskResponse{}, err
		}
		return nil, err
	}

	return &task, nil
}

func GetTaskByUserIdRepository(UserId int) ([]models.TaskResponse, error) {
	queryGet, err := config.DbConn.MySql.Prepare("SELECT id, title, description, status, user_id, created_at, updated_at FROM tasks WHERE user_id = ?")
	if err != nil {
		return nil, err
	}
	defer queryGet.Close()

	rows, err := queryGet.Query(UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.TaskResponse
	for rows.Next() {
		var u models.TaskResponse
		err := rows.Scan(&u.Id, &u.Title, &u.Description, &u.Status, &u.UserId, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

func InsertTaskRepository(request models.TaskRequest) (int64, error) {
	queryInsert, err := config.DbConn.MySql.Prepare("INSERT INTO tasks (title, description, status, user_id, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer queryInsert.Close()

	res, err := queryInsert.Exec(request.Title, request.Description, request.Status, request.UserId, time.Now())
	if err != nil {
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func UpdateTaskRepository(Id int, request models.TaskRequest) (int64, error) {
	result, err := config.DbConn.MySql.Exec(
		"UPDATE tasks SET title = ?, description = ?, status = ?, user_id = ?, updated_at = NOW() WHERE id = ?",
		request.Title, request.Description, request.Status, request.UserId, Id,
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

func DeleteTaskRepository(Id int) (int64, error) {
	result, err := config.DbConn.MySql.Exec(
		"DELETE FROM tasks WHERE id = ?",
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