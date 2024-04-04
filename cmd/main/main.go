package main

import (
	"awesomeProject/cmd/handlers"
	. "awesomeProject/cmd/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nats-io/nats.go"
)

const topic = "random-data"

func main() {
	app, err := DataSource()
	if err != nil {
		log.Fatalf("Failed to create main: %v", err)
	}

	// Подключение к серверу NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Подписка на тему "random-data" и обработка сообщений из NATS
	nc.Subscribe(topic, func(msg *nats.Msg) {
		var userData UserData
		if err := json.Unmarshal(msg.Data, &userData); err != nil {
			log.Printf("Failed to unmarshal data: %v\n", err)
			return
		}

		// Сохранение данных
		if err := app.SaveData(&userData); err != nil {
			log.Printf("Failed to save data: %v\n", err)
			return
		}

		// Вывод данных, сохраненных в памяти приложения, на экран
		app.GetDataFromMemory()
	})

	router := mux.NewRouter()
	router.HandleFunc("/data", handlers.HandleJSONData(app)).Methods(http.MethodPost)
	router.HandleFunc("/data/{id}", handlers.HandleGetDataByID(app)).Methods(http.MethodGet)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
