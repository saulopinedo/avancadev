package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-retryablehttp"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/process", process)
	http.ListenAndServe(":9090", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w, Result{})
}

func process(w http.ResponseWriter, r *http.Request) {
	//log.Println(r.FormValue("coupon"))
	//log.Println(r.FormValue("cc-number"))

	result := makeHttpCall(
		"http://localhost:9091",
		r.FormValue("coupon"),
		r.FormValue("cc-name"),
		r.FormValue("cc-number"),
		r.FormValue("cc-expiration"),
		r.FormValue("cc-cvv"))

	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w, result)
}

func makeHttpCall(urlMicroservice, coupon, ccName, ccNumber, ccExpiration, ccCvv string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)
	values.Add("ccName", ccName)
	values.Add("ccNumber", ccNumber)
	values.Add("ccExpiration", ccExpiration)
	values.Add("ccCvv", ccCvv)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "The mid server is offline!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("The mid server is offline!")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}
