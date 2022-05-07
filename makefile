build:
	go generate
	go build -o main main.go

wire:
	cd app && \
	rm wire_gen.go && \
	go run github.com/google/wire/cmd/wire && \
	cd ..

clean:
	rm -f main