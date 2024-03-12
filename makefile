build:
	GOEXPERIMENT=rangefunc go build .

run:
	GOEXPERIMENT=rangefunc go run .

test:
	GOEXPERIMENT=rangefunc go test -count=1 ./...
testv:
	GOEXPERIMENT=rangefunc go test -count=1 -v ./...

bench:
	GOEXPERIMENT=rangefunc go test -run=none -bench=. -benchmem ./...

bench_prof:
	GOEXPERIMENT=rangefunc go test -run=none -bench=. -benchmem \
		-cpuprofile=cpu.prof -memprofile=mem.prof ./...

bench_pgo:
	GOEXPERIMENT=rangefunc go test -run=none -bench=. -benchmem \
		-pgo=cpu.prof -cpuprofile=cpu.prof -memprofile=mem.prof ./...
