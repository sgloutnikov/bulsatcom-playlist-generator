package main

import (
    "fmt"
    "encoding/json"
)

type BSC struct {
    Channel string `json:"channel"`
    EPG_Id string `json:"epg_id"`
    EPG_Name string `json:"epg_name"`
    Title string `json:"title"`
    Genre string `json:"genre"`
    Logo_EPG string `json:"logo_epg"`
    Sources string `json:"sources"`
}

func main() {
    fmt.Println("Testor")

    data := []byte(`
    {
    "channel": "101",
    "epg_id": "1",
    "epg_name": "bnt1_hd",
    "title": "БНТ 1 HD",
    "genre": "ЕФИРНИ",
    "audio": "bul",
    "subtitles": "und",
    "logo": "https://api.iptv.bulsat.com//logos/101-BNT%201%20HD.png",
    "logo_selected": "https://api.iptv.bulsat.com//logos/101-BNT%201%20HD_selected.png",
    "logo_favorite": "https://api.iptv.bulsat.com//logos/101-BNT%201%20HD_favorite.png",
    "logo_mobile": "https://api.iptv.bulsat.com//logos/101-BNT%201%20HD_mobile.png",
    "logo_mobile_selected": "https://api.iptv.bulsat.com//logos/101-BNT%201%20HD_mobile_selected.png",
    "logo_epg": "https://api.iptv.bulsat.com//logos_epg/101-BNT%201%20HD_epg.png",
    "radio": false,
    "sources": "https://lb-m.iptv.bulsat.com/redirect/bnt/smil:bnt1_hd_mobile.smil?scheme=m3u8&bulsatCustomParameter=123372_3816411_123372_NULL_android_ce372703cd579e4c4bb0b66fb92a81eb_&bulsatendtime=1504049176&bulsatstarttime=1504012876&bulsathash=hsaN1Zx0gGdaowvxJg5fjOkgsSCMmZjWimiN45DepMU=",
    "lb": true,
    "quality": "360-SMIL",
    "parent": "",
    "offset": "",
    "pg": "free",
    "program": {
      "start": "20170829132500 +0000",
      "stop": "20170829134000 +0000",
      "startts": "1504013100",
      "stopts": "1504014000",
      "title": "Телепазарен прозорец",
      "desc": ""
    }
  }
  `)
    var mychan BSC
    err := json.Unmarshal(data, &mychan)
    if err == nil {
        fmt.Println(mychan.Channel)
        fmt.Println(mychan.EPG_Id)
        fmt.Println(mychan.EPG_Name)
        fmt.Println(mychan.Title)
        fmt.Println(mychan.Genre)
        fmt.Println(mychan.Logo_EPG)
        fmt.Println(mychan.Sources)
    }


}