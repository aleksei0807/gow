package main

import (
	"log"
	"os"
	"os/exec"
	gPath "path"
	"strings"
)

func isTrueExt(name string, ext map[string]bool) bool {
	nameArr := strings.Split(name, ".")
	myExt := nameArr[len(nameArr)-1]
	if ext[myExt] {
		return true
	}
	return false
}

func watchPath(path string, r bool, first bool, ext map[string]bool) {
	var fullpath string
	if first {
		log.Printf("path: %s", path)
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		dir, _ := gPath.Split(path)
		fullpath = gPath.Clean(pwd + "/" + dir)
		log.Printf("fullpath: %s", fullpath)
	} else {
		fullpath = gPath.Clean(path)
	}

	if first || r {
		cmd1 := exec.Command("sh", "-c", "ls -1 "+path)
		out1, err := cmd1.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		paths := strings.Split(string(out1), "\n")
		if len(path) > 0 {
			for _, p := range paths {
				if len(p) > 0 {
					myPath := gPath.Clean(fullpath + "/" + p)
					file, err := os.Open(myPath)

					if err != nil {
						log.Fatal(err)
					}
					defer file.Close()

					fi, err := file.Stat()
					if err != nil {
						log.Fatal(err)
					}
					if fi.IsDir() {
						watchPath(myPath, r, false, ext)
					} else {
						if len(ext) < 1 || isTrueExt(p, ext) {
							log.Printf("file: %s", myPath)
						}
					}
				}
			}
		}
	} else {
		file, err := os.Open(fullpath)

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fi, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		if fi.IsDir() {
			log.Printf("directory: %s", fullpath)
		} else {
			if len(ext) < 1 || isTrueExt(path, ext) {
				log.Printf("file: %s", fullpath)
			}
		}
	}
}
