package main

//
import (
	"fmt"
	"net/http"
)

// ---------------------------------------------------------------------------------------------------------------------------------

//
func helloWorldServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	if r.URL.Path != "/" {

		fmt.Fprintf(w, "Error ? Yeah....")

		//
	} else {

		//
		getEmojisFromGitHubAPIRequest := "https://api.github.com/emojis"

		//
		//getEmojisFromGitHubAPIResp, err := http.Get(getEmojisFromGitHubAPIRequest)

		//
		//getEmojisFromGitHubAPIJsonString, err := ioutil.ReadAll(getEmojisFromGitHubAPIResp.Body)

		//
		fmt.Fprintf(w, "%s", getEmojisFromGitHubAPIRequest)
	}
}

// ---------------------------------------------------------------------------------------------------------------------------------

//
func main() {

	//
	http.HandleFunc("/", helloWorldServerFunc)

	//
	http.ListenAndServe(":80", nil)
}
