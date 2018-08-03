MODULE_FILES            = modules/dumb/dumb.go
#MODULE_FILES           += modules/pwet/pwet.go
MODULE_FILES	       += modules/jira/jira.go
MODULE_FILES	       += modules/whoami/whoami.go
MODULE_FILES_PATTERN    = %.go
MODULE_SO_PATTERN       = %.so
MODULE_SO               = $(patsubst %.go, %.so, $(MODULE_FILES))
SYLVESTRE_BIN           = sylvestre

all: modules build

build: $(SYLVESTRE_BIN)

$(SYLVESTRE_BIN): $(shell find . -name \*.go)
	go build -o $(SYLVESTRE_BIN)

modules: $(MODULE_SO)

$(MODULE_SO_PATTERN) : $(MODULE_FILES_PATTERN)
	go build -buildmode=plugin -o $@ $<

clean: 
	rm -rf $(MODULE_SO) $(SYLVESTRE_BIN)
