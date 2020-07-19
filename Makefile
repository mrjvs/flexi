all: gateway
	@echo Fuck you

gateway:
	go build -o ./build/gateway ./cmd/gateway

run-gateway: gateway
	@./build/gateway

msg_handler:
	go build -o ./build/msg_handler ./cmd/msg_handler

run-msg_handler: msg_handler
	@./build/msg_handler

ribbon:
	go build -o ./build/ribbon ./cmd/ribbon

run-ribbon: ribbon
	@./build/ribbon