init: docker-build docker-up
docker-build:
	docker build -t fibonacci-go .

docker-up:
	docker run -p 8000:8000 fibonacci-go