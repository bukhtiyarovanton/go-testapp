package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/bukhtiyarovanton/go-testapp/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

//auth middleware
func Validate(protectedRoute http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			http.NotFound(w, r)
			return
		}

		token, err := jwt.ParseWithClaims(authHeader, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("go-testappsecret"), nil
		})
		if err != nil {
			http.NotFound(w, r)
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "UserClaims", *claims)
			protectedRoute(w, r.WithContext(ctx))
		} else {
			http.NotFound(w, r)
			return
		}
	})
}

// auth routes handlers
type Claims struct {
	UserID   uint   `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	userName := vars.Get("username")
	if len(userName) == 0 {
		// return json with "username is mandatory"
		WriteError(w, http.StatusBadRequest, "username is mandatory")
		return
	}

	password := vars.Get("password")
	if len(password) == 0 {
		// return json with "password is mandatory"
		WriteError(w, http.StatusBadRequest, "username is mandatory")
		return
	}

	// find user by name in db
	user, ok := UserRepoFindByName(userName)
	if ok != nil {
		WriteError(w, http.StatusNotFound, fmt.Sprintf("User with name %v does not exist.", userName))
		return
	}

	if ok := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); ok != nil {
		WriteError(w, http.StatusBadRequest, ok.Error())
		return
	}

	expireToken := time.Now().Add(time.Hour * 1).Unix()

	claims := Claims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:3000",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("go-testappsecret"))

	result := struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}{
		"bearer",
		signedToken,
	}

	WriteResponse(w, http.StatusOK, result)
}

// Todo routes handles
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("UserClaims").(Claims)
	if !ok {
		http.NotFound(w, r)
		return
	}

	WriteResponse(w, http.StatusOK, TodoRepoGetAllForUser(claims.UserID))
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("UserClaims").(Claims)
	if !ok {
		http.NotFound(w, r)
		return
	}

	vars := mux.Vars(r)

	var todoId uint64
	var err error

	if todoId, err = strconv.ParseUint(vars["todoId"], 0, 0); err != nil {
		panic(err)
	}

	todo := TodoRepoFindForUser(claims.UserID, uint(todoId))
	if todo.ID > 0 {
		WriteResponse(w, http.StatusOK, todo)
		return
	}

	WriteError(w, http.StatusNotFound, "Not Found")
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("UserClaims").(Claims)
	if !ok {
		http.NotFound(w, r)
		return
	}

	var todo models.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		WriteError(w, 422, err.Error())
	}

	t := TodoRepoCreateForUser(claims.UserID, todo)
	WriteResponse(w, http.StatusCreated, t)
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value("UserClaims").(Claims)
	if !ok {
		http.NotFound(w, r)
		return
	}

	vars := mux.Vars(r)

	var todoId uint64
	var err error

	if todoId, err = strconv.ParseUint(vars["todoId"], 0, 0); err != nil {
		panic(err)
	}

	todo := TodoRepoDeleteForUser(claims.UserID, uint(todoId))
	if todo.ID > 0 {
		WriteResponse(w, http.StatusOK, todo)
		return
	}

	WriteError(w, http.StatusNotFound, "Not found")
}

// User routes handlers

//Helpers
func WriteError(w http.ResponseWriter, status int, errorText string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	jsonErr := struct {
		Code int
		Text string
	}{
		status,
		errorText,
	}

	if err := json.NewEncoder(w).Encode(jsonErr); err != nil {
		panic(err)
	}
}

func WriteResponse(w http.ResponseWriter, status int, responseObj interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(responseObj); err != nil {
		panic(err)
	}
}
