package reader_writer

import (
	"fmt"
	"log"
	"os"
)

func WriteDataTo(path string, lines []string) error {

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	for _, line := range lines {
		fmt.Fprintln(f, line)

	}

	err = f.Close()

	if err != nil {
		return err
	}

	log.Println("Successfully written ", path)

	return nil
}
