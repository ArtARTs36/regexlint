docker-build:
	docker build -f Dockerfile . -t artarts36/regexlint:testing

# Usage as: make docker-run ARGS="go file.yaml headers.cors"
docker-run: docker-build
	docker run \
		-v "./:/app" \
		-w /app \
		--rm artarts36/regexlint:testing \
		$(ARGS)

test:
	go test ./...
