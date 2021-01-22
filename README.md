# Reflection vs Json

Run tests by using `go test -bench=.`

## Results

```
goos: darwin
goarch: amd64
pkg: reflection-vs-json
BenchmarkStructToMap-12         1000000000               0.000005 ns/op
BenchmarkJsonToMap-12           1000000000               0.000047 ns/op
PASS
```

Meaning - reflection is way faster than marshal/unmarshal.
