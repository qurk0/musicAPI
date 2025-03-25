package request

import (
	"musicLib/pkg/responce"
	"net/http"
)

/*
Метод для парсинга тела обработки:
-- Нам приходит запрос с некоторыми данными
-- Мы его парсим
-- В случае ошибки парсинга метод автоматически кидает респонс с кодом 400
-- Мы его проверяем на валидность (Тэги валидности указаны в payload.go соответствующей модели в БД)
-- В случае невалидности метод автоматически кидает респонс с кодом 400
-- Мы возвращаем из метода тело запроса и нулевую ошибку
*/
func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		responce.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = Valid(body)
	if err != nil {
		responce.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}
