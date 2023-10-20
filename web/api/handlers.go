package api

import (
	"encoding/json"
	"github.com/gorilla/securecookie"
	"io"
	"log"
	"net/http"
	"os"
)

var hashKey = securecookie.GenerateRandomKey(32)
var blockKey = securecookie.GenerateRandomKey(32)
var s = securecookie.New(hashKey, blockKey)
var logFile, _ = os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

// Установка вывода логов в файл и консоль
var logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)

type requestData struct {
	Query string `json:"query"`
}

type responseData struct {
	Response string `json:"response"`
}

func HandleData(w http.ResponseWriter, r *http.Request, name string) {
	logger.Println("sad")
	logger.Println(name)
	var data requestData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := "server " + data.Query

	value := map[string]string{
		"query":    data.Query,
		"response": response,
	}

	encoded, err := s.Encode(name, value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:  name,
		Value: encoded,
		Path:  "/",
	}

	http.SetCookie(w, &cookie)
	//json.NewEncoder(w).Encode(responseData{Response: response})

	// Log the request to console
	logger.Println("Handled data:", data, "Response:", response)
}

func GetData(w http.ResponseWriter, r *http.Request, name string) {
	cookie, err := r.Cookie(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := make(map[string]string)
	if err := s.Decode(name, cookie.Value, &value); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := value["response"]

	json.NewEncoder(w).Encode(responseData{Response: response})

	// Log the request to console
	logger.Println("Retrieved data:", value)
}
