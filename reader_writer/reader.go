package reader_writer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

type Reader struct {
	Lines []string
}

func (r *Reader) Read(path string) {

	strData := r.read(path)

	if strData == "" {
		log.Printf("Could not read %s.", path)
		return
	}

	r.Lines = strings.Split(strData, "\n")
	fmt.Printf("Total %d Lines read.\n", len(r.Lines))

}

func (r *Reader) read(path string) string {
	data, err := ioutil.ReadFile(path)
	if IsError(err) {
		return ""
	}

	return string(data)

}

func (r *Reader) Show() {

	for index, line := range r.Lines {
		fmt.Printf("\n%d => %s", index+1, line)
	}
}

func (r *Reader) getReplacement(line ,newValue string) string {


	key := strings.Split(line , ":")[0]
	newLine := key + " : " + newValue
	return newLine
}

func (r *Reader) ReplaceWith(data map[string]string) {

	train_reader_found := false

	for index  , line := range r.Lines {

		if strings.Contains(line , "train_input_reader") {
			train_reader_found = true
		}

		for key , value  := range data {

			currentKey := ""

			if key == "input_path_train" || key == "input_path_eval_test" {
				currentKey = "input_path"
			} else {
				currentKey = key
			}

			if strings.Contains(line , currentKey) == false {
				continue
			}

			if currentKey == "fine_tune_checkpoint" {
				if strings.Contains(line , "fine_tune_checkpoint_type") {
					continue
				}
			}
			currentValue := ""

			if currentKey != key {

				if train_reader_found {
					currentValue = data["input_path_train"]
					train_reader_found = false
				}  else {
					currentValue = data["input_path_eval_test"]
				}

			} else {
				currentValue = value
			}


			_ , err := strconv.ParseInt(currentValue , 10 , 32)

			var newLine string

			if err != nil  {
				newLine = r.getReplacement(line , "\"" +  currentValue + "\"")
			} else {
				newLine = r.getReplacement(line ,   currentValue)
			}
			r.Replace(index + 1 , newLine)
			break
		}
	}
}

func (r *Reader) Replace(lineNumber int, newLine string) error {

	if lineNumber <= 0 || lineNumber > len(r.Lines) {
		return errors.New("Invalid line number " + string(lineNumber))
	}

	r.Lines[lineNumber-1] = newLine
	return nil
}
