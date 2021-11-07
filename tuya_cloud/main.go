package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

const ()

var (
	Token string
)

type TokenResponse struct {
	Result struct {
		AccessToken  string `json:"access_token"`
		ExpireTime   int    `json:"expire_time"`
		RefreshToken string `json:"refresh_token"`
		UID          string `json:"uid"`
	} `json:"result"`
	Success bool  `json:"success"`
	T       int64 `json:"t"`
}

func main() {
	var (
		host      string
		clientID  string
		secret    string
		deviceIDs string
	)
	flag.StringVar(&host, "host", "https://openapi.tuyain.com", "")
	flag.StringVar(&clientID, "client_id", "", "")
	flag.StringVar(&secret, "secret", "", "")
	flag.StringVar(&deviceIDs, "device_ids", "", "")
	flag.Parse()

	if clientID == "" || secret == "" || deviceIDs == "" {
		log.Println("client_id, secret and device_ids are required")
		return
	}

	ids := strings.Split(deviceIDs, ",")
	GetToken(host, clientID, secret)
	GetDevices(host, clientID, secret, ids)
}

func GetToken(host, clientID, secret string) {
	method := "GET"
	body := []byte(``)
	req, _ := http.NewRequest(method, host+"/v1.0/token?grant_type=1", bytes.NewReader(body))

	buildHeader(req, body, clientID, secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bs, _ := ioutil.ReadAll(resp.Body)
	ret := TokenResponse{}
	json.Unmarshal(bs, &ret)
	log.Println("resp:", string(bs))

	if v := ret.Result.AccessToken; v != "" {
		Token = v
	}
}

func GetDevices(host, clientID, secret string, deviceIDs []string) {
	for _, deviceID := range deviceIDs {
		method := "GET"
		body := []byte(``)
		req, _ := http.NewRequest(method, host+"/v1.0/devices/"+deviceID, bytes.NewReader(body))

		buildHeader(req, body, clientID, secret)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		bs, _ := ioutil.ReadAll(resp.Body)
		log.Println("resp:", string(bs))
	}

}

func buildHeader(req *http.Request, body []byte, clientID, secret string) {
	req.Header.Set("client_id", clientID)
	req.Header.Set("sign_method", "HMAC-SHA256")

	ts := fmt.Sprint(time.Now().UnixNano() / 1e6)
	req.Header.Set("t", ts)

	if Token != "" {
		req.Header.Set("access_token", Token)
	}

	sign := buildSign(req, body, ts, clientID, secret)
	req.Header.Set("sign", sign)
}

func buildSign(req *http.Request, body []byte, t, clientID, secret string) string {
	headers := getHeaderStr(req)
	urlStr := getUrlStr(req)
	contentSha256 := Sha256(body)
	stringToSign := req.Method + "\n" + contentSha256 + "\n" + headers + "\n" + urlStr
	signStr := clientID + Token + t + stringToSign
	sign := strings.ToUpper(HmacSha256(signStr, secret))
	return sign
}

func Sha256(data []byte) string {
	sha256Contain := sha256.New()
	sha256Contain.Write(data)
	return hex.EncodeToString(sha256Contain.Sum(nil))
}

func getUrlStr(req *http.Request) string {
	url := req.URL.Path
	keys := make([]string, 0, 10)

	query := req.URL.Query()
	for key, _ := range query {
		keys = append(keys, key)
	}
	if len(keys) > 0 {
		url += "?"
		sort.Strings(keys)
		for _, keyName := range keys {
			value := query.Get(keyName)
			url += keyName + "=" + value + "&"
		}
	}

	if url[len(url)-1] == '&' {
		url = url[:len(url)-1]
	}
	return url
}

func getHeaderStr(req *http.Request) string {
	signHeaderKeys := req.Header.Get("Signature-Headers")
	if signHeaderKeys == "" {
		return ""
	}
	keys := strings.Split(signHeaderKeys, ":")
	headers := ""
	for _, key := range keys {
		headers += key + ":" + req.Header.Get(key) + "\n"
	}
	return headers
}

func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
