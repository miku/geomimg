TARGETS = geomimg

.PHONY: all
all: $(TARGETS)

%: cmd/%/main.go
	go build -o $@ $^

.PHONY: clean
clean:
	rm -f geomimg
