package token

import "regexp"

var vnDiacriticRegexs = []struct {
	regStr     string
	replaceStr string
}{
	{"à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ", "a"},
	{"è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ", "e"},
	{"ì|í|ị|ỉ|ĩ", "i"},
	{"ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ", "o"},
	{"ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ", "u"},
	{"ỳ|ý|ỵ|ỷ|ỹ", "y"},
	{"đ", "d"},
}

func removeVnDiacritics(s string) string {
	res := s
	for _, diac := range vnDiacriticRegexs {
		r := regexp.MustCompile(diac.regStr)
		res = r.ReplaceAllString(res, diac.replaceStr)
	}
	return res
}

func keywordsWithoutDiacritic(keywords map[string]TokenType) map[string]TokenType {
	for key, tok := range keywords {
		keywords[removeVnDiacritics(key)] = tok
	}
	return keywords
}
