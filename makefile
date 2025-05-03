APP_NAME := task-manager
CMD_DIR := cmd/${APP_NAME}

# default target
all: build
run:
	go run ${CMD_DIR}/main.go