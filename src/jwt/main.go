package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// User is representation for error response
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Error is representation for error response
type Error struct {
	Message string `json:"message"`
}

// JWT is representation token jwt response
type JWT struct {
	Token string `json:"token"`
}

var db *sql.DB

func connect() *sql.DB {
	connString := "root:erwindo123@tcp(localhost:3306)/jwt"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	port := "8000"
	db = connect()
	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/signup", signup).Methods(http.MethodPost)
	router.HandleFunc("/login", login).Methods(http.MethodPost)
	router.HandleFunc("/user", TokenVerifyMiddleware(protectedEndpoint)).Methods(http.MethodGet)

	fmt.Println("Starting Web Server at Port : " + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal()
	}
}

func respondWithError(w http.ResponseWriter, status int, err Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func signup(w http.ResponseWriter, req *http.Request) {
	user := User{}
	error := Error{}

	json.NewDecoder(req.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email Must be filled in"
		respondWithError(w, http.StatusBadRequest, error)
	}

	if user.Password == "" {
		error.Message = "Password Must be filled in"
		respondWithError(w, http.StatusBadRequest, error)
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

	responseJSON(w, user)
}

func generateToken(user User) (string, error) {
	var err error
	secret := "secret"

	validTime := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"valid": validTime,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func login(w http.ResponseWriter, req *http.Request) {
	var user User
	var jwt JWT
	var error Error

	json.NewDecoder(req.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is Missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is Misssing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	password := user.Password

	query := "SELECT password FROM user WHERE email = ?"
	row := db.QueryRow(query, user.Email)
	err := row.Scan(&user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			error.Message = "The User Does Not Exist"
			respondWithError(w, http.StatusBadRequest, error)
			return
		}
		log.Fatal(err)
	}

	hashedPassword := user.Password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		error.Message = "Invalid Password"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	token, err := generateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	jwt.Token = token

	responseJSON(w, jwt)
}

func protectedEndpoint(w http.ResponseWriter, req *http.Request) {
	user := User{
		ID:       1,
		Email:    "This is a email",
		Password: "This is a password",
	}

	responseJSON(w, user)

	authHeader := req.Header.Get("Authorization")
	fmt.Println(authHeader)
}

// TokenVerifyMiddleware is a route that called when make a request
func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var errorObject Error
		authHeader := req.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return []byte("secret"), nil
			})

			if err != nil {
				errorObject.Message = err.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}

			if token.Valid {
				next.ServeHTTP(w, req)
			} else {
				errorObject.Message = err.Error()
				respondWithError(w, http.StatusUnauthorized, errorObject)
				return
			}
		} else {
			errorObject.Message = "Invalid Token"
			respondWithError(w, http.StatusUnauthorized, errorObject)
			return
		}
	})
}
