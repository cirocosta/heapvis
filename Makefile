install:
	go install -v ./cmd/heapvis

test:
	go test ./...


profile: sample.out
	./$<
	go tool pprof ./profile.pprof

sample.out:
	go build -v -i -o $@ ./sample


