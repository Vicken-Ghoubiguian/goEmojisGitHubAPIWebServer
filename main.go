package main

//
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ---------------------------------------------- Internal types and structures to run this module --------------------------------------

// Definition of all colors used in this project
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
/*type listOfEmojisFromGitHub struct {
	Plus_hand          string `json:"+1"`
	Minus_hand         string `json:"-1"`
	Hundred            string `json:"100"`
	One_two_three_four string `json:"1234"`
	First_place_medal  string `json:"1st_place_medal"`
	Second_place_medal string `json:"2st_place_medal"`
	Third_place_medal  string `json:"3st_place_medal"`
	Eight_ball         string `json:"8ball"`
	A_emoji            string `json:"a"`
	AB_emoji           string `json:"ab"`
	Abacus             string `json:"abacus"`
	ABC_emoji          string `json:"abc"`
	ABCD_emoji         string `json:"abcd"`
}*/

// ---------------------------------------------- Internal functions to run this module --------------------------------------

//
func helloWorldServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	//var currentlistOfEmojisFromGitHub listOfEmojisFromGitHub
	var currentlistOfEmojisFromGitHub map[string]string

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

		truc := []byte(getEmojisFromGitHubAPIJsonString)

		//Single instruction to convert weather_json_string []byte variable to string
		err = json.Unmarshal(truc, &currentlistOfEmojisFromGitHub)

		//
		otherErrorHandlerFunction(err)

		//
		fmt.Println(currentlistOfEmojisFromGitHub)

		//
		fmt.Fprintf(w, green+"%s"+reset, currentlistOfEmojisFromGitHub["+1"])
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
