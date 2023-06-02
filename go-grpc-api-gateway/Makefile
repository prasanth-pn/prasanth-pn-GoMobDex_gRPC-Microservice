proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	export GIN_MODE=release && go run cmd/*.go
	