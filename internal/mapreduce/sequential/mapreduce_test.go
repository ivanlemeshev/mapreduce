package sequential_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ivanlemeshev/mapreduce/internal/mapreduce/funcs/wordcount"
	"github.com/ivanlemeshev/mapreduce/internal/mapreduce/sequential"
)

const (
	testInput  = "./test_input.txt"
	testOutput = "./test_output.txt"
)

func TestMapReduce_Process(t *testing.T) {
	mr := sequential.New(wordcount.New())

	err := mr.Process(testInput, testOutput)
	require.NoError(t, err)

	output, err := os.ReadFile(testOutput)
	require.NoError(t, err)
	require.Equal(t, 36, len(output))

	fileContent := string(output)

	require.Equal(t, 4, strings.Count(fileContent, "\n"))

	require.True(t, strings.Contains(fileContent, "Hello\t3"))
	require.True(t, strings.Contains(fileContent, "World\t2"))
	require.True(t, strings.Contains(fileContent, "again\t1"))
	require.True(t, strings.Contains(fileContent, "MapReduce\t1"))

	_ = os.Remove(testOutput)
}
