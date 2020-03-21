package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

func main() {
	strbyte, err := ioutil.ReadFile("/home/jeff/Downloads/google.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	str := string(strbyte)
	str = strings.ReplaceAll(str, "\"", " ")
	str = strings.ReplaceAll(str, "'", " ")
	strfield := strings.Fields(str)
	var links []string
	for _, v := range strfield {
		if strings.Contains(v, "http://") || strings.Contains(v, "https://") {
			links = append(links, v)
		}
	}
	for _, v := range links {
		fmt.Println(v)
	}
}