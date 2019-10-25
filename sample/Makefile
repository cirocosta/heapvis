build:
	go build -v -i -o sample.out

run: build
	./sample.out

profile: run
	go tool pprof ./profile.pprof

