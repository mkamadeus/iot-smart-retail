build:
	go generate
	go build -o main main.go

wire:
	cd app && \
	go run github.com/google/wire/cmd/wire && \
	cd ..

clean:
	rm -f main
.PHONY: build
