package main
 
import (
	"fmt"
	"sync"
	"log"
	"net/http"
 )

 var urls=[]string{
	"http://geeksforgeeks.com",
	"http://leetcode.com",
 }
 
 func callapiendpoints(w http.ResponseWriter,r *http.Request){
	var wg sync.WaitGroup
	for _,url:=range urls {
        wg.Add(1)
		go func(url string) {
			resp,err:=http.Get(url)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Fprintf(w,"%+v\n",resp)
			wg.Done()
		} (url)
		
	}
	wg.Wait()
 }
 func main() {
	fmt.Println("Go WaitGroup demo")
	http.HandleFunc("/",callapiendpoints)
	log.Fatal(http.ListenAndServe(":8080",nil))

 }