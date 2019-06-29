GOCMD = go
GOBUILD = $(GOCMD) build

BINARY_NAME = RUN

build:
	$(GOBUILD) -o $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)
clean:
	rm -f $(BINARY_NAME)


	
