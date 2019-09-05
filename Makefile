IMAGE_NAME=guoxf/govanityurls

build:
	CGO_ENABLED=0 go build -o govanityurls main.go
	sudo docker build -t $(IMAGE_NAME) .
run:
	sudo docker run -p 8088:8088 $(IMAGE_NAME)