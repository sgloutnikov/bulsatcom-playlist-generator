package model

type BSC struct {
	Channel  string `json:"channel"`
	EPG_Id   string `json:"epg_id"`
	EPG_Name string `json:"epg_name"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Logo     string `json:"logo_selected"`
	Sources  string `json:"sources"`
}

type BSCChannels []BSC

func SetSourcesUrl(newUrl string, channel *BSC) {
	channel.Sources = newUrl
}