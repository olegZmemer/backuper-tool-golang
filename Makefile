build:
	go build -o ./backuper

init: build
	sudo ./backuper init

add: build
	sudo ./backuper add