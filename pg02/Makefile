NAME = main
GO_FILES:=$(shell find . -type f -name '*.go' -print)
TEXT_FILE = src.txt
COPY_FILE = copy.txt

build: 
	go build -o $(NAME) $(GO_FILES)

run: build
	./$(NAME) $(TEXT_FILE)

clean:
	rm -f $(NAME) $(COPY_FILE)

.PHONY: build run clean