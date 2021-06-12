package main

import (
	"encoding/json"
	"golang-docker/db"
	"golang-docker/model"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// var c context.Context

	err, db := db.ConnDB()
	if err != nil {
		os.Exit(1)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("port is required")
	}
	instanceID := os.Getenv("INSTANCE_ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		text := "Hello world "
		if text != "" {
			text += " from" + instanceID
		}

		w.Write([]byte(text))
	})

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		var user []model.User
		db.Model(&user).Scan(&user)

		listUser, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte("error occured while getting user lists"))
		}

		w.Write([]byte(listUser))

	})

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		var tasks []model.Task
		db.Model(&tasks).Scan(&tasks)

		listTasks, err := json.Marshal(tasks)
		if err != nil {
			w.Write([]byte("error occured while getting user lists"))
		}

		w.Write([]byte(listTasks))

	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0:" + port

	log.Println("server started at ", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
