package main

import (
	"encoding/json"
	"regexp"
	"strconv"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"os"
	"time"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	ErrorInternal       = "Internal server error. Try again later."
	ErrorBadJSON        = "You have supplied an invalid JSON body."
	ErrorBadCredentials = "Invalid credentials. Invalid email or password."
	ErrorUserExist      = "Email is already in use."
	ErrorEmailFormat    = "Invalid email format."
	ErrorShortPass      = "Password is too short (should be >= 8)."
	ErrorBadToken       = "Token is invalid or expired."
	ErrorAccess         = "You don't have permissions to make this request."
	ErrorNameFormat     = "Name should only consist of alphabetical characters."
	ErrorNotFound       = "Resource not found."
)

type User struct {
	ID        uint
	Email     string
	Password  []byte
	Token     string
	ExpiresAt int64
}

type UserData struct {
	ID        uint
	UserID    uint
	FirstName string
	LastName  string
}

type Route struct {
	Route   string
	Methods string
}

var (
	Secret	string
	Port	string
	DSN	string
	DB	*gorm.DB
	r	*mux.Router
)

// Init funcs
func
InitDB() error {
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&User{}, &UserData{})
	if err != nil {
		return err
	}

	return nil
}

func
InitConfig() error {
	DSN = fmt.Sprintf("%s:%s@tcp(stusy-db)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	Port = ":" + os.Getenv("PORT")
	Secret = os.Getenv("SECRET")

	if Port == "" || Secret == "" {
		return fmt.Errorf("some of environment variables are blank or missing (make sure to set them right)")
	}

	return nil
}

func
InitRouter() *mux.Router {
	r = mux.NewRouter()

	r.HandleFunc("/", List()).Methods("GET", "OPTIONS")
//	r.HandleFunc("/users", Count()).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/login", SignIn()).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/register", SignUp()).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/{id:[0-9]+}", IsAuth(Info())).Methods("GET", "PUT", "OPTIONS")

	r.Use(mux.CORSMethodMiddleware(r))

	return r
}

// Handler funcs
func
List() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			output := make(map[string]string, 15)
			pathTemplate, err := route.GetPathTemplate()
			if err == nil {
				output["route"] = pathTemplate
			}
			methods, err := route.GetMethods()
			if err == nil {
				output["methods"] = strings.Join(methods, ",")
			}

			listMessage(&w, output["route"], output["methods"])
			return nil
		})

		if err != nil {
			WriteError(&w, http.StatusInternalServerError, ErrorInternal)
			log.Println(err)
			return
		}
	}
}

func
SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}

		data, err := ReadJSON(&w, r)
		if err != nil {
			return
		}

		if len(data["password"]) < 8 {
			WriteError(&w, http.StatusBadRequest, ErrorShortPass)
			return
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

		re := regexp.MustCompile(`^[a-zA-Z\d.\-]+@[a-zA-Z\d\-]+\.[a-zA-Z.]`).MatchString
		if !re(data["email"]) {
			WriteError(&w, http.StatusBadRequest, ErrorEmailFormat)
			return
		}

		user := User{
			Email:    strings.ToLower(data["email"]),
			Password: password,
		}

		var temp User
		DB.Where("email = ?", user.Email).First(&temp)
		if temp.ID != 0 {
			WriteError(&w, http.StatusForbidden, ErrorUserExist)
			return
		}

		DB.Create(&user)

		registerMessage(&w, &user.ID)
	}
}

func
SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}

		data, err := ReadJSON(&w, r)
		if err != nil {
			return
		}

		var user User

		DB.Where("email = ?", strings.ToLower(data["email"])).First(&user)

		if user.ID == 0 {
			WriteError(&w, http.StatusForbidden, ErrorBadCredentials)
			return
		}

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
			WriteError(&w, http.StatusForbidden, ErrorBadCredentials)
			return
		}

		if len(user.Token) > 0 {
			if ok, _ := ValidateToken(user.Token); ok {
				loginMessage(&w, &user.Token, &user.ID, &user.ExpiresAt)
				return
			}
		}

		user.Token, user.ExpiresAt, err = CreateToken(strconv.Itoa(int(user.ID)))
		if err != nil {
			log.Println(err)
			WriteError(&w, http.StatusInternalServerError, ErrorInternal)
			return
		}

		DB.Save(&user)

		loginMessage(&w, &user.Token, &user.ID, &user.ExpiresAt)
	}
}

func
Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var user User
		var userData UserData

		DB.Where("token = ?", token[1]).First(&user)
		if user.ID != uint(id) {
			WriteError(&w, http.StatusForbidden, ErrorAccess)
			return
		}

		if r.Method == "GET" {
			DB.Where("user_id = ?", uint(id)).First(&userData)
			if userData.UserID == 0 {
				WriteError(&w, http.StatusNotFound, ErrorNotFound)
				return
			}

			infoMessage(&w, &userData.FirstName, &userData.LastName)
			return
		}

		data, err := ReadJSON(&w, r)
		if err != nil {
			return
		}

		re := regexp.MustCompile(`^\p{L}+$`).MatchString
		if !re(data["first_name"]) || !re(data["last_name"]) {
			WriteError(&w, http.StatusBadRequest, ErrorNameFormat)
			return
		}

		DB.Where("user_id = ?", uint(id)).First(&userData)

		userData.LastName = data["last_name"]
		userData.FirstName = data["first_name"]

		if userData.ID == 0 {

			res, _ := json.MarshalIndent(struct {
				FirstName string `json:"first_name"`
				LastName  string `json:"last_name"`
			}{
				FirstName: userData.LastName,
				LastName:  userData.FirstName,
			}, "", "	")

			userData.UserID = user.ID
			DB.Create(&userData)
			w.WriteHeader(http.StatusCreated)
			w.Write(res)
			return
		}

		w.WriteHeader(http.StatusForbidden)
	}
}

// Response funcs
// TODO: Implement generic func
func listMessage(w *http.ResponseWriter, path, methods string) {
	res, _ := json.MarshalIndent(struct {
		Path    string `json:"route"`
		Methods string `json:"methods"`
	}{
		Path:    path,
		Methods: methods,
	}, "", "	")

	(*w).Write(res)
}

func
registerMessage(w *http.ResponseWriter, id *uint) {
	res, _ := json.MarshalIndent(struct {
		UserID uint   `json:"user_id"`
		Href   string `json:"href"`
	}{
		UserID: *id,
		Href:   "https://api.studentsystem.xyz/profile/" + strconv.Itoa(int(*id)),
	}, "", "	")

	(*w).WriteHeader(http.StatusCreated)
	(*w).Write(res)
}

func
loginMessage(w *http.ResponseWriter, token *string, id *uint, expiresAt *int64) {
	res, _ := json.MarshalIndent(struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
		UserID      uint   `json:"user_id"`
	}{
		TokenType:   "bearer",
		AccessToken: *token,
		ExpiresIn:   *expiresAt - time.Now().Unix(),
		UserID:      *id,
	}, "", "	")

	(*w).WriteHeader(http.StatusOK)
	(*w).Write(res)
}

func
infoMessage(w *http.ResponseWriter, first, last *string) {
	res, _ := json.MarshalIndent(struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{
		FirstName: *first,
		LastName:  *last,
	}, "", "	")

	(*w).WriteHeader(http.StatusOK)
	(*w).Write(res)
}


// Util funcs
func
ReadJSON(w *http.ResponseWriter, r *http.Request) (map[string]string, error) {
	var data map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, ErrorInternal)
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		WriteError(w, http.StatusBadRequest, ErrorBadJSON)
		return nil, err
	}

	return data, nil
}

func // TODO: Must be generic to work with any struct
WriteError(w *http.ResponseWriter, status int, err string) {
	res, _ := json.MarshalIndent(struct {
		Message string `json:"error_message"`
	}{
		Message: err,
	}, "", "	")

	(*w).WriteHeader(status)
	(*w).Write(res)
}

// Middleware
func
IsAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			return
		}

		header := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(header) != 2 || header[0] == "null" {
			WriteError(&w, http.StatusBadRequest, ErrorBadToken)
			return
		}

		if ok, _ := ValidateToken(header[1]); !ok {
			WriteError(&w, http.StatusUnauthorized, ErrorBadToken)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func
CreateToken(issuer string) (string, int64, error) {
	expiresAt := time.Now().Add(time.Hour * 12).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: expiresAt,
	})

	token, err := claims.SignedString([]byte(Secret))
	if err != nil {
		return "", 0, err
	}
	return token, expiresAt, nil
}

func
ValidateToken(tokenString string) (bool, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}

func
parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Secret), nil
	})

	return token, err
}



func
main() {
	var err error
	log.Println("Initializing config structure")
	err = InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Initializing database connections pool")
		err = InitDB()
		if err == nil {
			break
		} else {
			time.Sleep(3 * time.Second)
		}
	}

	log.Println("Starting local server")
	err = http.ListenAndServe(Port, InitRouter())
	if err != nil {
		log.Fatal(err)
	}
}
