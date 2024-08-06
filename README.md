Benchmark of Go array conversion

```
goos: darwin
goarch: arm64
pkg: github.com/axdoomer/array_conv_bench
BenchmarkArrayConversion/UnsafeConversion-8             1000000000     0.4964 ns/op
BenchmarkArrayConversion/CopyConversion-8               282516         4198 ns/op
BenchmarkArrayConversion/UnsafeConversion#01-8          1000000000     0.4974 ns/op
BenchmarkArrayConversion/CopyConversion#01-8            73630          16417 ns/op
BenchmarkArrayConversion/UnsafeConversion#02-8          1000000000     0.5048 ns/op
BenchmarkArrayConversion/CopyConversion#02-8            34322          34695 ns/op
PASS
ok      github.com/axdoomer/array_conv_bench    7.014s
```

