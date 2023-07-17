package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func simpleRequest(Url string, client *http.Client, tr *http.Transport) (string, string, string) {
	var response string
	var scode string
	var headers string

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return response, scode, headers
	}
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		headers = headerDump(resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		scode = strconv.Itoa(resp.StatusCode)
		response = string(body)
	}

	return response, scode, headers
}

func sameResponse(url string) bool {
	state := true
	rsp1 := ""
	rsp2 := ""

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			rsp1 = string(body)
		}
	}

	client2 := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req2, err := http.NewRequest("GET", fuzzUrl(url), nil)
	if err == nil {
		resp, err := client2.Do(req2)
		if err == nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			rsp2 = string(body)
		}
	}

	if rsp1 != rsp2 {
		state = false
	}
	return state
}

func requesterNoRedarict(url string) (response string) {
	response = ""
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			scode := strconv.Itoa(resp.StatusCode)
			if scode == "200" {
				response = string(body)
			}

		}
	}

	return response

}
func headerDump(headers http.Header) string {
	var sb strings.Builder
	for name, values := range headers {
		for _, value := range values {
			head := name + ": " + value + "\n"
			sb.WriteString(head)
		}
	}
	return sb.String()
}

func createHTTPClient() (*http.Client, *http.Transport) {
	var tr *http.Transport

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: tr,
	}
	return client, tr
}
