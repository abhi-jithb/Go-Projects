package main
import (
	"fmt"
	"net/http"
	"html/template"
	"log"
	)
	
	var temp = template.Must(template.New("form").Parse(
	`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Simple Server</title>
		</head>
		<body>
			<u><h1>Enter your details</h1></u>
			<form action="/submit" method="POST">
				<p>Name: <input type="text" name="name"></p><br>
				<p>Email: <input type="email" name="email"></p><br>
				<input type="submit" value="Submit">
			</form>
		</body>
		</html>
		
	`))
	func homeHandler(w http.ResponseWriter, r *http.Request){
		temp.Execute(w, nil)
	}
	func submitHandler(w http.ResponseWriter, r *http.Request){
		if(r.Method != http.MethodPost){
			http.Redirect(w,r, "/", http.StatusSeeOther)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		
		fmt.Fprintf(w, "<h2>Hello, %s!</h2><p>Your email: %s</p>", name, email)
	}
func main(){
	fmt.Println("This is a simple webserver")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/submit", submitHandler)
	fmt.Println("Port is running successfully at https://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
	}
