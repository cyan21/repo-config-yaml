build:
	rm -f gen*.yml
	go build -o repo-mgmt

run:
	./repo-mgmt export local

all: build run