package main

// Import all of the used packages in this web application (all coming from the standard library - see https://pkg.go.dev/std)
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

// Definition of all colors used in this project (actually more)
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

// Definition of the Page structure used in this project to define a page
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

// Function which manage the filled in URL and all of this web application
func onlyAndMainHandlerFunc(w http.ResponseWriter, r *http.Request) {

	// Definition of the 'currentlistOfEmojisFromGitHub' which is a map containing all emojis collected from GitHub
	var currentlistOfEmojisFromGitHub map[string]string

	//
	if r.URL.Path != "/" {

		// To display information message in cyan
		fmt.Println(red + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] Wrong URL error: the application is at http://localhost..." + reset)

		// Definition of the error template
		t := template.New("Error tmpl")

		// To parse all tmpl files which contain the error template
		t = template.Must(t.ParseFiles("tmpl/header.tmpl", "tmpl/error.tmpl", "tmpl/footer.tmpl", "tmpl/end.tmpl"))

		//
		err := t.ExecuteTemplate(w, "error", Page{"goEmojisGitHubAPIWebServer", nil})

		// Manage the possible occured error
		otherErrorHandlerFunction(err)

		// In the other case...
	} else {

		// Definition of the HTTPS request's URL to get all emojis from GitHub
		getEmojisFromGitHubAPIRequest := "https://api.github.com/emojis"

		// Execution of the Get HTTPS request to get all emojis from GitHub
		getEmojisFromGitHubAPIResp, err := http.Get(getEmojisFromGitHubAPIRequest)

		// Manage the possible occured error
		otherErrorHandlerFunction(err)

		//
		getEmojisFromGitHubAPIJsonString, err := ioutil.ReadAll(getEmojisFromGitHubAPIResp.Body)

		// Manage the possible occured error
		otherErrorHandlerFunction(err)

		//
		getEmojisFromGitHubAPIJsonStringAsByte := []byte(getEmojisFromGitHubAPIJsonString)

		// Single instruction to convert weather_json_string []byte variable to string
		err = json.Unmarshal(getEmojisFromGitHubAPIJsonStringAsByte, &currentlistOfEmojisFromGitHub)

		// Manage the possible occured error
		otherErrorHandlerFunction(err)

		// Definition of the main template
		t := template.New("Main tmpl")

		// To parse all tmpl files which contain the main template
		t = template.Must(t.ParseFiles("tmpl/header.tmpl", "tmpl/main.tmpl", "tmpl/footer.tmpl", "tmpl/end.tmpl"))

		// To display information message in green
		fmt.Println(green + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] New device connected: " + getIP(r) + "..." + reset)

		//
		err = t.ExecuteTemplate(w, "main", Page{"goEmojisGitHubAPIWebServer", currentlistOfEmojisFromGitHub})

		// Manage the possible occured error
		otherErrorHandlerFunction(err)
	}

	// To display the blue separator
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

		// To display information message in cyan
		fmt.Println(cyan + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] Goodbye, we will miss you (" + strconv.Itoa(os.Getpid()) + ")..." + reset)

		// To display the purple separator
		fmt.Println(purple + "---------------------------------" + reset)

		// Exit the program (with exit code 0 (without any error))
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

		// To display information message
		fmt.Println(cyan + "[UTC time: " + time.Now().UTC().Format("January 02 2006 03:04:05") + "] Pressed Ctrl+z, suspended process " + strconv.Itoa(os.Getpid()) + "..." + reset + "\n")

		// To display the purple separator
		fmt.Println(purple + "---------------------------------" + reset)

	}()
}

// Function which display other errors when they occurs
func otherErrorHandlerFunction(err error) {

	// If the error is null (in this case, there is no error)
	if err != nil {

		// To display the error in red
		fmt.Println(red + err.Error() + reset)

		// Exit the program (with exit code 1 (with error))
		os.Exit(1)
	}
}

// Function to execute the whole program
func main() {

	// To display the purple separator
	fmt.Println(purple + "---------------------------------" + reset)

	// To handle the 'assets' directory to use images, css and js files
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Implementation of the CTRL - C handler
	setup_ctrl_c_handler()

	// Implementation of the CTRL - Z handler
	setup_ctrl_z_handler()

	// To register the 'onlyAndMainHandlerFunc' function and to correspond it to the root URL
	http.HandleFunc("/", onlyAndMainHandlerFunc)

	// To listen on the current local TCP network address at the port 80
	http.ListenAndServe(":80", nil)
}
