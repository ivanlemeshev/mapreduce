# Implementing MapReduce in Go

An implementation of MapReduce framework in Go to learn how it works and
and understand the concepts behind it.

It is based on the paper
[MapReduce: Simplified Data Processing on Large Clusters](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf)
by Jeffrey Dean and Sanjay Ghemawat from Google.

It uses word count as the example to implement MapReduce.

## Run the sequential implementation

```bash
go run ./cmd/sequential/main.go -input ./data/input.txt -output ./data/output.txt
```
