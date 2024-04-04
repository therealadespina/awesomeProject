package handlers

import (
	. "awesomeProject/cmd/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func HandleJSONData(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var userData UserData
		if err := decoder.Decode(&userData); err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// Сохранение данных
		if err := app.SaveData(&userData); err != nil {
			http.Error(w, "Failed to save data", http.StatusInternalServerError)
			return
		}

		// Вывод данных, сохраненных в памяти приложения, на экран
		app.GetDataFromMemory()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Data received and saved successfully"))
	}
}

func HandleGetDataByID(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		userData, err := app.GetDataByID(id)
		if err != nil {
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(userData)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "main/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
