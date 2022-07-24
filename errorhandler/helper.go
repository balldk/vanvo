package errorhandler

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func printLine(line int, input string) {

}

// func showError(message string, fromToken, toToken token.Token) {
// 	var buf bytes.Buffer

// 	red := color.New(color.FgHiRed)
// 	blue := color.New(color.FgBlue)
// 	green := color.New(color.FgHiGreen)
// 	white := color.New(color.FgWhite)

// 	if filepath != "" {
// 		green.Fprintln(&buf, "-->", filepath)
// 	}
// }
