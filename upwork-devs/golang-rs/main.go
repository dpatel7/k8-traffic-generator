package main

import (
	"github.com/tggg/generateurl"
	"github.com/tggg/log"
)

func main() {

	log.Filelog("urltxt.txt")

	generateurl.Thefile("urltxt.txt", "open")
	generateurl.Thefile("urltxt.txt", "download")
	generateurl.Thefile("urltxt.txt", "upload")

}
