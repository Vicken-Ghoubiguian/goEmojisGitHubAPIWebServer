package main

//
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ---------------------------------------------- Internal types and structures to run this module --------------------------------------

// Definition of all colors used in the weatherModule package
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

//
type listOfEmojisFromGitHub struct {
	Plus_hand  string
	Minus_hand string
	Hundred    string
}

// ---------------------------------------------- Internal functions to run this module --------------------------------------

//
func helloWorldServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	//var currentlistOfEmojisFromGitHub listOfEmojisFromGitHub

	//
	if r.URL.Path != "/" {

		//
		fmt.Fprintf(w, "Error ? Yeah....")

		//
		fmt.Println(red + "Error ? Yeah...." + reset)

		//
	} else {

		//
		getEmojisFromGitHubAPIRequest := "https://api.github.com/emojis"

		//
		getEmojisFromGitHubAPIResp, err := http.Get(getEmojisFromGitHubAPIRequest)

		//
		otherErrorHandlerFunction(err)

		//
		getEmojisFromGitHubAPIJsonString, err := ioutil.ReadAll(getEmojisFromGitHubAPIResp.Body)

		//
		otherErrorHandlerFunction(err)

		//
		fmt.Fprintf(w, green+"%s"+reset, getEmojisFromGitHubAPIJsonString)
	}
}

// Function which display other errors when they occurs...
func otherErrorHandlerFunction(err error) {

	//
	if err != nil {

		//
		fmt.Println(red + err.Error() + reset)

		//
		os.Exit(1)
	}
}

// ---------------------------------------------- External function to use this module --------------------------------------

//
func main() {

	//
	http.HandleFunc("/", helloWorldServerFunc)

	//
	http.ListenAndServe(":80", nil)
}
