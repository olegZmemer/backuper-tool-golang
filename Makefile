build:
	cd cmd && go build -o ../backuper

init: build
	sudo ./backuper init

add: build
	sudo ./backuper add

list: build
	sudo ./backuper list

remove: build
	sudo ./backuper remove 

exec: build
	sudo ./backuper exec  

bin:
	sudo cp ./backuper /usr/local/bin