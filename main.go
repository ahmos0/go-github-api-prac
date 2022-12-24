package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

var grassURL string = loadEnvURL()

func main() {

	getGrass(grassURL)
}

func loadEnvToken() string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".envの読み込みに失敗しました: %v", err)
	}

	token := os.Getenv(("AccessToken"))

	return token
}

func loadEnvURL() string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf(".envの読み込みに失敗しました: %v", err)
	}

	url := os.Getenv(("GrassURL"))

	return url
}

func getGrass(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("urlの取得がうまくできませんでした")
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	regex1 := regexp.MustCompile(`data-count`)
	regex2 := regexp.MustCompile(`data-date`)
	regex3 := regexp.MustCompile(`ry="2" `)

	result := regex3.Split(string(body), -1)
	for i := 0; i < len(result); i++ {
		arr := strings.Split(result[i], " ")
		for j := 0; j < len(arr); j++ {
			if regex2.MatchString(arr[j]) {
				fmt.Println(arr[j])
			} else if regex1.MatchString(arr[j]) {
				fmt.Println(arr[j])
			}

		}
		fmt.Println()
	}
}
