package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hello root route")
		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			fmt.Printf("hello to root route via GET method")
			w.Write([]byte("Hello from root route via GET method"))
		case http.MethodPost:
			fmt.Printf("hello to root route via POST method")
			w.Write([]byte("Hello from root route via POST method"))
		case http.MethodPut:
			fmt.Printf("hello to root route via PUT method")
			w.Write([]byte("Hello from root route via PUT method"))
		case http.MethodDelete:
			fmt.Printf("hello to root route via DELETE method")
			w.Write([]byte("Hello from root route via DELETE method"))
		case http.MethodPatch:
			fmt.Printf("hello to root route via PATCH method")
			w.Write([]byte("Hello from root route via PATCH method"))
		default:
			fmt.Printf("hello to root route via unknown method")
			w.Write([]byte("Hello from root route via unknown method"))
		}

	})

	http.HandleFunc("/teachers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hello techer route")
		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			fmt.Printf("hello to teacher route via GET method")
			w.Write([]byte("Hello from teacher route via GET method"))
		case http.MethodPost:
			fmt.Printf("hello to teacher route via POST method")
			w.Write([]byte("Hello from teacher route via POST method"))
		case http.MethodPut:
			fmt.Printf("hello to teacher route via PUT method")
			w.Write([]byte("Hello from teacher route via PUT method"))
		case http.MethodDelete:
			fmt.Printf("hello to teacher route via DELETE method")
			w.Write([]byte("Hello from teacher route via DELETE method"))
		case http.MethodPatch:
			fmt.Printf("hello to teacher route via PATCH method")
			w.Write([]byte("Hello from teacher route via PATCH method"))
		default:
			fmt.Printf("hello to teacher route via unknown method")
			w.Write([]byte("Hello from teacher route via unknown method"))
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hello student route")
		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			fmt.Printf("hello to students route via GET method")
			w.Write([]byte("Hello from students route via GET method"))
		case http.MethodPost:
			fmt.Printf("hello to students route via POST method")
			w.Write([]byte("Hello from students route via POST method"))
		case http.MethodPut:
			fmt.Printf("hello to students route via PUT method")
			w.Write([]byte("Hello from students route via PUT method"))
		case http.MethodDelete:
			fmt.Printf("hello to students route via DELETE method")
			w.Write([]byte("Hello from students route via DELETE method"))
		case http.MethodPatch:
			fmt.Printf("hello to students route via PATCH method")
			w.Write([]byte("Hello from students route via PATCH method"))
		default:
			fmt.Printf("hello to students route via unknown method")
			w.Write([]byte("Hello from students route via unknown method"))
		}

	})

	http.HandleFunc("/execs", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("hello execs route")
		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodGet:
			fmt.Printf("hello to execs route via GET method")
			w.Write([]byte("Hello from execs route via GET method"))
		case http.MethodPost:
			fmt.Printf("hello to execs route via POST method")
			w.Write([]byte("Hello from execs route via POST method"))
		case http.MethodPut:
			fmt.Printf("hello to execs route via PUT method")
			w.Write([]byte("Hello from execs route via PUT method"))
		case http.MethodDelete:
			fmt.Printf("hello to execs route via DELETE method")
			w.Write([]byte("Hello from execs route via DELETE method"))
		case http.MethodPatch:
			fmt.Printf("hello to execs route via PATCH method")
			w.Write([]byte("Hello from execs route via PATCH method"))
		default:
			fmt.Printf("hello to execs route via unknown method")
			w.Write([]byte("Hello from execs route via unknown method"))
		}

	})

	port := ":3000"
	fmt.Println("Starting thr server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("server failed to start due to error", err)
	}
}
