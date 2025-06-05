package main
import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found :(",http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
        }
        fmt.Fprintf(w, "hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "form.html")
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := r.FormValue("name")
		address := r.FormValue("address")
		phone := r.FormValue("phone")

		fmt.Fprintf(w, `
			<h2>Form Submitted Successfully</h2>
			<p><strong>Name:</strong> %s</p>
			<p><strong>Address:</strong> %s</p>
			<p><strong>Phone:</strong> %s</p>
			<p><a href="/form">Go Back</a></p>
		`, name, address, phone)
	}
}


func main(){
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Port is running successfully at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil{
	log.Fatal(err)
	}
}
