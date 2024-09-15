package wordcount

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/ivanlemeshev/mapreduce/internal/mapreduce"
)

// wordCount implements the MapReducer interface to count the number
// of occurrences of each word in a dataset.
type wordCount struct{}

// New creates a new wordCount instance.
func New() *wordCount {
	return &wordCount{}
}

// Map accepts a key that identifies the data and a value that is the data.
// It returns a list of key-value pairs where the key is a word and the value
// is "1". It doesn't matter what the value is, but it is a common convention
// to use "1" as the value when counting occurrences of a words. The use of "1"
// as the value allows for easy counting during the reduce phase, where all
// occurrences of each word are aggregated. This approach efficiently counts
// the frequency of each word in large datasets. Words with different cases are
// considered different words.
func (wc *wordCount) Map(key, value string) []mapreduce.KeyValue {
	words := wc.unicodeStringToWords(value)

	kvs := make([]mapreduce.KeyValue, 0, len(words))
	for _, w := range words {
		kv := mapreduce.KeyValue{
			Key:   w,
			Value: "1",
		}
		kvs = append(kvs, kv)
	}

	return kvs
}

// Reduce counts the number of occurrences of each word. The key is the word
// and the values are the number of occurrences of that word.
func (wc *wordCount) Reduce(key string, values []string) string {
	// An individual value is always "1" so we can just count the number of values.
	return strconv.Itoa(len(values))
}

// Split the value into words. It ignores any non-letter characters.
func (wc *wordCount) unicodeStringToWords(value string) []string {
	return strings.FieldsFunc(value, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}
