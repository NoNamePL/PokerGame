package app

import (
	"awesomeProject/models"
	u "awesomeProject/utils"
	"go/token"
	"net/http"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"api/user/new", "api/user/login"} // список эндпоинтов, для которых не требуется авторизация
		requestPath := r.URL.Path                             // текущий путь запроса

		//проверяем, не требуется ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokerHeader := r.Header.Get("Authorization") // получение токена

		if tokerHeader == "" { // Токен отсутствует, возвращаем 403 http-код Unauthorized
			response = u.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Response(w, response)
			return
		}

		splitted := strings.Split(tokerHeader, " ") // Токен обычно поставляется в формате `Bearer {token-body}`,мы проверяем, соответствует ли полученный токен этому требованию
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			return
		}

		tokenPart := splitted[1] // получаем вторую часть токена
		tk := &models.Token{}

	})
}
