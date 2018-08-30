compile:
	GOARCH=wasm GOOS=js go build -o synth.wasm main.go

start:
	go run server.go