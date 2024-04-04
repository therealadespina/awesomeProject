package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ApplicationConfig struct {
	AwesomeProject struct {
		DataSource string `yaml:"dataSourcePath"`
	} `yaml:"awesomeProject"`
}

func DataSource() (*Application, error) {
	// Чтение файла application.yaml
	yamlFile, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		log.Fatalf("Ошибка чтения файла application.yaml: %v", err)
	}

	// Создание экземпляра структуры для хранения конфигурации
	var appConfig ApplicationConfig

	// Разбор YAML данных
	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		log.Fatalf("Ошибка разбора YAML: %v", err)
	}

	// Использование извлеченных переменных
	fmt.Println("DataSourcePath is:", appConfig.AwesomeProject.DataSource)

	if err != nil {
		log.Fatalf("Ошибка чтения файла YAML: %v", err)
	}
	conn, err := pgx.Connect(context.Background(), appConfig.AwesomeProject.DataSource)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Application{
		Database: conn,
		Memory:   make([]*UserData, 0),
	}, nil
}
