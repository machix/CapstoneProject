geo-api-linux:
	GOOS=linux go build -o geo-api-linux main.go

geo-service1.docker: geo-api-linux
	docker build -t geo-service1 .
	docker image save --output=geo-service1.docker geo-service1

rsync: geo-service1.docker
	rsync -a geo-service1.docker root@159.203.178.86:/

restart:
	ssh root@159.203.178.86 "docker image load --input=geo-service1.docker; \
		docker kill geo-service1;\
		docker run --env-file ./env.list -p 8000:8000 -d geo-service1"

clean:
	rm *.docker geo-api-linux
