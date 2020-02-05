package main

import (
	"reader_writer/reader_writer"
	"flag"
	"log"
	"strings"
)

func editConfigFrom(src_config , dst_config string) {
	reader := reader_writer.Reader{}
	reader.Read(src_config)

	dataMap := map[string]string {
		"num_classes" : "2" ,
		"batch_size" : "12",
		"num_steps" : "2000",
		"label_map_path" : "/home/musa/python_pro/trainer_package/generated_data/label.pbtxt",
		"num_examples" : "59",
		"input_path_train" : "/home/musa/python_pro/trainer_package/generated_data/train.record",
		"input_path_eval_test" : "/home/musa/python_pro/trainer_package/generated_data/test.record",
	}

	reader.ReplaceWith(dataMap)
	reader_writer.WriteDataTo(dst_config , reader.Lines)
}

func getDstConfigPath(srcPath string) string  {
	return strings.Split(srcPath , ".config")[0] + "_edited.config"
}

func main() {
	
	srcPathPtr := flag.String("src" , "" , "src pipeline config path")


	flag.Parse()

	if *srcPathPtr == ""  {
		log.Fatalln("Usage: -src pipeline.config")
	}

	dstPath := getDstConfigPath(*srcPathPtr)
	editConfigFrom(*srcPathPtr , dstPath)
}
