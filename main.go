package main

import (
	"flag"
	"log"
	"strings"
)

var (
	path = flag.String("path", "", "path to watch")
	r    = flag.Bool("r", false, "is recursive watching needed")
	ext  = flag.String("ext", "go", `file extensions to watch. can be a list with coma separator.`)
)

func main() {
	flag.Parse()
	if len(*path) > 0 {
		extArr := strings.Split(*ext, ",")
		extMap := map[string]bool{}
		for _, p := range extArr {
			if len(p) > 0 {
				extMap[strings.TrimSpace(p)] = true
			}
		}

		log.Printf("extensions: %v", extMap)
		watchPath(*path, *r, true, extMap)
	}
}
