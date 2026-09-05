# 2026-09-05

Benchmark ran using:
* validator-go-proto: d046119c4f0ef48dd2c843d30a65eb02e7a9f3c2
* validator-testsuite: f3ab0e5d2f81ccc89acb91c13f52408a776d6b6b

```
(cd ./src/katydid.org.za/go/validator-go-proto && make paper_benchmarks)
go test -v -run=none -bench=BenchmarkSuite/Conf ./testsuite/intern/
goos: darwin
goarch: arm64
pkg: katydid.org.za/go/validator-go-proto/testsuite/intern
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14         	  206952	      5715 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14         	  184242	      5995 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14   	   43018	     27426 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                      	  218264	      5220 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14           	  204454	      5448 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14     	   43784	     26968 ns/op
PASS
ok  	katydid.org.za/go/validator-go-proto/testsuite/intern	8.678s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=100x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: katydid.org.za/go/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14         	     100	      1055 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14         	     100	      1199 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14   	     100	      1998 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                      	     100	       433.3 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14           	     100	       514.6 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14     	     100	      1153 ns/op
PASS
ok  	katydid.org.za/go/validator-go-proto/testsuite/mem	0.793s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=1000x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: katydid.org.za/go/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14         	    1000	       239.0 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14         	    1000	       392.8 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14   	    1000	       589.5 ns/op
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14                  	    1000	       922.8 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14       	    1000	      1096 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14 	    1000	      1183 ns/op
PASS
ok  	katydid.org.za/go/validator-go-proto/testsuite/mem	0.792s
go test -v -run=none -bench=BenchmarkSuite/Conf -benchtime=10000x ./testsuite/mem/
goos: darwin
goarch: arm64
pkg: katydid.org.za/go/validator-go-proto/testsuite/mem
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14         	   10000	       876.9 ns/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14         	   10000	       910.0 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14   	   10000	      1115 ns/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                      	   10000	       223.4 ns/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14           	   10000	       308.0 ns/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14     	   10000	       419.5 ns/op
PASS
ok  	katydid.org.za/go/validator-go-proto/testsuite/mem	0.839s
go test -v -run=none -bench=BenchmarkSuite/Conf ./testsuite/auto/
goos: darwin
goarch: arm64
pkg: katydid.org.za/go/validator-go-proto/testsuite/auto
cpu: Apple M4 Pro
BenchmarkSuite
BenchmarkSuite/ConfIs2026Json
BenchmarkSuite/ConfIs2026Json-14         	 1420489	       852.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuite/ConfIsComputerScienceJson
BenchmarkSuite/ConfIsComputerScienceJson-14         	 1376670	       874.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUJson-14   	 1000000	      1037 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuite/ConfIs2026Pb
BenchmarkSuite/ConfIs2026Pb-14                      	 7992178	       146.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuite/ConfIsComputerSciencePb
BenchmarkSuite/ConfIsComputerSciencePb-14           	 5197860	       230.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb
BenchmarkSuite/ConfIsIn2026OrLate2025AndEUPb-14     	 3706016	       322.8 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	katydid.org.za/go/validator-go-proto/testsuite/auto	10.312s
```