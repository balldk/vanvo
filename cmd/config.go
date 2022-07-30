package cmd

import "os"

func initConfig() {
	os.Setenv("IDENT_SIZE", "4")
}
