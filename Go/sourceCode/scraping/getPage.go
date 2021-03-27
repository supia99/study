package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"./model"

	"github.com/PuerkitoBio/goquery"
)

const (
	// CD情報を取得する元となるURL
	RESOURCE_URL             = "https://www.lantis.jp/imas/"
	RELEASE_PAGE_PARTIAL_URL = "release.html"
	SAVED_PAGE_PATH          = "/mnt/d/down_load/gdrive/github/Go/sourceCode/scraping/htmlPage/"
)

var (
	// buf    bytes.Buffer
	logger = log.New(os.Stderr, "logger:", log.Lshortfile)
)

// urlにgetリクエストを投げて、ページを取得する
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

// スクレイピングする
func GetCDPartialURL(fileContent string) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(fileContent))
	if err != nil {
		return nil, err
	}

	// cd areaにおける、CD単体のページURLを取得する
	selectionCDDiv := doc.Find("div.CD_area")
	partialURLs := make([]string, selectionCDDiv.Find("a").Size())
	selectionCDDiv.Find("a").Each(func(index int, s *goquery.Selection) {
		// // log
		// fmt.Println("href ", index, " :")
		text, _ := s.Attr("href")
		// fmt.Println(text)
		partialURLs[index] = text
	})

	return partialURLs, nil
}

// CD単体のページをスクレイピングする
func ScrapeDiskPage(fileContent string) error {

	disk := new(model.Disk)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(fileContent))
	if err != nil {
		return err
	}

	// fmt.Println("title:", doc.Find("title").Text())
	titleDatas := strings.Split(doc.Find("div.titles").Find("h2").Text(), "\n")
	switch len(titleDatas) {
	case 1:
		disk.Series = strings.Trim(titleDatas[0], " ")
	case 2:
		if "" == strings.Trim(titleDatas[1], " ") {
			disk.Series = strings.Trim(titleDatas[0], " ")
		} else {
			disk.Series = strings.Trim(titleDatas[1], " ")
		}
	case 3:
		disk.Series = strings.Trim(titleDatas[1], " ")
	case 4:
		disk.Series = strings.Trim(titleDatas[1], " ")
	}
	disk.Series = strings.TrimPrefix(disk.Series, "タイトル：")
	// // log
	// fmt.Println("!title:", disk.Series)

	selection := doc.Find("div.release_contents") //pセレクタの内容を取得

	// fmt.Println("children:", selection.Children().Size())
	var songInfos []string
	// trueの場合： index(曲情報)のselection
	selectionFlag := false
	selection.Children().Each(func(index int, child *goquery.Selection) {
		// fmt.Print(index, selectionFlag, ":")
		// fmt.Println(goquery.NodeName(child))

		if "h3" == goquery.NodeName(child) {
			attr, isAttr := child.Find("img").Attr("alt")
			if true == isAttr {
				if "index" == attr {
					selectionFlag = true
				} else if "info" == attr {
					selectionFlag = false
				}
			}
		} else if true == selectionFlag {
			songInfos = append(songInfos, strings.Split(child.Text(), "\n")...)
		}
	})

	// 邪魔な文字列を除外し、曲情報をstring配列に格納する
	for i := len(songInfos) - 1; 0 <= i; i-- {
		songInfos[i] = strings.Trim(songInfos[i], " ")
		// fmt.Println(i, "/", len(songInfos), "target:", songInfos[i])

		// strings.Contains(songInfos[i], "※") ||
		if "" == songInfos[i] || strings.Contains(songInfos[i], "他") || strings.Contains(songInfos[i], "ほか、") || strings.Contains(songInfos[i], "収録") || strings.Contains(songInfos[i], "生産") {
			if 0 == i {
				songInfos = songInfos[1:]
			} else if len(songInfos)-1 == i {
				songInfos = songInfos[:i]
			} else {
				songInfos = append(songInfos[:i], songInfos[i+1:]...)
			}
		}
	}
	// // log
	// for i := 0; i < len(songInfos); i++ {
	// 	fmt.Println(i, "|", songInfos[i], "|", len(songInfos[i]))
	// }

	startIndex := -1
	for index, songInfo := range songInfos {
		if -1 == startIndex {
			startIndex = index
		} else if containsString(songInfo, []string{"作詞", "作曲", "編曲", "歌", "CV"}) {
			//nop
		} else {
			// fmt.Println("  !else index:", startIndex, index)
			disk.Songs = append(disk.Songs, addSongInfo(songInfos[startIndex:index]))
			startIndex = index
		}
	}
	disk.Songs = append(disk.Songs, addSongInfo(songInfos[startIndex:]))

	disk.PrintExcel()
	return nil
}

const (
	COMPOSER_KEYWORD = "作詞"
	LYRIST_KEYWORD   = "作曲"
	ARRANGER_KEYWORD = "編曲"
	SINGER_KEYWORD   = "歌"
)

// 1曲分の情報からCDモデルを作成する
func addSongInfo(infos []string) (song model.Song) {

	// fmt.Println("   !!!", strings.Join(infos, "|"))
	song.Name = infos[0]

	for _, info := range infos {
		splitedInfo := strings.Split(info, "：")

		// if 1 == len(splitedInfo) {
		// 	continue
		// }
		if strings.Contains(info, COMPOSER_KEYWORD) {
			song.Composer = deleteSuffixes(splitedInfo[searchIndex(splitedInfo, COMPOSER_KEYWORD)+1],
				[]string{COMPOSER_KEYWORD, LYRIST_KEYWORD, ARRANGER_KEYWORD, SINGER_KEYWORD})
		}
		if strings.Contains(info, LYRIST_KEYWORD) {
			song.Lyrist = deleteSuffixes(splitedInfo[searchIndex(splitedInfo, LYRIST_KEYWORD)+1],
				[]string{COMPOSER_KEYWORD, LYRIST_KEYWORD, ARRANGER_KEYWORD, SINGER_KEYWORD})
		}
		if strings.Contains(info, ARRANGER_KEYWORD) {
			song.Arranger = deleteSuffixes(splitedInfo[searchIndex(splitedInfo, ARRANGER_KEYWORD)+1],
				[]string{COMPOSER_KEYWORD, LYRIST_KEYWORD, ARRANGER_KEYWORD, SINGER_KEYWORD})
		}
		if strings.Contains(info, SINGER_KEYWORD) {
			str := strings.Join((splitedInfo[searchIndex(splitedInfo, SINGER_KEYWORD)+1:]), "")
			song.Singer = str //eleteSuffixes(str,
			// []string{COMPOSER_KEYWORD, LYRIST_KEYWORD, ARRANGER_KEYWORD, SINGER_KEYWORD})
		}

		// fmt.Println(" index:", strings.Index(info, "CV"), info)
		if 0 == strings.Index(info, "[") {
			song.Singer += info
			// fmt.Println("  !!!singer add:", song.Singer, strings.Index(info, "["))
		}
	}

	return
}

// substrsの単語が1つでもtargetに含まれていれば、trueを返す
func containsString(target string, substrs []string) bool {
	if nil == substrs {
		return false
	}
	for _, substr := range substrs {
		if strings.Contains(target, substr) {
			return true
		}
	}
	return false
}

// strs から、subStrを含む添字を取得する
// 存在しなければ、-1を返す
func searchIndex(strs []string, subStr string) int {
	//fmt.Println("  !!! searchIndex:", strings.Join(strs, "|"))
	for index, str := range strs {
		if strings.Contains(str, subStr) {
			if index == len(strs)-1 {
				// fmt.Println("  !!! searchIndex:", strings.Join(strs, "|"))
				return len(strs) - 2
			} else {
				return index
			}

		}
	}

	return -1
}

// target に含まれる subStrと" "を削除する
func deleteSuffixes(target string, subStr []string) string {

	for _, str := range subStr {
		target = strings.Split(target, str)[0]
	}

	return strings.Trim(target, " ")
}

// ページデータを取得し、そのページからCDデータを取得し、ファイルに保存する。
func getPageData() {

	// releaseページをgetリクエストで取得する
	page, err := GetPage(strings.Join([]string{RESOURCE_URL, RELEASE_PAGE_PARTIAL_URL}, ""))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var partialURLs []string
	partialURLs, err = GetCDPartialURL(page)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//
	for _, partialURL := range partialURLs {
		page, err = GetPage(strings.Join([]string{RESOURCE_URL, partialURL}, ""))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = OutputToFile(page, strings.Join([]string{SAVED_PAGE_PATH, partialURL}, ""))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

// 特定フォルダ以下のファイルを読み込み、CDデータを出力する
func extractData() {

	// ディレクトリ以下のファイル名を読む
	files, err := ioutil.ReadDir(SAVED_PAGE_PATH)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// ファイルごとに読み込む
	for _, file := range files {
		// fmt.Println(index, ":", file.Name())

		// ダウンロード済みのページ情報を読み込む
		// //log
		// fmt.Println("!URL:", file.Name())
		page, err := InputFile(strings.Join([]string{SAVED_PAGE_PATH, file.Name()}, ""))
		if err != nil {
			logger.Print(err.Error())
		}

		err = ScrapeDiskPage(page)
		if err != nil {
			logger.Print(err.Error())
		}
	}
}

func editPage(pageName, old, new string) error {

	// ファイル読み込みテスト
	page, err := InputFile(strings.Join([]string{SAVED_PAGE_PATH, pageName}, ""))
	if err != nil {
		return err
	}

	// fmt.Println("EDIT(", old, "->", new, "):", strings.Contains(page, old))
	page = strings.Replace(page, old, new, -1)

	err = OutputToFile(page, strings.Join([]string{SAVED_PAGE_PATH, pageName}, ""))
	if err != nil {
		return err
	}

	return nil
}

func editHtmls() error {

	// 変更箇所が書かれたデータを読み込む
	page, err := InputFile(strings.Join([]string{SAVED_PAGE_PATH, "../changeData.csv"}, ""))
	if err != nil {
		return err
	}

	for _, splited := range strings.Split(page, "\n") {
		// fmt.Println(index, strings.Split(splited, "|")[2])
		changeDatas := strings.Split(splited, "|")
		// 最後の要素に改行が残ってしまうため、
		if 3 > len(changeDatas) {
			// fmt.Println("edit skip:", changeDatas, len(changeDatas))
			continue
		}

		err = editPage(changeDatas[0], changeDatas[1], changeDatas[2])
		if err != nil {
			logger.Print(err.Error())
			os.Exit(1)
		}
	}

	return nil
}

// func main() {
//
// 	// // ファイル読み込みテスト
// 	// page, err := InputFile(strings.Join([]string{SAVED_PAGE_PATH, "release_LACA-15436.html"}, ""))
// 	// if err != nil {
// 	// 	logger.Print(err.Error())
// 	// }
//
// 	// getPageData()
//
// 	editHtmls()
// 	extractData()
//
// 	// err = ScrapeDiskPage(page)
// 	// if err != nil {
// 	// 	logger.Print(err.Error())
// 	// }
//
// 	// err = OutputToFile(page, "/mnt/d/down_load/gdrive/github/Go/sourceCode/scraping/test.html")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// 	os.Exit(1)
// 	// }
//
// }
