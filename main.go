package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var url []string

func readfile() {
	f, e := os.Open("url.txt")
	defer f.Close()
	if e != nil {
		fmt.Println("error:", e)
	} else {
		str := bufio.NewReader(f)
		for {
			r, err := str.ReadString('\n')
			url = append(url, strings.TrimSpace(r))
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("readfile error:", err)
			}
		}
	}
}
func papi(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(url) - 1)
	log.Println(len(url))
	w.Header().Add("Location", url[i])
	w.WriteHeader(307)
}
func main() {
	readfile()
	http.HandleFunc("/papi", papi)
	err := http.ListenAndServe(":1701", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
