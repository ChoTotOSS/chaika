build:
	CGO_ENABLED=0 GOOS=linux go build -o dist/chaika -a -tags netgo -ldflags '-w' .
