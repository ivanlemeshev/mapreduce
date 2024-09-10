package wordcount_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ivanlemeshev/mapreduce/internal/wordcount"
)

func TestMap(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		output int
	}{
		{
			name:   "empty",
			input:  "",
			output: 0,
		},
		{
			name:   "one word",
			input:  "hello",
			output: 1,
		},
		{
			name:   "two words",
			input:  "hello world",
			output: 2,
		},
		{
			name:   "two words with punctuation",
			input:  "hello, world",
			output: 2,
		},
		{
			name:   "two words with numbers",
			input:  "hello 123 world",
			output: 2,
		},
		{
			name:   "two words with special characters",
			input:  "hello @#$%^&*() world",
			output: 2,
		},
		{
			name:   "two same words",
			input:  "hello hello",
			output: 2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			wc := wordcount.New()
			kvs := wc.Map(tc.name, tc.input)
			require.Equal(t, len(kvs), tc.output)
		})
	}
}

func TestReduce(t *testing.T) {
	tt := []struct {
		name   string
		input  []string
		output string
	}{
		{
			name:   "empty",
			input:  []string{},
			output: "0",
		},
		{
			name:   "one word",
			input:  []string{"1"},
			output: "1",
		},
		{
			name:   "two words",
			input:  []string{"1", "1"},
			output: "2",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			wc := wordcount.New()
			result := wc.Reduce(tc.name, tc.input)
			require.Equal(t, result, tc.output)
		})
	}
}
