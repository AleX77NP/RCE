start-server:
	cd cmd/rce && go run main.go

clean:
	docker rm $$(docker ps -a -f status=exited -q)
