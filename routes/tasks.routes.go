package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-gorm-rest-api/db"
	"go-gorm-rest-api/models"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, req *http.Request) {
	var tasks []models.Task
	w.Header().Set("Content-Type", "application/json")

	db.DB.Find(&tasks)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task models.Task
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(req.Body).Decode(&task)
	createdTask := db.DB.Create(&task)
	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(createdTask.Error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&task)
}

func GetTaskHandler(w http.ResponseWriter, req *http.Request) {

	var task models.Task
	params := mux.Vars(req)

	db.DB.First(&task, params["id"])

	w.Header().Set("Content-Type", "application/json")

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&task)
}

func DeleteTaskHandler(w http.ResponseWriter, req *http.Request) {
	var task models.Task
	params := mux.Vars(req)

	db.DB.First(&task, params["id"])

	w.Header().Set("Content-Type", "application/json")

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Task not found")
		return
	}

	db.DB.Delete(&task)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(&task)
}
