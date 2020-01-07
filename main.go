package main

import "reader_writer/reader_writer"

func main() {
	reader := reader_writer.Reader{}
	reader.Read("./main.go")
	reader.Show()
	reader.Replace(9, "   reader.Test()")
	reader.Show()

	reader_writer.WriteDataTo("./test.txt", reader.Lines)
}
