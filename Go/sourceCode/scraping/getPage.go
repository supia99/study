package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger:", log.Lshortfile)
)

type ScrapingError struct {
	msg string
}

func (err ScrapingError) Error() string {
	return fmt.Sprintf("ScrapingError %s ", err.msg)
}

// urlにgetリクエストを投げる
func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	// logger.Print(string(byteArray))
	return string(byteArray), nil // htmlをstringで取得
}

// contentをfileNameが指すファイルに書き出す
func OutputToFile(content string, fileName string) error {

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(([]byte)(content))
	return nil
}

// fileNameが指すファイルを読み込む
func InputFile(fileName string) (string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var contentBytes []byte
	contentBytes, err = ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}

//
func Scrape(fileContent string) error {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(fileContent))
	if err != nil {
		return err
	}
	fmt.Println("title:", doc.Find("title").Text())
	selection := doc.Find("div.release_contents") //pセレクタの内容を取得
	selection = selection.Find("a")
	fmt.Println(selection.Attr("href"))
	// for i := range selection.Nodes {
	// 	fmt.Println(" |", i, "|", selection.Eq(i).Text())
	// }

	// if selection != nil {
	// 	fmt.Println("=finder=\n", selection.Text())
	// } else {
	// 	fmt.Println("no target selection!")
	// }

	return nil
}

func main() {

	// page, err := GetPage("https://www.lantis.jp/imas/release.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	// err = OutputToFile(page, "/mnt/d/down_load/gdrive/github/Go/sourceCode/scraping/htmlPage/millionRealeaseTop.html")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	page, err := InputFile("/mnt/d/down_load/gdrive/github/Go/sourceCode/scraping/htmlPage/millionRealeaseTop.html")
	if err != nil {
		logger.Print(err.Error())
	}

	// fmt.Println("==file content==\n", page)
	Scrape(page)

}
