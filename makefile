chitty-server-build:
	@docker build -t chitty-server -f DockerfileServer .

chitty-server-run:
	@docker run chitty-server

chitty-client-build:
	@docker build -t chitty-client -f DockerfileClient .

chitty-client-run:
	@docker run chitty-client

