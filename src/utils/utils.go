package utils

import (
	"fmt"
	constants "go_bot/src/const"
	"regexp"
	"strconv"
	"strings"
)

type Utils struct{}

func (t Utils) Convert(msg string) string {

	// 最初の改行の番目を取得
	index := strings.Index(msg, "\n")
	// TLXXまでを取得する。スペース全削除
	_h := constants.SPACE_REGEX.ReplaceAllString(msg, "")
	// TLのマッチパッケージ
	_r := regexp.MustCompile(`(?mi)(tl|ｔｌ)`)
	// その後tlを削除する
	_time := _r.ReplaceAllString(_h, "")
	var i, _ = strconv.Atoi(_time)

	// デフォルト時間よりマイナスした時間を算出
	var coTime int = constants.DEFAULT_TIME - i

	// TLコマンド以降の文字列
	tls := constants.INDENT_REGEX.Split(msg[index+1:], -1)

	for _, tl := range tls {
		fmt.Printf("[%s]", tl) // -> [Alfa][Bravo][Charlie][Delta][Echo][Foxtrot][Golf]
	}

	fmt.Printf(string(coTime))

	return _time
}
