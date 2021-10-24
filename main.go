package main

//
import (
	"fmt"
	"net/http"
	"os"
)

//

// Definition of all colors used in the weatherModule package
const (
	red   = "\033[31m"
	green = "\033[32m"
	cyan  = "\033[36m"
	reset = "\033[0m"
)

// ---------------------------------------------- Internal functions to run this module --------------------------------------

//
func helloWorldServerFunc(w http.ResponseWriter, r *http.Request) {

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
		//getEmojisFromGitHubAPIResp, err := http.Get(getEmojisFromGitHubAPIRequest)

		//

		//
		//getEmojisFromGitHubAPIJsonString, err := ioutil.ReadAll(getEmojisFromGitHubAPIResp.Body)

		//

		//
		fmt.Fprintf(w, "%s", getEmojisFromGitHubAPIRequest)
	}
}

// Function which display other errors when they occurs...
func OtherErrorHandlerFunction(err error, color string, reset string) {

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
