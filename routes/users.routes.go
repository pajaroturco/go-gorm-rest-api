package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-gorm-rest-api/db"
	"go-gorm-rest-api/models"
	"net/http"
)

func GetUsersHandler(w http.ResponseWriter, req *http.Request) {
	var users []models.User
	w.Header().Set("Content-Type", "application/json")

	db.DB.Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, req *http.Request) {

	var user models.User
	params := mux.Vars(req)

	db.DB.First(&user, params["id"])

	w.Header().Set("Content-Type", "application/json")

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	json.NewDecoder(req.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(createdUser.Error.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, req *http.Request) {
	var user models.User
	params := mux.Vars(req)

	db.DB.First(&user, params["id"])

	w.Header().Set("Content-Type", "application/json")

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	db.DB.Delete(&user)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(&user)
}
