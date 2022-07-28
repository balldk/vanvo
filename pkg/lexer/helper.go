package lexer

import (
	"vila/pkg/token"
)

var VNALPHA = arrToMap([]rune("aAàÀảẢãÃáÁạẠăĂằẰẳẲẵẴắẮặẶâÂầẦẩẨẫẪấẤậẬbBcCdDđĐeEèÈẻẺẽẼéÉẹẸêÊềỀểỂễỄếẾệỆfFgGhHiIìÌỉỈĩĨíÍịỊjJkKlLmMnNoOòÒỏỎõÕóÓọỌôÔồỒổỔỗỖốỐộỘơƠờỜởỞỡỠớỚợỢpPqQrRsStTuUùÙủỦũŨúÚụỤưƯừỪửỬữỮứỨựỰvVwWxXyYỳỲỷỶỹỸýÝỵỴzZ"))

func appendToken(tok1 *token.Token, tok2 token.Token) {
	tok1.Literal = append(tok1.Literal, ' ')
	tok1.Literal = append(tok1.Literal, tok2.Literal...)
	tok1.Type = token.LookupKeyword(tok1.Literal)
}

func arrToMap(arr []rune) map[rune]bool {
	m := make(map[rune]bool)
	for _, each := range arr {
		m[each] = true
	}
	return m
}

func isVietnameseLetter(ch rune) bool {
	if _, ok := VNALPHA[ch]; ok {
		return true
	}
	return false
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || isVietnameseLetter(ch) || isDigit(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}
