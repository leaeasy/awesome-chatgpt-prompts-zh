package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

var (
	jsonFile   = "../prompts-zh.json"
	trimPrefix = []string{"扮演", "担任", "作为", "一个"}
)

type Prompts struct {
	Act    string `json:"act"`
	Cmd    string `json:"cmd"`
	Prompt string `json:"prompt"`
}

func main() {
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var prompts []Prompts

	if err := json.Unmarshal(byteValue, &prompts); err != nil {
		log.Fatal(err)
	}

	// 默认
	a := pinyin.NewArgs()
	// pinyin.FirstLetter
	a.Style = pinyin.FirstLetter

	result := make([]Prompts, 0)

	for _, p := range prompts {
		for _, prefix := range trimPrefix {
			p.Act = strings.TrimPrefix(p.Act, prefix)
		}

		p.Cmd = strings.Join(pinyin.LazyPinyin(p.Act, a), "")
		result = append(result, p)
	}

	file, _ := json.MarshalIndent(result, "", " ")
	ioutil.WriteFile("../prompts-zh-result.json", file, 0644)
}
