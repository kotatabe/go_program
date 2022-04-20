NAME = main
GO_FILES:=$(shell find . -type f -name '*.go' -print)
TEXT_FILE = sample.txt

build: 
	go build -o $(NAME) $(GO_FILES)

run: build
	./$(NAME) $(TEXT_FILE)

clean:
	rm -f $(NAME)

.PHONY: build run clean