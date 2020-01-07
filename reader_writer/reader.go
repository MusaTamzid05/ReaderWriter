package reader_writer

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Reader struct {
	lines []string
}

func (r *Reader) Read(path string) {

	strData := r.read(path)

	if strData == "" {
		log.Printf("Could not read %s.", path)
		return
	}

	r.lines = strings.Split(strData, "\n")
	fmt.Printf("Total %d lines read.", len(r.lines))

}

func (r *Reader) read(path string) string {
	data, err := ioutil.ReadFile(path)
	if IsError(err) {
		return ""
	}

	return string(data)

}

func (r *Reader) Show() {

	for index, line := range r.lines {
		fmt.Printf("\n%d => %s", index+1, line)
	}
}
