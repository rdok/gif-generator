start:
	rm -rf .aws-sam
	rm -f hello-world/main
	cd hello-world && go build -o hello-world main.go
	sam local start-api