out/swagger.json:
	go run cmd/gen/main.go > out/swagger.json

.PHONY: clean
clean:
	rm -rf out/*