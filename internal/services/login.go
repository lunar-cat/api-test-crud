package services

import (
	"api-test-crud/config"
	"api-test-crud/internal/models"
	"api-test-crud/internal/utils"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type LoginHandler struct{}

func (h *LoginHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	userIsValid := parseUser(&user, r.Body)
	if !userIsValid {
		utils.BadRequestResponse(w, r, "Error: Body formado incorrectamente")
		return
	}

	foundUser, ok := models.SearchUser(&user)
	if !ok {
		utils.BadRequestResponse(w, r, "Error: Usuario o contraseña incorrectas")
		return
	}

	token, err := generateToken(foundUser)
	if err != nil {
		utils.InternalServerErrorResponse(w, r, "Error: Problema al generar el token")
		return
	}

	data := map[string]any{
		"token": token,
		"user": map[string]string{
			"name":  foundUser.Name,
			"email": foundUser.Email,
		},
	}
	utils.OkResponse(w, r, "Inicio de sesión correcto", data)
}

// Helpers

func parseUser(user *models.User, body io.ReadCloser) bool {
	if body == nil {
		return false
	}

	if err := json.NewDecoder(body).Decode(user); err != nil {
		return false
	}

	payloadValid := isLoginPayloadValid(user)

	return payloadValid
}

func isLoginPayloadValid(user *models.User) bool {
	if user.Username == "" {
		return false
	}

	if user.Password == "" {
		return false
	}

	return true
}

func generateToken(user *models.User) (string, error) {
	hours, errEnv := strconv.Atoi(os.Getenv("TOKEN_DURATION_HOURS"))
	if errEnv != nil {
		return "", errEnv
	}

	tokenExpiry := time.Now().Add(time.Duration(hours) * time.Hour).Unix()
	claims := map[string]interface{}{
		"id":   user.ID,
		"name": user.Username,
		"exp":  tokenExpiry,
	}
	_, tokenString, errToken := config.TokenAuth.Encode(claims)
	return tokenString, errToken
}
