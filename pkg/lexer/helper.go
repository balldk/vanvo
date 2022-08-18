package lexer

import (
	"vanvo/pkg/token"
)

var VNALPHA = arrToMap([]rune("aAàÀảẢãÃáÁạẠăĂằẰẳẲẵẴắẮặẶâÂầẦẩẨẫẪấẤậẬbBcCdDđĐeEèÈẻẺẽẼéÉẹẸêÊềỀểỂễỄếẾệỆfFgGhHiIìÌỉỈĩĨíÍịỊjJkKlLmMnNoOòÒỏỎõÕóÓọỌôÔồỒổỔỗỖốỐộỘơƠờỜởỞỡỠớỚợỢpPqQrRsStTuUùÙủỦũŨúÚụỤưƯừỪửỬữỮứỨựỰvVwWxXyYỳỲỷỶỹỸýÝỵỴzZ"))

func mergeToken(tok1 token.Token, tok2 token.Token) token.Token {
	if len(tok2.Literal) == 0 {
		return tok1
	}
	tok := token.Token{}
	tok.Literal = tok1.Literal
	tok.Literal = append(tok.Literal, ' ')
	tok.Literal = append(tok.Literal, tok2.Literal...)
	tok.Type = token.LookupKeyword(tok.Literal)

	return tok
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
