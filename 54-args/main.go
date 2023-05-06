package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// public static void main([]String args){

// }

// public static void Main([]string args){

// }

var (
	PORT   string
	DBCONN *string
	HTTP   bool
)

func main() {

	flag.StringVar(&PORT, "port", "8081", "--port=<port number> or -port=<port number>; Port must not be greater than 65000")

	DBCONN = flag.String("db", "some conn", "--db=<db connection string> or -db=<db connection string>")

	flag.BoolVar(&HTTP, "http-type", true, "--http-type=<true | false> or -db=<http-type=<true | false>")

	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Sony")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ok")
	})

	fmt.Println("Server started and listening on port", PORT)

	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		// println("There seems to be a problem in the server",err)
		// os.Exit(1)
		log.Fatal("There seems to be an error", err)
	}
	fmt.Println("Hello Sony")
}

// golang http servers are selfhosted
// that means there is no need of apache or iis or similar web server to run http based applications
// wifi, ether,lo, bluetooth,
