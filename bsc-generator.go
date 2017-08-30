package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"./model"
	"log"
	"encoding/json"
	"os"
	"bufio"
	"github.com/spf13/viper"
)

var authToken = ""

func main() {
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig()
    check(err)
    fmt.Println("Generating playlist in " + viper.GetString("playlist.path"))

    login(viper.GetString("login.username"), viper.GetString("login.passwordHash"))
    channels := getChannel()
    generatePlaylist(viper.GetString("playlist.path"), channels)

    fmt.Println(authToken)
}

func generatePlaylist(path string, channels model.BSCChannels) {
	file, err := os.Create(path)
	check(err)
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("#EXTM3U\n")

	// Write out available channels
	var buffer bytes.Buffer
	for _, channel := range channels {
		buffer.WriteString("#EXTINF:-1 ")
		buffer.WriteString("tvg-name=\"")
		buffer.WriteString(channel.EPG_Name)
		buffer.WriteString("\" tvg-logo=\"")
		buffer.WriteString(channel.Logo)
		buffer.WriteString("\" group-title=\"")
		buffer.WriteString(channel.Genre)
		buffer.WriteString("\",")
		buffer.WriteString(channel.Title)
		buffer.WriteString("\n")
		buffer.WriteString(channel.Sources)
		buffer.WriteString("\n")

		writer.WriteString(buffer.String())
		buffer.Reset()
	}

	writer.Flush()
}

func getChannel() model.BSCChannels {
	myChannels := model.BSCChannels{}
	request, _ := http.NewRequest(http.MethodGet, "https://api.iptv.bulsat.com/tv/full/live", nil)
	request.Header.Add("SSBULSATAPI", authToken)
	client := http.Client{}

	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(body, &myChannels)

	//Follow redirect urls to get direct url (Some players have problems with indirect links)
	totalChannels := len(myChannels)
	fmt.Printf("Found %d channels.\n", totalChannels)
	for i := range myChannels {
		fmt.Printf("[%d of %d] Getting direct URL for %s.\n", i+1, totalChannels, myChannels[i].Title)
		request, _ = http.NewRequest(http.MethodGet, myChannels[i].Sources, nil)
		response, _ = client.Do(request)
		model.SetSourcesUrl(response.Request.URL.String(), &myChannels[i])
	}

	return myChannels
}

func login(user, passHash string) bool {
	apiUrl := "https://api.iptv.bulsat.com/?auth"

	payload := url.Values{}
	payload.Set("user", user)
	payload.Set("pass", passHash)
	payload.Set("device_id", "123000001")
	payload.Set("device_name", "bscgen1")
	payload.Set("os_version", "1.0")
	payload.Set("os_type", "android")
	payload.Set("app_version", "0.3.6")

	client := http.Client{}
	request, _ := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBufferString(payload.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, _ := client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	check(err)
	authToken = response.Header["Ssbulsatapi"][0]
	logged := model.BSCLogin{}
	err = json.Unmarshal(body, &logged)
	check(err)
	return logged.Logged == "true"
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}