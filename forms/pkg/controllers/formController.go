package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/santosh/forms/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// ========== Start JWT =============

var secretKey = []byte("secret-key")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

//  ========== End JWT =============

func Login(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, filepath.Join("views/auth", "login.html"))
	} else if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")
		isAuthenticated := models.IsAuthenticated(email, password)

		if !isAuthenticated {
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(map[string]string{"message": "Invalid credentials"})
			return
		}
		tokenString, err := CreateToken(email)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(map[string]string{"message": "Token Invalid!"})
			return
		}
		// Set token as a cookie
		http.SetCookie(res, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		// http.Redirect(res, req, "/", http.StatusSeeOther)
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(map[string]string{"message": "Login Successfully!", "token": tokenString})
		return
	}
}

func Register(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		http.ServeFile(res, req, filepath.Join("views/auth", "register.html"))
	} else if req.Method == "POST" {

		// var user models.User
		// if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		// 	http.Error(res, err.Error(), http.StatusBadRequest)
		// 	return
		hashPass, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), 14)
		if err != nil {
			res.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(res).Encode(map[string]string{"message": "Failed to Register!"})
			return
		}
		imagePath := UploadImage(res, req)
		user := models.User{
			Name:     req.FormValue("name"),
			Email:    req.FormValue("email"),
			Password: string(hashPass),
			Profile:  imagePath,
		}
		createdUser := user.RegisterUser()
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(createdUser)
	}
}

func UploadImage(res http.ResponseWriter, req *http.Request) string {
	var imagePath string
	// Validate size of file
	req.ParseMultipartForm(10 << 20)

	// Get File From form
	file, handler, err2 := req.FormFile("image")
	if err2 != nil {
		fmt.Println(err2)
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"message": "Failed to Parse File!"})
		return ""
	}
	defer file.Close()

	// Create File
	des, err3 := os.Create(filepath.Join("assets/uploads", handler.Filename))
	imagePath = "assets/uploads" + handler.Filename
	if err3 != nil {
		fmt.Println(err3)
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"message": "Failed to Create File!"})
		return ""
	}
	defer des.Close()

	// Store File
	_, err5 := io.Copy(des, file)
	if err5 != nil {
		fmt.Println(err5)
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]string{"message": "Failed to Store File!"})
		return ""
	}
	return imagePath

}
func Home(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, filepath.Join("views", "index.html"))
	// http.ServeFile(res, req, filepath.Join("views/auth", "register.html"))
}

func Example(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{"message": "You have access to this protected route"})
}
