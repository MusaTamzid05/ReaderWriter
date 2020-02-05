package main

import "reader_writer/reader_writer"

func main() {
	reader := reader_writer.Reader{}
	reader.Read("./pipeline.config")

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
	reader.Show()
}
