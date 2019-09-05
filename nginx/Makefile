IMAGE_NAME=guoxf/govanityurls-nginx

run-nginx:
	 sudo docker run -p 80:80 -p 443:443 $(IMAGE_NAME)

gen-ca:
	openssl genrsa -out ./cert/rootCA.key 2048
	openssl req -x509 -new -nodes -key ./cert/rootCA.key -config openssl.cnf -days 5000 -out ./cert/rootCA.pem -subj "/C=CN/ST=Shanghai/L=Shanghai/O=devops/OU=devops/CN=gxfxx.com"
	openssl genrsa -out ./cert/cert.key 2048
	openssl req -new -key ./cert/cert.key -config openssl.cnf -out ./cert/cert.csr -subj "/C=CN/ST=Shanghai/L=Shanghai/O=devops/OU=devops/CN=gxfxx.com"
	openssl x509 -req -in ./cert/cert.csr -CA ./cert/rootCA.pem -CAkey ./cert/rootCA.key -CAcreateserial -out ./cert/cert.crt -days 5000

add-root:
	cp /etc/ssl/certs/ca-certificates.crt ./
	cat ./cert/rootCA.pem >> ./ca-certificates.crt
	sudo mv ./ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

build-image:
	sudo docker build -t $(IMAGE_NAME) .