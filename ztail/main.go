package main

import (
	"fmt"
	"os"
)

func ztailAtoi(s string) int {
	res := 0
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			res = res*10 + int(ch-'0')
		} else {
			return -1
		}
	}
	return res
}

func main() {
	numBytes := 0
	fileStartIdx := 1

	if len(os.Args) >= 2 {
		if os.Args[1] == "-c" {
			if len(os.Args) >= 3 {
				numBytes = ztailAtoi(os.Args[2])
				fileStartIdx = 3
			}
		} else if len(os.Args[1]) > 2 && os.Args[1][:2] == "-c" {
			numBytes = ztailAtoi(os.Args[1][2:])
			fileStartIdx = 2
		}
	}

	if numBytes < 0 {
		return
	}

	exitCode := 0
	files := os.Args[fileStartIdx:]

	for i, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("%v\n", err)
			exitCode = 1
			continue
		}

		stat, err := file.Stat()
		if err != nil {
			fmt.Printf("%v\n", err)
			exitCode = 1
			file.Close()
			continue
		}

		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", filename)
		}

		size := stat.Size()
		startIdx := int64(0)
		if size > int64(numBytes) {
			startIdx = size - int64(numBytes)
		}

		_, err = file.Seek(startIdx, 0)
		if err != nil {
			fmt.Printf("%v\n", err)
			exitCode = 1
			file.Close()
			continue
		}

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if n > 0 {
				fmt.Printf("%s", string(buf[:n]))
			}
			if err != nil {
				break
			}
		}
		file.Close()
	}

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}
