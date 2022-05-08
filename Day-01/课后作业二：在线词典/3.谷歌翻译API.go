package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strings"
)

func CN2EN(text string) (string, error) {
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&dt=t&sl=zh-CN&tl=en&q=%s", url2.QueryEscape(text))
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	//字符串拆解response
	str := string(bs)
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.ReplaceAll(str, "null", "")
	str = strings.Trim(str, `"`)
	ps := strings.Split(str, `","`)
	return ps[0], nil
}

func main() {
	str, err := CN2EN("作业")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}
