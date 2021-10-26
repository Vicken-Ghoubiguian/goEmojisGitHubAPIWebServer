package main

//
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

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

// Function to return client's IP adress
func getIP(r *http.Request) string {

	//
	forwarded := r.Header.Get("X-FORWARDED-FOR")

	//
	if forwarded != "" {

		//
		return forwarded
	}

	//
	return r.RemoteAddr
}

//
func helloWorldServerFunc(w http.ResponseWriter, r *http.Request) {

	//
	var currentlistOfEmojisFromGitHub map[string]string

	//
	fmt.Println(green + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] New device connected: " + getIP(r) + "..." + reset)

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
		//fmt.Println(green + currentlistOfEmojisFromGitHub["+1"] + reset)

		//
		fmt.Fprintf(w, "%s", currentlistOfEmojisFromGitHub)
	}

	//
	fmt.Println(blue + "---------------------------------" + reset)
}

// Function to handle Ctrl+c signal
func setup_ctrl_c_handler() {

	//
	c := make(chan os.Signal, 2)

	//
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	//
	go func() {

		//
		<-c

		//
		fmt.Println(cyan + "Goodbye, we will miss you (" + strconv.Itoa(os.Getpid()) + ")..." + reset)

		//
		fmt.Println(purple + "\n---------------------------------" + reset)

		//
		os.Exit(0)
	}()
}

// Function to handle Ctrl+Z signal
func setup_ctrl_z_handler() {

	//
	z := make(chan os.Signal, 20)

	//
	signal.Notify(z, os.Interrupt, syscall.SIGTSTP)

	//
	go func() {

		//
		<-z

		//
		fmt.Println(cyan + "Pressed Ctrl+z, suspended process " + strconv.Itoa(os.Getpid()) + "..." + reset + "\n")

	}()
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

//
func main() {

	//
	fmt.Println(purple + "\n---------------------------------" + reset)

	//
	setup_ctrl_c_handler()

	//
	setup_ctrl_z_handler()

	//
	http.HandleFunc("/", helloWorldServerFunc)

	//
	http.ListenAndServe(":80", nil)
}
