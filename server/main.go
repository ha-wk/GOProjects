package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

//var mysignkey = os.Getenv("MY_TOKEN")

//OR

var mysignkey = []byte("bhtsecure")

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "bht secure key")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return mysignkey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "not authorized")
		}
	})
}
func handleRequests() {
	http.Handle("/", isAuthorized(homepage))

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("my server")

	handleRequests()

}
