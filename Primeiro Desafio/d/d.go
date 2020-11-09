package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

type CreditCard struct {
	Name       string
	Number     string
	ExpiryDate string
	Cvv        string
}

type Cards struct {
	Cards []CreditCard
}

func (card Cards) Check(name, number, expirydate, cvv string) string {
	for _, it := range card.Cards {
		if (name == it.Name) && (number == it.Number) && (expirydate == it.ExpiryDate) && (cvv == it.Cvv) {
			return "valid"
		}
	}

	return "invalid"
}

var cards Cards

func main() {

	mycard := CreditCard{
		Name:       "Saulo Pinedo",
		Number:     "1111222233334444",
		ExpiryDate: "10/28",
		Cvv:        "999",
	}

	cards.Cards = append(cards.Cards, mycard)

	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ccName := r.PostFormValue("ccName")
	ccNumber := r.PostFormValue("ccNumber")
	ccExpiration := r.PostFormValue("ccExpiration")
	ccCvv := r.PostFormValue("ccCvv")

	valid := cards.Check(ccName, ccNumber, ccExpiration, ccCvv)

	jsonResult, err := json.Marshal(Result{Status: valid})

	if err != nil {
		log.Fatal("Converting JSON task has failed.")
	}

	fmt.Fprintf(w, string(jsonResult))
}
