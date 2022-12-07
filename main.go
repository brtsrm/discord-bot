package main

import (
	"log"
	"os"
)

var (
	outfile, _ = os.Create("RSSParse.log")
	l          = log.New(outfile, "", log.LstdFlags|log.Lshortfile)
)

func main() {
	ConnectToDc()
}
