build:
	docker build -t github.com/rafaeros/fk-api:latest .

run:
	docker run -d -p 8080:8080 --name fk-api github.com/rafaeros/fk-api:latest