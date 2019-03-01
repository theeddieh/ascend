# Manually set GOPATH to the root of this project
REPO_PATH = github.com/theeddieh/ascend
PROJ_PATH = src/${REPO_PATH}
export GOPATH = $(shell pwd | sed 's|/${PROJ_PATH}||')

# Go commands
GO_TOOL    = go
GO_ENV     = ${GO_TOOL} env
GO_VERSION = ${GO_TOOL} version
GO_TEST    = ${GO_TOOL} test
GO_BUILD   = ${GO_TOOL} build
GO_RUN     = ${GO_TOOL} run
GO_INSTALL = ${GO_TOOL} install

# Target platforms
TARGET_MAC = GOOS=darwin  GOARCH=amd64
TARGET_NIX = GOOS=linux   GOARCH=amd64
TARGET_WIN = GOOS=windows GOARCH=amd64

# Project-specific variables
BINARY_NAME = ascendb
INPUT_FILE  = input-1.log
SANITY_FILE = sanity.log
VERSION     = 0.0.1

# Targets
build:
	${GO_BUILD} -o ${BINARY_NAME}

all:
	${TARGET_MAC} ${GO_BUILD} -o ${BINARY_NAME}_mac
	${TARGET_NIX} ${GO_BUILD} -o ${BINARY_NAME}_nix
	${TARGET_WIN} ${GO_BUILD} -o ${BINARY_NAME}_win

test:
	${GO_TEST} ./... -v

run:
	${GO_RUN} ${REPO_PATH} ${INPUT_FILE}

check:
	${GO_VERSION}
	${GO_ENV}

sanity: clean check build
	./${BINARY_NAME} ${SANITY_FILE} -v

clean:
	rm -f ${BINARY_NAME}