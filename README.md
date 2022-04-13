## Simple golang functional programming example with benchmark 

- go run main.go
- go test -bench=.

# Benchmark results:
goos: linux  
goarch: amd64  
pkg: go-fn-prog-01  
cpu: Intel(R) Core(TM) i7-4510U CPU @ 2.00GHz  

Benchmark                               Time(ns)        CPU(ns/op)
----------------------------------------------------------------------
BenchmarkFilterWithFP-4   	            5538358	        225.4 ns/op  
BenchmarkWithloops-4      	            7089026	        173.0 ns/op  
BenchmarkWithOneloop-4    	            25567000	    42.63 ns/op  
PASS  
ok  	go-fn-prog-01	5.043s  
