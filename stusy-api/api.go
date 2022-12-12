package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type Error struct {
	Message string `json:"error_message"`
}

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

var (
	secret	string
	port	string
	dsn	string
	db	*gorm.DB
	r	*mux.Router
)

// Init funcs
func
initdb() error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&User{}, &UserData{})
	if err != nil {
		return err
	}

	return nil
}

func
initcfg() error {
	dsn = fmt.Sprintf("%s:%s@tcp(stusy-db)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	port = ":" + os.Getenv("PORT")
	secret = os.Getenv("SECRET")

	if port == "" || secret == "" {
		return fmt.Errorf("some of the environment variables are blank")
	}

	return nil
}

func
initrouter() *mux.Router {
	r = mux.NewRouter()

	r.HandleFunc("/", listroutes()).Methods("GET", "OPTIONS")
//	r.HandleFunc("/users", Count()).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/login", userlogin()).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/register", usercreation()).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/{id:[0-9]+}", auth(userinfo())).Methods("GET", "PUT", "OPTIONS")

	r.Use(mux.CORSMethodMiddleware(r))

	return r
}

// Handler funcs
func
listroutes() http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		err := r.Walk(func(route *mux.Route, r *mux.Router, ancestors []*mux.Route) error {
			output := make(map[string]string, 15)
			
			if path, err := route.GetPathTemplate(); err == nil {
				output["route"] = path
			}

			if methods, err := route.GetMethods(); err == nil {
				output["methods"] = strings.Join(methods, ",")
			}

			respond(&w, http.StatusOK, struct {
					Path    string `json:"route"`
					Methods string `json:"methods"`
				}{
					Path: output["route"],
					Methods: output["methods"],
				})
			
			return nil
		})

		if err != nil {
			respond(&w, http.StatusInternalServerError, Error{
					Message: ErrorInternal,
				})

			log.Println(err)
			return
		}
	}
}

func
usercreation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}

		data, err := parsereq(&w, r)
		if err != nil {
			return
		}

		if len(data["password"]) < 8 {
			respond(&w, http.StatusBadRequest, Error{
					Message: ErrorShortPass,
				})

			return
		}

		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)

		re := regexp.MustCompile(`^[a-zA-Z\d.\-]+@[a-zA-Z\d\-]+\.[a-zA-Z.]`).MatchString
		if !re(data["email"]) {
			respond(&w, http.StatusBadRequest, Error{
					Message: ErrorEmailFormat,
				})

			return
		}

		u := User{
			Email:    strings.ToLower(data["email"]),
			Password: password,
		}

		var tmp User
		db.Where("email = ?", u.Email).First(&tmp)
		if tmp.ID != 0 {
			respond(&w, http.StatusForbidden, Error{
					Message: ErrorUserExist,
				})
			return
		}

		db.Create(&u)

		respond(&w, http.StatusCreated, struct {
				UserID uint `json:"user_id"`
			}{
				UserID: u.ID,
			})
	}
}

func
userlogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "OPTIONS" {
			return
		}

		data, err := parsereq(&w, r)
		if err != nil {
			return
		}

		var u User

		db.Where("email = ?", strings.ToLower(data["email"])).First(&u)
		if u.ID == 0 {
			respond(&w, http.StatusForbidden, Error{
					Message: ErrorBadCredentials,
				})
			return
		}

		if err := bcrypt.CompareHashAndPassword(u.Password, []byte(data["password"])); err != nil {
			respond(&w, http.StatusForbidden, Error{
					Message: ErrorBadCredentials,
				})
			return
		}

		if len(u.Token) == 0 || !ValidateToken(u.Token) {
			u.Token, u.ExpiresAt, err = CreateToken(strconv.Itoa(int(u.ID)))
			if err != nil {
				log.Println(err)
				respond(&w, http.StatusInternalServerError, Error{
						Message: ErrorInternal,
					})
				return
			}

			db.Save(&u)
		}


		respond(&w, http.StatusOK, struct {
				TokenType   string `json:"token_type"`
				AccessToken string `json:"access_token"`
				ExpiresIn   int64  `json:"expires_in"`
				UserID      uint   `json:"user_id"`
			}{
				TokenType:   "bearer",
				AccessToken: u.Token,
				ExpiresIn:   u.ExpiresAt - time.Now().Unix(),
				UserID:      u.ID,
			})
	}
}

func
userinfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var ( 
			u	User
			udata	UserData
		)

		db.Where("token = ?", token[1]).First(&u)
		if u.ID != uint(id) {
			respond(&w, http.StatusForbidden, Error{
					Message: ErrorAccess,
				})
			return
		}

		if r.Method == "GET" {
			db.Where("user_id = ?", uint(id)).First(&udata)
			if udata.UserID == 0 {
				respond(&w, http.StatusNotFound, Error{
						Message: ErrorNotFound,
					})
				return
			}

			respond(&w, http.StatusOK, struct {
					FirstName string `json:"first_name"`
					LastName  string `json:"last_name"`
				}{
					FirstName: udata.FirstName,
					LastName:  udata.LastName,
				})

			return
		}

		data, err := parsereq(&w, r)
		if err != nil {
			return
		}

		re := regexp.MustCompile(`^\p{L}+$`).MatchString
		if !re(data["first_name"]) || !re(data["last_name"]) {
			respond(&w, http.StatusBadRequest, Error{
					Message: ErrorNameFormat,
				})

			return
		}

		db.Where("user_id = ?", uint(id)).First(&udata)

		udata.LastName = data["last_name"]
		udata.FirstName = data["first_name"]

		if udata.ID == 0 {
			respond(&w, http.StatusCreated, struct {
					FirstName string `json:"first_name"`
					LastName  string `json:"last_name"`
				}{
					FirstName: udata.FirstName,
					LastName:  udata.LastName,
				})

			udata.UserID = u.ID
			db.Create(&udata)

			return
		}

		w.WriteHeader(http.StatusForbidden)
	}
}

// Util funcs
func
respond(w *http.ResponseWriter, status int, data interface{}) {
	res, _ := json.MarshalIndent(data, "", "	")

	(*w).WriteHeader(status)
	(*w).Write(res)
}

func
parsereq(w *http.ResponseWriter, r *http.Request) (map[string]string, error) {
	var data map[string]string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respond(w, http.StatusInternalServerError, Error{
				Message: ErrorInternal,
			})
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		respond(w, http.StatusBadRequest, Error{
				Message: ErrorBadJSON,
			})
		return nil, err
	}

	return data, nil
}

// Middleware
func
auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			return
		}

		header := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(header) != 2 || header[0] == "null" {
			respond(&w, http.StatusBadRequest, Error{
					Message: ErrorBadToken,
				})
			return
		}

		if !ValidateToken(header[1]) {
			respond(&w, http.StatusUnauthorized, Error{
					Message: ErrorBadToken,
				})
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

	token, err := claims.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}
	return token, expiresAt, nil
}

func
ValidateToken(tokenString string) bool {
	token, err := parseToken(tokenString)
	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}

func
parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	return token, err
}

// Driver
func
main() {

	log.Println("Initializing config structure")
	if err := initcfg(); err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing database connections pool")
	if err := initdb(); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting local server")
	if err := http.ListenAndServe(port, initrouter()); err != nil {
		log.Fatal(err)
	}
}
