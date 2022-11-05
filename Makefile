GO = go

ALL = authenticator cleanup
LIBRARIES = $(shell find pkg -type f -iname '*.go')

all: $(ALL)

$(ALL): %: cmd/%/main.go $(LIBRARIES)
	$(GO) build -o $@ $<

clean:
	$(RM) $(ALL)

.PHONY: all clean