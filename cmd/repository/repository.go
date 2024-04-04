package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type UserData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Application struct {
	Database *pgx.Conn
	Memory   []*UserData
}

func (app *Application) SaveData(userData *UserData) error {
	// Сохранение в базу данных
	if _, err := app.Database.Exec(
		context.Background(),
		"INSERT INTO users (name, age) VALUES ($1, $2)",
		userData.Name,
		userData.Age,
	); err != nil {
		return fmt.Errorf("failed to save data to database: %w", err)
	}

	// Сохранение в память приложения
	app.Memory = append(app.Memory, userData)

	return nil
}

func (app *Application) GetDataByID(id int) (*UserData, error) {
	var userData UserData
	err := app.Database.QueryRow(
		context.Background(),
		"SELECT id, name, age FROM users WHERE id = $1",
		id,
	).Scan(&userData.ID, &userData.Name, &userData.Age)
	if err != nil {
		return nil, fmt.Errorf("failed to get data from database: %w", err)
	}

	return &userData, nil
}

func (app *Application) GetDataFromMemory() {
	fmt.Println("Data saved in memory:")
	for _, userData := range app.Memory {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", userData.ID, userData.Name, userData.Age)
	}
}
