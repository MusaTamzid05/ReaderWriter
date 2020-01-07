package reader_writer

import "log"

func IsError(err error) bool {

	if err != nil {
		log.Println(err)
		return true
	}

	return false
}
