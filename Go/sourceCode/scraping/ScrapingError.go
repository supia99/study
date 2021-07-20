package main

import "fmt"

// 独自エラー構造体
type ScrapingError struct {
	msg string
	err error
}

func (err ScrapingError) Error() string {
	return fmt.Sprintf("ScrapingError %s ", err.msg)
}
