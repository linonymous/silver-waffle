package http

import (
	"bytes"
	"cowin/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	HttpClient *http.Client
	URL    string
	AuthToken  string
}

func InitializeClient(client *http.Client, url, auth string) *Client {
	return &Client{
		HttpClient: client,
		URL:    url,
		AuthToken: auth,
	}
}

func (h *Client) GetBeneficiary() *model.Beneficiary {
	origin := fmt.Sprintf("%s/appointment/beneficiaries", h.URL)
	request, err := http.NewRequest(http.MethodGet, origin, nil)
	if err != nil {
		log.Println("Could not make request for GetSchedules")
		return nil
	}
	request.Header.Add("authority", "cdn-api.co-vin.in")
	request.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	request.Header.Add("accept", "application/json, text/plain, */*")
	request.Header.Add("authorization", fmt.Sprintf("Bearer %s", h.AuthToken))
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("sec-fetch-dest", "empty")
	request.Header.Add("referer", "https://selfregistration.cowin.gov.in/")
	request.Header.Add("accept-language", "en-US,en;q=0.9")

	response, err := h.HttpClient.Do(request)
	if err != nil {
		log.Println("Could not GetSchedules")
		return nil
	}
	if response.StatusCode == http.StatusUnauthorized {
		log.Println("auth token is invalid! Retry with a new token")
		os.Exit(400)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("could not read response for GetSchedules ")
		return nil
	}
	var result model.Beneficiary
	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("could not unmarshal the response for get beneficiary")
	}
	return &result
}

func (h *Client) BookSlot(requestBody *model.ScheduleRequest) *model.ScheduleResponse {
	requestB, _ := json.Marshal(requestBody)
	origin := fmt.Sprintf("%s/appointment/schedule", h.URL)
	request, err := http.NewRequest(http.MethodPost, origin, bytes.NewReader(requestB))
	if err != nil {
		log.Println("Could not make request for GetSchedules")
		return nil
	}
	request.Header.Add("authority", "cdn-api.co-vin.in")
	request.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	request.Header.Add("accept", "application/json, text/plain, */*")
	request.Header.Add("authorization", fmt.Sprintf("Bearer %s", h.AuthToken))
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("sec-fetch-dest", "empty")
	request.Header.Add("referer", "https://selfregistration.cowin.gov.in/")
	request.Header.Add("accept-language", "en-US,en;q=0.9")

	response, err := h.HttpClient.Do(request)
	if err != nil {
		log.Println("Could not GetSchedules")
		return nil
	}
	if response.StatusCode == http.StatusUnauthorized {
		log.Println("auth token is invalid! Retry with new auth token!")
		os.Exit(401)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("could not read response for GetSchedules ")
		return nil
	}
	var result model.ScheduleResponse
	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("could not unmarshal the response for get beneficiary")
		return nil
	}
	return &result
}

func (h *Client) GetSchedules(pin string) *model.Centers {
	date := "11-05-2021"
	origin := fmt.Sprintf("%s/appointment/sessions/calendarByPin?pincode=%s&date=%s", h.URL, pin, date)
	request, err := http.NewRequest(http.MethodGet, origin, nil)
	if err != nil {
		log.Println("Could not make request for GetSchedules")
		return nil
	}
	request.Header.Add("authority", "cdn-api.co-vin.in")
	request.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	request.Header.Add("accept", "application/json, text/plain, */*")
	request.Header.Add("authorization", fmt.Sprintf("Bearer %s", h.AuthToken))
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("sec-fetch-dest", "empty")
	request.Header.Add("referer", "https://selfregistration.cowin.gov.in/")
	request.Header.Add("accept-language", "en-US,en;q=0.9")

	response, err := h.HttpClient.Do(request)
	if err != nil {
		log.Println("Could not GetSchedules")
		return nil
	}
	if response.StatusCode == http.StatusUnauthorized {
		log.Println("auth token is invalid! Retry with new auth token!")
		os.Exit(401)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("could not read response for GetSchedules ")
		return nil
	}
	var result model.Centers
	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("could not unmarshal the response for get beneficiary")
		return nil
	}
	return &result
}


func (h *Client) GetCaptcha() *model.BSCaptcha {
	origin := fmt.Sprintf("%s/auth/getRecaptcha", h.URL)
	request, err := http.NewRequest(http.MethodPost, origin, strings.NewReader("{}"))
	if err != nil {
		log.Println("Could not make request for GetSchedules")
		return nil
	}
	request.Header.Add("authority", "cdn-api.co-vin.in")
	request.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"90\", \"Google Chrome\";v=\"90\"")
	request.Header.Add("accept", "application/json, text/plain, */*")
	request.Header.Add("authorization", fmt.Sprintf("Bearer %s", h.AuthToken))
	request.Header.Add("sec-ch-ua-mobile", "?0")
	request.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	request.Header.Add("sec-fetch-mode", "cors")
	request.Header.Add("sec-fetch-dest", "empty")
	request.Header.Add("referer", "https://selfregistration.cowin.gov.in/")
	request.Header.Add("accept-language", "en-US,en;q=0.9")

	response, err := h.HttpClient.Do(request)
	if err != nil {
		log.Println("Could not GetSchedules")
		return nil
	}
	if response.StatusCode == http.StatusUnauthorized {
		log.Println("auth token is invalid! Retry with new auth token!")
		os.Exit(401)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("could not read response for GetSchedules ")
		return nil
	}
	var result model.BSCaptcha
	if err = json.Unmarshal(body, &result); err != nil {
		log.Println("could not unmarshal the response for get beneficiary")
		return nil
	}
	return &result
}
