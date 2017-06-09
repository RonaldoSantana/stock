package sandbox

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"github.com/dgrijalva/jwt-go/request"
)

const signingKey = "Some long secret key"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SessionTokenResponse struct {
	Token string `json:"sessionToken"`
}

var users = map[string]User{}

var sessions = map[string]User{}

func writeJSON(w http.ResponseWriter, status int, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, status int, response interface{}) error {
	return writeJSON(w, status, ErrorResponse{Error: response.(string)})
}

func readUser(r *http.Request) (User, error) {
	var u User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}

func startSession(u User) string {
	sessionID := base64.URLEncoding.EncodeToString(uuid.NewV4().Bytes())
	sessions[sessionID] = u // oops - not thread safe
	return sessionID
}

func generateToken(sessionID string) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        sessionID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(signingKey))
}

func register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u, err := readUser(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	_, userExists := users[u.Email]
	if userExists {
		writeError(w, http.StatusBadRequest, "email already registered")
		return
	}
	
	users[u.Email] = u
	sessionID := startSession(u)
	tokenString, err := generateToken(sessionID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	err = writeJSON(w, http.StatusOK, SessionTokenResponse{Token: tokenString})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}

func login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u, err := readUser(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Read user")
	
	user, userExists := users[u.Email]
	if !userExists || user.Password != u.Password {
		//writeError(w, http.StatusBadRequest, "user email or password does not match")
		return
	}
	
	sessionID := startSession(user)
	
	tokenString, err := generateToken(sessionID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	err = writeJSON(w, http.StatusOK, SessionTokenResponse{Token: tokenString})
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}

func user(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	extractor := request.HeaderExtractor{
		"Authorization",
	}
	token, err := request.ParseFromRequestWithClaims(r, extractor, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	
	fmt.Println("Read token from request")
	
	if err != nil || !token.Valid {
		writeError(w, http.StatusBadRequest, err.Error())
		fmt.Printf("Error on token: %v, %s", err, token)
		return
	}
	
	claims := token.Claims.(*jwt.StandardClaims)
	fmt.Println(claims)
	
	user, validSession := sessions[claims.Id]
	if !validSession {
		writeError(w, http.StatusBadRequest, "invalid session")
		fmt.Printf("Invalid Session: %s, %v, %s", sessions, claims.Id, token.Claims)
		return
	}
	fmt.Println("Nothing is wrong so far")
	
	writeJSON(w, http.StatusOK, map[string]string{"email": user.Email})
}

func main() {
	router := httprouter.New()
	router.POST("/register", register)
	router.PUT("/login", login)
	router.GET("/user", user)
	http.ListenAndServe(":8080", router)
}