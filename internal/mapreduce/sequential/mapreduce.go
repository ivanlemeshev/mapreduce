package sequential

import (
	"fmt"
	"io"
	"os"

	"github.com/ivanlemeshev/mapreduce/internal/mapreduce"
)

// MapReduce implements the sequential MapReduce algorithm.
type MapReduce struct {
	mapReducer mapreduce.MapReducer
}

// New creates a new MapReduce instance.
func New(mapReducer mapreduce.MapReducer) *MapReduce {
	return &MapReduce{
		mapReducer: mapReducer,
	}
}

// Process reads the input file, processes it using the Map and Reduce
// functions, and writes the output to the output file.
func (mr *MapReduce) Process(inputFilepath, outputFilepath string) error {
	input, err := os.Open(inputFilepath)
	if err != nil {
		return fmt.Errorf(
			"failed to open the input file %q: %w",
			inputFilepath,
			err,
		)
	}
	defer input.Close()

	output, err := os.Create(outputFilepath)
	if err != nil {
		return fmt.Errorf(
			"failed to create the output file %q: %w",
			outputFilepath,
			err,
		)
	}

	data, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf(
			"failed to read the input file %q: %w",
			inputFilepath,
			err,
		)
	}

	kvs := mr.mapReducer.Map(inputFilepath, string(data))

	shuffled := mr.shuffle(kvs)

	for key, values := range shuffled {
		result := mr.mapReducer.Reduce(key, values)
		_, err := fmt.Fprintf(output, "%s\t%s\n", key, result)
		if err != nil {
			return fmt.Errorf(
				"failed to write the output file %q: %w",
				outputFilepath,
				err,
			)
		}
	}

	return nil
}

func (mr *MapReduce) shuffle(kvs []mapreduce.KeyValue) map[string][]string {
	result := make(map[string][]string)
	for _, kv := range kvs {
		result[kv.Key] = append(result[kv.Key], kv.Value)
	}

	return result
}
