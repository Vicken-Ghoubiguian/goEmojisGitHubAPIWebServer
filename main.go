package main

//
import (
	"encoding/json"
	"fmt"
	"html/template"
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

//
type Page struct {
	Title                  string
	ListOfEmojisFromGitHub map[string]string
}

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
	if r.URL.Path != "/" {

		//
		fmt.Println(red + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] Error ? Yeah...." + reset)

		//
		t := template.New("Error tmpl")

		//
		t = template.Must(t.ParseFiles("tmpl/error.tmpl"))

		//
		err := t.ExecuteTemplate(w, "error", Page{"goEmojisGitHubAPIWebServer", nil})

		//
		otherErrorHandlerFunction(err)

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
		t := template.New("Main tmpl")

		//
		t = template.Must(t.ParseFiles("tmpl/main.tmpl"))

		//
		fmt.Println(green + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] New device connected: " + getIP(r) + "..." + reset)

		//
		err = t.ExecuteTemplate(w, "main", Page{"goEmojisGitHubAPIWebServer", currentlistOfEmojisFromGitHub})

		//
		otherErrorHandlerFunction(err)
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
		fmt.Println(cyan + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] Goodbye, we will miss you (" + strconv.Itoa(os.Getpid()) + ")..." + reset)

		//
		fmt.Println(purple + "---------------------------------" + reset)

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
		fmt.Println(cyan + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] Pressed Ctrl+z, suspended process " + strconv.Itoa(os.Getpid()) + "..." + reset + "\n")

		//
		fmt.Println(purple + "---------------------------------" + reset)

	}()
}

// Function which display other errors when they occurs
func otherErrorHandlerFunction(err error) {

	//
	if err != nil {

		//
		fmt.Println(red + err.Error() + reset)

		//
		os.Exit(1)
	}
}

// Function to execute the whole program
func main() {

	//
	fmt.Println(purple + "---------------------------------" + reset)

	//
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	//
	setup_ctrl_c_handler()

	//
	setup_ctrl_z_handler()

	//
	http.HandleFunc("/", helloWorldServerFunc)

	//
	http.ListenAndServe(":80", nil)
}
