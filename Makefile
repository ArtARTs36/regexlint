# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

docker-build:
	docker build -f Dockerfile . -t artarts36/regexlint:testing

# Usage as: make docker-run ARGS="go file.yaml headers.cors"
docker-run:
	docker run \
		-v "./:/app" \
		-w /app \
		--rm artarts36/regexlint:testing \
		$(ARGS)

test:
	go test ./...
