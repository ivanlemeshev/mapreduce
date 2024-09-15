package main

import (
	"flag"
	"log"

	"github.com/ivanlemeshev/mapreduce/internal/mapreduce/funcs/wordcount"
	"github.com/ivanlemeshev/mapreduce/internal/mapreduce/sequential"
)

var (
	input  = flag.String("input", "", "input file path")
	output = flag.String("output", "", "output file path")
)

func main() {
	flag.Parse()

	mr := sequential.New(wordcount.New())

	err := mr.Process(*input, *output)
	if err != nil {
		log.Fatalln("failed to process the input file:", err)
	}
}
