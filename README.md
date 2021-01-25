# Reflection vs Json

Run tests by using `go test -bench=.`

## Results


```
goos: darwin
goarch: amd64
pkg: reflection-vs-json
BenchmarkBig/StructToMap-12             1000000000               0.000003 ns/op
BenchmarkBig/JsonToMap-12               1000000000               0.000090 ns/op
BenchmarkNested/StructToMap-12          1000000000               0.000003 ns/op
BenchmarkNested/JsonToMap-12            1000000000               0.0104 ns/op
BenchmarkStructToMap-12                 1000000000               0.000004 ns/op
BenchmarkJsonToMap-12                   1000000000               0.000019 ns/op
PASS
```

Meaning - reflection is way faster than marshal/unmarshal.
