package main

import "net/http"

func sayHello(w http.ResponseWriter, r *http.Request) {
	planet := r.URL.Query().Get("planet")
	w.Write([]byte("Hello, world" + planet))
}

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}

// firefox 127.0.0.1:5050

// firefox http://localhost:5050/?planet=World
