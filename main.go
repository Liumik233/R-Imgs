package main

import (
	"log"
	"net/http"
)

func papi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Location", "https://od.liumik.tech/sjhl/Picture")
	w.WriteHeader(307)
}
func main() {
	http.HandleFunc("/papi", papi)
	err := http.ListenAndServe(":1701", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
