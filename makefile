draw:
	go run cmd/main.go | \
		gvpr -c -f ./scripts/weight.gvpr | dot -Tpdf | open -f -a Preview

fuzz:
	go test -fuzz=. -count=1 .

test:
	go test -count=1 ./...

testshort:
	go test -short -count=1 ./...

test_cover:
	go test -coverprofile=c.out -count=1 ./...

testv:
	go test -count=1 -v ./...

bench:
	go test -run=none -bench=. -benchmem ./...

bench_prof:
	go test -run=none -bench=. -benchmem \
		-cpuprofile=cpu.prof -memprofile=mem.prof ./...

bench_pgo:
	go test -run=none -bench=. -benchmem \
		-pgo=cpu.prof -cpuprofile=cpu.prof -memprofile=mem.prof ./...
