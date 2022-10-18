build:
	go build -o tt

build-dev:
	go build -o tt -ldflags "-X main.sessionFile=./sessions.json"