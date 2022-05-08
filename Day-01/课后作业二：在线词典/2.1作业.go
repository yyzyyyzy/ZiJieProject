package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type DictRep struct {
	ReturnPhrase []string `json:"returnPhrase"`
	Query        string   `json:"query"`
	ErrorCode    string   `json:"errorCode"`
	L            string   `json:"l"`
	TSpeakURL    string   `json:"tSpeakUrl"`
	Web          []struct {
		Value []string `json:"value"`
		Key   string   `json:"key"`
	} `json:"web"`
	RequestID   string   `json:"requestId"`
	Translation []string `json:"translation"`
	Dict        struct {
		URL string `json:"url"`
	} `json:"dict"`
	Webdict struct {
		URL string `json:"url"`
	} `json:"webdict"`
	Basic struct {
		ExamType   []string `json:"exam_type"`
		UsPhonetic string   `json:"us-phonetic"`
		Phonetic   string   `json:"phonetic"`
		UkPhonetic string   `json:"uk-phonetic"`
		UkSpeech   string   `json:"uk-speech"`
		Explains   []string `json:"explains"`
		UsSpeech   string   `json:"us-speech"`
	} `json:"basic"`
	IsWord   bool   `json:"isWord"`
	SpeakURL string `json:"speakUrl"`
}

func dictQuery(word string) {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`q=%s&from=en&to=zh-CHS`, word))
	req, err := http.NewRequest("POST", "https://aidemo.youdao.com/trans", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "aidemo.youdao.com")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6,ja;q=0.5")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", "https://ai.youdao.com")
	req.Header.Set("referer", "https://ai.youdao.com/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Microsoft Edge";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse DictRep
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(word)
	for _, item := range dictResponse.Basic.Explains {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 { //os.Args()第一个参数为默认路径，第二个参数为用户输入的参数
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD example: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	dictQuery(word)
}
