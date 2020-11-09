package main

import (
	"encoding/json"
	"fmt"
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
	http.ListenAndServe(":9091", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	coupon := r.PostFormValue("coupon")
	ccName := r.PostFormValue("ccName")
	ccNumber := r.PostFormValue("ccNumber")
	ccExpiration := r.PostFormValue("ccExpiration")
	ccCvv := r.PostFormValue("ccCvv")

	result := Result{Status: "Compra concluída!"}

	resultCard := CardServiceCaller("http://localhost:9093", ccName, ccNumber, ccExpiration, ccCvv)

	if resultCard.Status == "invalid" {
		result.Status = "This card is invalid. Please check its informations."
	}

	if len(coupon) != 0 {
		resultCoupon := CouponServiceCaller("http://localhost:9092", coupon)
		if resultCoupon.Status == "invalid" {
			result.Status = "Invalid coupon. Please insert a valid one."
		}
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing json.")
	}

	fmt.Fprintf(w, string(jsonData))
}

func CouponServiceCaller(urlMicroservice, coupon string) Result {
	values := url.Values{}
	values.Add("coupon", coupon)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Coupon server is offline!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Coupon server is offline!")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result

}

func CardServiceCaller(urlMicroservice, ccName, ccNumber, ccExpiration, ccCvv string) Result {
	values := url.Values{}
	values.Add("ccName", ccName)
	values.Add("ccNumber", ccNumber)
	values.Add("ccExpiration", ccExpiration)
	values.Add("ccCvv", ccCvv)

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5

	res, err := retryClient.PostForm(urlMicroservice, values)
	if err != nil {
		result := Result{Status: "Card server is offline!"}
		return result
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Card server is offline!")
	}

	result := Result{}

	json.Unmarshal(data, &result)

	return result
}
