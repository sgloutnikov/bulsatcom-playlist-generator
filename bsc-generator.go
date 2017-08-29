package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var authToken = ""

func main() {
	fmt.Println("Hello world!")

	apiUrl := "https://api.iptv.bulsat.com/?auth"
	payload := url.Values{}
	payload.Set("user", "")
	payload.Set("pass", "")
	payload.Set("device_id", "")
	payload.Set("device_name", "")
	payload.Set("os_version", "")
	payload.Set("os_type", "android")
	payload.Set("app_version", "0.3.6")

	client := http.Client{}
	r, _ := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBufferString(payload.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(resp.Header)

	authtoken := resp.Header["Ssbulsatapi"][0]
	fmt.Println(authtoken)

	// GET
	r, _ = http.NewRequest(http.MethodGet, "https://api.iptv.bulsat.com/tv/full/live", nil)
	r.Header.Add("SSBULSATAPI", authtoken)

	resp, _ = client.Do(r)
	fmt.Println(resp.Status)
	//TODO: Handle JSON Response...
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
    fmt.Println(authtoken)


}
