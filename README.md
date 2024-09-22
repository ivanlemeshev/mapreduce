# Implementing MapReduce in Go

![Build](https://github.com/ivanlemeshev/mapreduce/actions/workflows/build.yml/badge.svg)
![Test](https://github.com/ivanlemeshev/mapreduce/actions/workflows/test.yml/badge.svg)
![Lint](https://github.com/ivanlemeshev/mapreduce/actions/workflows/lint.yml/badge.svg)
![Sec](https://github.com/ivanlemeshev/mapreduce/actions/workflows/test.yml/badge.svg)

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
