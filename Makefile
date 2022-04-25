NAME = main
GO_FILES:= main.go ./...
TEXT_FILE = text/sample.txt text/num.txt

build: 
	go build -o $(NAME) $(GO_FILES)

run:
	go run . $(TEXT_FILE)

clean:
	rm -f $(NAME)

.PHONY: build run clean