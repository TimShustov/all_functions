package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Метод Handler. Данный метод будет обрабатывать HTTP запросы поступающие к функции
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	// Чтение тела запроса
	body, _ := ioutil.ReadAll(r.Body)

	// Логирование входящего запроса
	log.Printf("Request received: %s\nMethod: %s", string(body), r.Method)

	// Подготовка и возврат ответа на вызов
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello from Go function!\nYou said: %s", string(body))))
}
