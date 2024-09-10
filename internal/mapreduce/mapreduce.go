package mapreduce

// KeyValue represents an intermediate key-value pair in the MapReduce process.
type KeyValue struct {
	Key   string
	Value string
}

// MapReducer that represents the Map and Reduce functions.
type MapReducer interface {
	// Map processes the input key-value pair and returns a list of intermediate
	// key-value pairs.
	Map(key, value string) []KeyValue

	// Reduce processes the intermediate key and a list of values associated with
	// that key and returns the final output value.
	Reduce(key string, values []string) string
}
