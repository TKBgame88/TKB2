package constants

import (
	"regexp"
)

/** スペースのパッケージ */
var SPACE_REGEX = regexp.MustCompile(`\s`)

/** 改行のパッケージ */
var INDENT_REGEX = regexp.MustCompile(`\r\n|\n\r|\n|\r`)

/** 数値のパッケージ */
var NUM_REGEX = regexp.MustCompile(`(?i)/d`)

/** 持越しTLチャンネル*/
var CARRY_OVER_REGEX = regexp.MustCompile(`(?mi)^(tl|ｔｌ)(\s\[0-9０-９]{0,2}|[0-9０-９]{0,2})`)

/** 持越しTLチャンネルで時間の正規表現マッチ */
var CARRY_OVER_TIME_REGEX = regexp.MustCompile(`(?i)[0-9０-９]{1,2}(:|：)[0-9０-９]{1,2}`)

/** 持ち越しTLチャンネルで時間以外での秒数記載の正規表現マッチ */
var CARRY_OVER_OTHER_TIME_REGEX = regexp.MustCompile(`(?i)[0-9０-９]{1,2}(秒|s|ｓ)`)

/** 持越しTLチャンネル:基準時間(90秒) */
const DEFAULT_TIME int = 90
