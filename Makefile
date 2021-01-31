start:
	sam validate
	sam build
	sam local start-api

deploy:
	sam build && sam deploy --no-confirm-changeset

formatter:
	gofmt -d . 1>&2 || exit 1

linter: # https://golangci-lint.run/usage/quick-start/
	golangci-lint run ./gif-generator/