# 2026-07-30

Benchmark ran using:
* validator-go-proto: 1dea64be11dec4d75de2cbbcf8a4de8c8da704fd
* validator-testsuite: a2ffd57196cffe6af0038e8e9b413284662667bc

```
% make paper_benchmarks
go test -v -run=none -bench=BenchmarkSuite/Conf ./testsuite/intern/
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                  198308              5926 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14               177614              6198 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14          41535             28587 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                            210386              5341 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                 193557              5668 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14            41914             27558 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/intern  9.358s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=100x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                     100              1192 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14                  100              1288 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14            100              2146 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                               100               605.0 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                    100               497.5 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14              100              1254 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/mem     0.904s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=1000x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                    1000               902.9 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14                 1000              1046 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14           1000              1223 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                              1000               285.8 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                   1000               326.9 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14             1000               531.7 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/mem     0.829s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=10000x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                   10000               919.1 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14                10000               952.6 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14          10000              1157 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                             10000               214.4 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                  10000               322.1 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14            10000               446.7 ns/op
PASS
ok      github.com/katydid/validator-go-proto/testsuite/mem     0.888s
go test -v -run=none -bench=BenchmarkSuite/Conf ./testsuite/auto/
goos: darwin
goarch: arm64
pkg: github.com/katydid/validator-go-proto/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                 1348864               887.0 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14              1313613               902.7 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14        1000000              1076 ns/op               0 B/op          0 allocs/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                           7903616               150.7 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14                4978732               239.2 ns/op             0 B/op          0 allocs/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14          3551433               350.2 ns/op             0 B/op          0 allocs/op
PASS
```