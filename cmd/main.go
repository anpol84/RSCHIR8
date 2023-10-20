package main

import (
	"awesomeProject/web/api"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Получение порта из аргументов командной строки
	port := flag.String("port", "8080", "port to run the server on")
	name := flag.String("name", "name", "")
	flag.Parse()

	// Открытие файла для логирования
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла логов: %v", err)
	}

	// Установка вывода логов в файл и консоль
	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		api.HandleData(w, r, *name) // передать имя cookie "name"
	})
	http.HandleFunc("/api/getData", func(w http.ResponseWriter, r *http.Request) {
		api.GetData(w, r, *name) // передать имя cookie "name"
	})
	http.HandleFunc("/api/linearGet", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second) // имитация долгого вычисления
		api.GetData(w, r, *name)    // передать имя cookie "name"
	})

	http.HandleFunc("/api/concurrentGet", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			time.Sleep(5 * time.Second) // имитация долгого вычисления
			api.GetData(w, r, *name)    // передать имя cookie "name"
		}()
	})
	logger.Printf("Сервер запущен на порту %s", *port)
	logger.Fatal(http.ListenAndServe(":"+*port, nil))

}
