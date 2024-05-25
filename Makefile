# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Main application name
APP_NAME=Troupie-Tracker

# Directories
SRC_DIR=./
CLI_APP_DIR=$(SRC_DIR)/app/cli
WEB_APP_DIR=$(SRC_DIR)/app/web

# Binary names
CLI_BINARY=$(CLI_APP_DIR)/cli
WEB_BINARY=$(WEB_APP_DIR)/web

# Targets
.PHONY: all cli web clean run

all: cli web

cli:
	$(GOBUILD) -o $(CLI_BINARY) $(CLI_APP_DIR)/...

web:
	$(GOBUILD) -o $(WEB_BINARY) $(WEB_APP_DIR)/...

run-cli: cli
	$(CLI_BINARY)

run-web: web
	$(WEB_BINARY)

clean:
	@echo cleaning...
	@$(GOCLEAN)
	@rm -f $(CLI_BINARY) $(WEB_BINARY)

