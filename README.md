# Implementing MapReduce in Go

An implementation of MapReduce framework in Go to learn how it works and
and understand the concepts behind it.

It is based on the paper
[MapReduce: Simplified Data Processing on Large Clusters](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf)
by Jeffrey Dean and Sanjay Ghemawat from Google.

I use word count as the example to implement MapReduce.

- [x] Sequential implementation
- [ ] Concurrent implementation
- [ ] Distributed implementation

## Run the sequential implementation

```bash
go run main.go -input input.txt -output output.txt
```
