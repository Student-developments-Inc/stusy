package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB Models
type User struct {
	ID        uint		`json:"id"`
	Email     string	`json:"email"`
	Password  []byte	`json:"password"`
	Token     string	`json:"token"`
	ExpiresAt int64		`json:"expires_at"`
}

type UserData struct {
	ID		uint
	UserID		uint	`json:"id"`
	FirstName	string	`json:"first_name"`
	MiddleName	string	`json:"middle_name"`
	LastName	string	`json:"last_name"`
}

type Course struct {
	ID		uint	`json:"id"`
	Author		uint	`json:"author"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
}

// Command structs
type Api struct {
	Config
	Router	*mux.Router
	DB	*gorm.DB
}

type Config struct {
	Port	string
	DSN	string
	Secret	string
}

// Main command functions' implementations
func (a *Api) Init() {
	var err error

	log.Println("Initializing config struct..")
	err = a.Config.Init()
	if err != nil {
		log.Fatal(err)
	}

	for {
		a.DB, err = gorm.Open(mysql.Open(a.Config.DSN), &gorm.Config{})
		if err != nil {
			log.Println("Failed to establish database connection, retrying...")
			time.Sleep(time.Second * 3)
		} else {
			log.Println("Database connection established.")
			break
		}
	}

	err = a.DB.AutoMigrate(&User{}, &UserData{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initializing routes..")
	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (c *Config) Init() error {
	c.DSN = fmt.Sprintf("%s:%s@tcp(stusy-db)/%s", os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	c.Port = os.Getenv("PORT")

	c.Secret = os.Getenv("SECRET")

	if len(c.Port) == 0 || len(c.Secret) == 0 {
		return fmt.Errorf("Config init failed. Some of the environment variables are blank.")
	}

	return nil
}

func (a *Api) initRoutes() {
	a.Router.HandleFunc("/", a.listRoutes).Methods("GET")

	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	a.Router.HandleFunc("/user/auth", a.loginUser).Methods("POST")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.isAuth(a.checkPerms(a.getUser))).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.isAuth(a.checkPerms(a.updateUser))).Methods("PUT")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")

	a.Router.HandleFunc("/courses", a.isAuth(a.getCourses)).Methods("GET")
	a.Router.HandleFunc("/course", a.isAuth(a.createCourse)).Methods("POST")
	a.Router.HandleFunc("/course/{id:[0-9]+}", a.isAuth(a.getCourse)).Methods("GET")
	a.Router.HandleFunc("/course/{id:[0-9]+}", a.isAuth(a.updateCourse)).Methods("PUT")
	a.Router.HandleFunc("/course/{id:[0-9]+}", a.isAuth(a.deleteCourse)).Methods("DELETE")
}

func (a *Api) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}

// Driver
func main() {
	a := Api{}
	a.Init()

	log.Println("Starting local server..")
	a.Run(":" + a.Config.Port)
}

// Utilities
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

func RespondWithJSON(w *http.ResponseWriter, status int, payload interface{}) {
	res, _ := json.MarshalIndent(payload, "", "	")

	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	(*w).Write(res)
}

func RespondWithError(w *http.ResponseWriter, status int, err string) {
	RespondWithJSON(w, status, map[string]string{"error": err})
}

// Middleware
func (a *Api) isAuth(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		header := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		if len(header) != 2 || !a.validateToken(header[1]) {
			RespondWithError(&w, http.StatusBadRequest, ErrorBadToken)
			return
		}

		n.ServeHTTP(w, req)
	}
}

func (a *Api) generateToken(iss string) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 12).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    iss,
		ExpiresAt: exp,
	})

	tok, err := claims.SignedString([]byte(a.Config.Secret))
	if err != nil {
		return "", 0, err
	}

	return tok, exp, nil
}

func (a *Api) validateToken(s string) bool {
	tok, err := a.parseToken(s)
	if err != nil {
		return false
	}

	if _, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
		return true
	}

	return false
}

func (a *Api) parseToken(s string) (*jwt.Token, error) {
	tok, err := jwt.Parse(s, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}

		return []byte(a.Config.Secret), nil
	})

	return tok, err
}

// Endpoint functions' implementations

// /
func (a *Api) listRoutes (w http.ResponseWriter, _ *http.Request) {
	err := a.Router.Walk(func(ro *mux.Route, r *mux.Router, anc []*mux.Route) error {
		var (
			path string
			methods []string
			err error
		)

		path, err = ro.GetPathTemplate()
		if err != nil {
			return err
		}

		methods, err = ro.GetMethods()
		if err != nil {
			return err
		}
			
		RespondWithJSON(&w, http.StatusOK, map[string]string{
			"path": path, "methods": strings.Join(methods, ",")})
		
		return nil
	})

	if err != nil {
		RespondWithError(&w, http.StatusInternalServerError, ErrorInternal)
		log.Println(err)
		return
	}
}

// /users
func (a *Api) getUsers(w http.ResponseWriter, req *http.Request) {
	
}

// /user
func (a *Api) createUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var pl struct {
		Email		string `json:"email"`
		Password	string `json:"password"`
	}

	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()

	if !validateRegData(pl.Password, pl.Email) {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(pl.Password), 12)
	u := User{
		Email:    strings.ToLower(pl.Email),
		Password: pass,
	}

	var t User
	a.DB.Where("email = ?", u.Email).First(&t)
	if t.ID != 0 {
		RespondWithError(&w, http.StatusForbidden, ErrorUserExist)
		return
	}

	a.DB.Create(&u)

	RespondWithJSON(&w, http.StatusCreated, map[string]uint{"uid": u.ID}) 
}

func validateRegData(pass, email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\d.\-]+@[a-zA-Z\d\-]+\.[a-zA-Z.]`).MatchString

	if len(pass) < 8 || !re(email) {
		return false
	}

	return true
}

// /user/auth
func (a *Api) loginUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var pl struct {
		Email		string `json:"email"`
		Password	string `json:"password"`
	}

	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()

	var u User
	a.DB.Where("email = ?", strings.ToLower(pl.Email)).First(&u)
	if u.ID == 0 || !validatePassHash(u.Password, []byte(pl.Password)) {
		RespondWithError(&w, http.StatusForbidden, ErrorBadCredentials)
		return
	}

	if len(u.Token) == 0 || !a.validateToken(u.Token) {
		var err error
		u.Token, u.ExpiresAt, err = a.generateToken(strconv.Itoa(int(u.ID)))
		if err != nil {
			log.Println(err)
			RespondWithError(&w, http.StatusInternalServerError, ErrorInternal)
			return
		}

		a.DB.Save(&u)
	}


	RespondWithJSON(&w, http.StatusOK, struct {
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

func validatePassHash(a []byte, b []byte) bool {
	err := bcrypt.CompareHashAndPassword(a, b)
	if err != nil {
		return false
	}
	return true
}

// /user/id
func (a *Api) getUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var udata UserData
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	a.DB.Where("user_id = ?", uint(id)).First(&udata)
	if udata.UserID == 0 {
		RespondWithError(&w, http.StatusNotFound, ErrorNotFound)
		return
	}

	RespondWithJSON(&w, http.StatusOK, map[string]string{
			"first_name": udata.FirstName,
			"middle_name": udata.MiddleName,
			"last_name": udata.LastName})
}

func (a *Api) updateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var pl struct {
		First	string `json:"first_name"`
		Middle	string `json:"middle_name"`
		Last	string `json:"last_name"`
	}

	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()

	re := regexp.MustCompile(`^\p{L}+$`).MatchString
	if !re(pl.First) || !re(pl.Middle) || !re(pl.Last) {
		RespondWithError(&w, http.StatusBadRequest, ErrorNameFormat)
		return
	}

	var udata UserData
	id, _ := strconv.Atoi(mux.Vars(req)["id"])

	a.DB.Where("user_id = ?", uint(id)).First(&udata)

	if udata.ID != 0 {
		RespondWithError(&w, http.StatusForbidden, "You can't change your name")
	}

	udata.FirstName = pl.First
	udata.MiddleName = pl.Middle
	udata.LastName = pl.Last
	udata.UserID = uint(id)

	a.DB.Create(&udata)

	RespondWithJSON(&w, http.StatusCreated, map[string]string{
			"first_name": udata.FirstName,
			"middle_name": udata.MiddleName,
			"last_name": udata.LastName})

}

func (a *Api) deleteUser(w http.ResponseWriter, req *http.Request) {

}

func (a *Api) checkPerms(n http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tok := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		id, _ := strconv.Atoi(mux.Vars(req)["id"])

		var u User
		a.DB.Where("token = ?", tok[1]).First(&u)
		if u.ID != uint(id) {
			RespondWithError(&w, http.StatusForbidden, ErrorAccess)
			return
		}

		n.ServeHTTP(w, req)
	}
}

// /courses
func (a *Api) getCourses(w http.ResponseWriter, req *http.Request) {

}

// /course
func (a *Api) getCourse(w http.ResponseWriter, req *http.Request) {

}

func (a *Api) createCourse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var pl struct {
		Name	string `json:"name"`
		Desc	string `json:"description"`
	}

	dec := json.NewDecoder(req.Body)
	if err := dec.Decode(&pl); err != nil {
		RespondWithError(&w, http.StatusBadRequest, ErrorBadJSON)
		return
	}
	defer req.Body.Close()

	var c = Course{
		Name: pl.Name,
		Description: pl.Desc,
	}

	a.DB.Create(&c)

	RespondWithJSON(&w, http.StatusCreated, map[string]uint{"id": c.ID}) 
}

func (a *Api) updateCourse(w http.ResponseWriter, req *http.Request) {

}

func (a *Api) deleteCourse(w http.ResponseWriter, req *http.Request) {

}
