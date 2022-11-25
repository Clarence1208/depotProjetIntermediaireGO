package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func clockHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		now := time.Now()
		fmt.Fprintf(w, "Il est %02dh%02dmin%02ds", now.Hour(), now.Minute(), now.Second())
	case http.MethodPost:
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
	}
}

func diceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		dice := r1.Intn(1000) + 1
		fmt.Fprintf(w, "Le résultat du dé a 1000 faces est: %04d", dice)
	case http.MethodPost:
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
	}
}
func dicesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		var differentDice = [8]int{2, 4, 6, 8, 10, 12, 20, 100}
		var diceValue int
		for i := 0; i < 15; i++ {
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			r2 := rand.New(s1)
			slctDice := r2.Intn(8)
			if req.URL.Query().Get("type") == "" {
				diceValue = differentDice[slctDice]
			} else {
				int1, err := strconv.Atoi(req.URL.Query().Get("type"))
				if err != nil {
					errors.New("NullPointerException")
					return
				}
				diceValue = int1
			}
			dice := r1.Intn(diceValue) + 1
			if diceValue < 10 {
				fmt.Fprintf(w, "%d ", dice)
			} else if diceValue < 100 {
				fmt.Fprintf(w, "%02d ", dice)
			} else {
				fmt.Fprintf(w, "%03d ", dice)
			}
		}
	case http.MethodPost:
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
	}
}

func randWordsHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		/*for key, value := range req.PostForm {
			if key == "words"{
				for len(value)
			}

		}

		*/
	}
}

func semiCapitalHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Println("Something went bad")
		fmt.Fprintln(w, "Something went bad")
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}
	}
}

func main() {
	http.HandleFunc("/", clockHandler)
	http.HandleFunc("/dice", diceHandler)
	http.HandleFunc("/dices", dicesHandler)
	http.HandleFunc("/randomize-words", randWordsHandler)
	http.HandleFunc("/semi-capitalize-sentence", semiCapitalHandler)
	http.ListenAndServe(":4567", nil)
}
