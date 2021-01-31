start:
	sam validate
	sam build
	sam local start-api

deploy:
	sam build && sam deploy --no-confirm-changeset