package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"jwt/models"
	"jwt/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Signup is aaa
func (c Controller) Signup(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := models.User{}

		json.NewDecoder(req.Body).Decode(&user)

		if user.Email == "" {
			utils.ResponseError(w, http.StatusBadRequest, "Email Must be filled in")
		}

		if user.Password == "" {
			utils.ResponseError(w, http.StatusBadRequest, "Password Must be filled in")
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			log.Fatal(err)
		}

		user.Password = string(hash)

		query := "INSERT INTO user (email, password) VALUES (?, ?)"
		resp, err := db.Exec(query, user.Email, user.Password)
		if err != nil {
			log.Fatal(err)
		}

		id, err := resp.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		user.ID = int(id)

		utils.ResponseSuccess(w, user)
	}
}

// Login is aaa
func (c Controller) Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user models.User
		var jwt models.JWT

		json.NewDecoder(req.Body).Decode(&user)

		if user.Email == "" {
			utils.ResponseError(w, http.StatusBadRequest, "Email is Missing")
			return
		}

		if user.Password == "" {
			utils.ResponseError(w, http.StatusBadRequest, "Password is Misssing")
			return
		}

		password := user.Password

		query := "SELECT password FROM user WHERE email = ?"
		row := db.QueryRow(query, user.Email)
		err := row.Scan(&user.Password)

		hashedPassword := user.Password

		if err != nil {
			if err == sql.ErrNoRows {
				utils.ResponseError(w, http.StatusBadRequest, "User Does Not Exist")
				return
			}
			log.Fatal(err)
		}

		isValidPass := utils.ComparePassword(hashedPassword, []byte(password))

		if isValidPass {
			token, err := utils.GenerateToken(user)
			if err != nil {
				log.Fatal(err)
			}

			jwt.Token = token
			utils.ResponseSuccess(w, jwt)
		} else {
			utils.ResponseError(w, http.StatusBadRequest, "Invalid Password")
			return
		}
	}
}

// TokenVerifyMiddleware is a route that called when make a request
func (c Controller) TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte(os.Getenv("SECRET")), nil
			})

			if err != nil {
				utils.ResponseError(w, http.StatusUnauthorized, err.Error())
				return
			}

			if token.Valid {
				next.ServeHTTP(w, req)
			} else {
				utils.ResponseError(w, http.StatusUnauthorized, err.Error())
				return
			}
		} else {
			utils.ResponseError(w, http.StatusUnauthorized, "Invalid Token")
			return
		}
	})
}
