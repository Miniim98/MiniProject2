chitty-server-build:
	@docker build -t chitty-server -f DockerfileServer .

chitty-server-run:
	@docker run -i -t -p 8008:80 chitty-server

chitty-client-build:
	@docker build -t chitty-client -f DockerfileClient .

chitty-client-run:
	@docker run -i -t  chitty-client

server-rebuild:
	@docker build --no-cache -t chitty-server -f DockerfileServer .

client-rebuild:
	@docker build --no-cache -t chitty-client -f DockerfileClient .
