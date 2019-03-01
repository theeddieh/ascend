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

# Project-specific variables
BINARY_NAME  = ascendb
INPUT_FILE   = input-1.log
SANITY_FILE  = sanity.log


# Targets
run:
	${GO_RUN} ${REPO_PATH} ${INPUT_FILE}

build:
	${GO_BUILD} -o ${BINARY_NAME}

test:
	${GO_TEST} ./... -v

check:
	${GO_VERSION}
	${GO_ENV}

sanity: clean check build
	./${BINARY_NAME} ${SANITY_FILE} -v

clean:
	rm -f ${BINARY_NAME}