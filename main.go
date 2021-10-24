package main

//
import (
	"fmt"
	"net/http"
)

//
func main() {

	//
	http.HandleFunc("/", helloWorldServerFunc)

	//
	http.ListenAndServe(":80", nil)
}

//
func helloWorldServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	if r.URL.Path != "/" {

		fmt.Fprintf(w, "Error ? Yeah....")

		//
	} else {

		fmt.Fprintf(w, "Hello, world !")
	}
}
