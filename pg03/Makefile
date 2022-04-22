NAME = main
GO_FILES:= main.go sumfile/sumfile.go
TEXT_FILE = sample.txt

build: 
	go build -o $(NAME) $(GO_FILES)

run:
	go run . $(TEXT_FILE)

clean:
	rm -f $(NAME)

.PHONY: build run clean