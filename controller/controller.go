package controller

import (
	"catalogService/model"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)


type application struct {
	auth struct {
		username string
		password string
	}
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}


func InitController(r *mux.Router, serviceDb model.IServiceDb) {

	app := new(application)
	app.auth.username = "pankaj"
	app.auth.password = "1234"
	r.HandleFunc("/restaurants", app.basicAuth(CreateRestaurant(serviceDb))).Methods("POST")
	r.HandleFunc("/restaurants/{id}", app.basicAuth(AddMenuItems(serviceDb))).Methods("POST")
	r.HandleFunc("/restaurants/{id}", app.basicAuth(GetRestaurant(serviceDb))).Methods("GET")
	r.HandleFunc("/restaurants", app.basicAuth(GetAllRestaurants(serviceDb))).Methods("GET")
	
	
}

func CreateRestaurant(serviceDb model.IServiceDb) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {


		w.Header().Set("Content-Type", "application/json")
		var restaurants model.Restaurants
		json.NewDecoder(r.Body).Decode(&restaurants)

		restaurant,err := serviceDb.CreateRestaurant(restaurants)

		if err!=nil{
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(restaurant)

	}
}


func GetRestaurant(serviceDb model.IServiceDb) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {


		w.Header().Set("Content-Type", "application/json")
		param := mux.Vars(r)

		userId, err := strconv.Atoi(param["id"])
		if err != nil {
			json.NewEncoder(w).Encode("Invalid user ID")
			return
		}


		restaurant,err := serviceDb.GetRestaurant(userId)

		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(restaurant)

	}
}

func AddMenuItems(serviceDb model.IServiceDb) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {


		w.Header().Set("Content-Type", "application/json")
		param := mux.Vars(r)

		userId, err := strconv.Atoi(param["id"])
		if err != nil {
			json.NewEncoder(w).Encode("Invalid user ID")
			return
		}

		var items model.MenuItems
		json.NewDecoder(r.Body).Decode(&items)
		menuItem,err := serviceDb.AddMenuItems(userId,items)

		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode(menuItem)

	}
}

func GetAllRestaurants(serviceDb model.IServiceDb) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {


		w.Header().Set("Content-Type", "application/json")

		restaurants,err := serviceDb.GetAllRestaurants()

		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewEncoder(w).Encode(restaurants)

	}
}