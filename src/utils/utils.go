package utils

import (
	"fmt"
	constants "pricess_connect_lite_bot/src/const"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	moji "github.com/ktnyt/go-moji"
)

type Utils struct{}

func (t Utils) Convert(msg string, s *discordgo.Session, m *discordgo.MessageCreate) {

	// 最初の改行の番目を取得
	desc := ""
	index := strings.Index(msg, "\n")
	if index != -1 {
		// TLXXまでを取得する。スペース全削除
		_h := constants.SPACE_REGEX.ReplaceAllString(msg[0:index], "")
		// TLのマッチパッケージ
		_r := regexp.MustCompile(`(?mi)(tl|ｔｌ)`)
		// その後tlを削除する
		_time := _r.ReplaceAllString(_h, "")
		var i, _ = strconv.Atoi(_time)

		// デフォルト時間よりマイナスした時間を算出
		var coTime int = constants.DEFAULT_TIME - i

		// TLコマンド以降の文字列
		tls := constants.INDENT_REGEX.Split(msg[index+1:], -1)

		var slice []string

		for _, tl := range tls {
			var isAddFlag bool = true
			var _mainTime = constants.CARRY_OVER_TIME_REGEX.FindAllString(tl, -1)
			var _subTime = constants.CARRY_OVER_OTHER_TIME_REGEX.FindAllString(tl, -1)

			// サブ時間を変換
			if len(_subTime) != 0 {
				for _, st := range _subTime {
					__r := regexp.MustCompile(`[0-9０-９]{0,2}`)
					__org := moji.Convert(__r.FindString(st), moji.ZE, moji.HE)
					var __i, _ = strconv.Atoi(__org)
					_newTime := __i - coTime
					tl = strings.Replace(tl, __org, strconv.Itoa(_newTime), -1)
				}
			}

			if len(_mainTime) != 0 {
				for _, mt := range _mainTime {
					_h = moji.Convert(mt, moji.ZE, moji.HE)
					var _left, _ = strconv.Atoi(strings.Split(_h, ":")[0])
					var _right, _ = strconv.Atoi(strings.Split(_h, ":")[1])

					_t := time.Date(2021, 1, 1, _left, _right, 0, 0, time.UTC)
					_newTime := _t.Add(-time.Duration(coTime) * time.Minute)
					_newLeft := strconv.Itoa(_newTime.Hour())
					_newRight := fmt.Sprintf("%02d", _newTime.Minute())

					if (_newTime.Minute() == 0 && _newTime.Hour() == 0) || _newTime.Hour() >= 2 {
						isAddFlag = false
						break
					}
					tl = strings.Replace(tl, _h, _newLeft+":"+_newRight, -1) + "\n"
				}
				if isAddFlag {
					slice = append(slice, tl)
				}

			} else {
				slice = append(slice, tl)
			}
		}
		if slice == nil {
			desc = "```\n" + strconv.Itoa(i) + "秒では持ち越しTLが組めないぞ!" + "```"
		} else {
			desc = "```\n" + strconv.Itoa(i) + "秒の持ち越しTLだ!\n\n" + strings.Join(slice, "") + "```"
		}
	} else {
		desc = "```\n" + "正しいフォーマットで入力してください。" + "```"
	}

	s.ChannelMessageSend(m.ChannelID, desc)
}
