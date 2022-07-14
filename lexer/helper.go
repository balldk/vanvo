package lexer

import "vila/token"

var VNALPHA = arrToMap([]rune("aAàÀảẢãÃáÁạẠăĂằẰẳẲẵẴắẮặẶâÂầẦẩẨẫẪấẤậẬbBcCdDđĐeEèÈẻẺẽẼéÉẹẸêÊềỀểỂễỄếẾệỆfFgGhHiIìÌỉỈĩĨíÍịỊjJkKlLmMnNoOòÒỏỎõÕóÓọỌôÔồỒổỔỗỖốỐộỘơƠờỜởỞỡỠớỚợỢpPqQrRsStTuUùÙủỦũŨúÚụỤưƯừỪửỬữỮứỨựỰvVwWxXyYỳỲỷỶỹỸýÝỵỴzZ"))

func appendIdent(ident1 *token.Token, ident2 token.Token) {
	if ident1.Type != token.IDENT || ident2.Type != token.IDENT {
		panic("Cannot merge non-identifier")
	}
	ident1.Literal = append(ident1.Literal, ' ')
	ident1.Literal = append(ident1.Literal, ident2.Literal...)
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
