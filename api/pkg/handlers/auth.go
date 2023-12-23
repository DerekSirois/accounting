package handlers

import (
	"accounting/pkg/auth"
	"accounting/pkg/models"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("JWT_SECRET")

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			var jwtToken = request.Header["Token"][0]
			var userClaim auth.UserClaim
			token, err := jwt.ParseWithClaims(jwtToken, &userClaim, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})
			if err != nil {
				errorJson(writer, err, http.StatusBadRequest)
				return
			}
			if !token.Valid {
				errorJson(writer, fmt.Errorf("invalid token"), http.StatusBadRequest)
				return
			}
			endpointHandler(writer, request)
		} else {

			errorJson(writer, fmt.Errorf("missing token"), http.StatusBadRequest)
			return
		}
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.UsersRegister
	err := readJson(r, &u)
	if err != nil {
		errorJson(w, err, http.StatusBadRequest)
		return
	}

	u.Password, err = auth.HashPassword(u.Password)
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}

	err = u.Create()
	if err != nil {
		errorJson(w, fmt.Errorf("can't create the user in the database"), http.StatusInternalServerError)
		return
	}

	res := JsonResponse{
		Error:   false,
		Message: "User created successfully",
	}

	writeJson(w, http.StatusOK, res)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var u models.UsersLogin
	err := readJson(r, &u)
	if err != nil {
		errorJson(w, err, http.StatusBadRequest)
		return
	}

	var uDb models.Users
	err = uDb.GetByEmail(u.Email)
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}

	if !auth.CheckPasswordHash(u.Password, uDb.Password) {
		errorJson(w, fmt.Errorf("wrong password"), http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateJWTToken(uDb.Id, uDb.Email)
	if err != nil {
		errorJson(w, fmt.Errorf("can't generate token"), http.StatusInternalServerError)
		return
	}

	res := JsonResponse{
		Error:   false,
		Message: token,
	}

	writeJson(w, http.StatusOK, res)
}
