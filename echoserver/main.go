package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("开启echoserver!")
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {{
		fmt.Fprintln(w,"Hello world!")
	}})
	http.HandleFunc("/url",func(w http.ResponseWriter, r *http.Request) {{
		data := fmt.Sprintf("%+v",r)
		data = strings.Join(strings.Split(data," "),"\n") 
		fmt.Fprintln(w,data)
	}})
	http.ListenAndServe(":8080",nil)
}