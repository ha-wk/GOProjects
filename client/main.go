package main

import (
	"fmt"
	"log"
	"net/http"

	//"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//var mysignkey = os.Getenv("MY_TOKEN")

//OR

var mysignkey = []byte("bhtsecure")

func homepage(w http.ResponseWriter, r *http.Request) {

	validToken, err := generatejwt()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, validToken)
}

func generatejwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "AYUSH KUMAR"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mysignkey)

	if err != nil {
		fmt.Errorf("something went wrong:%s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func handleRequests() {
	http.HandleFunc("/", homepage)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("my client")

	handleRequests()

}
