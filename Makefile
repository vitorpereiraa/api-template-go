run: clean build docs
	./api-template-go.exe

build:
	go build -o ./api-template-go.exe

clean:
	rm -f ./api-template-go.exe 

docs: docs-clean
	swag init ./

docs-clean:
	rm -rf ./docs/

dr: docker-clean docker-build  
	docker run -d -p 8080:8080 --name api-template-go api-template-go:dev

docker-build: 
	docker build . -t api-template-go:dev

docker-clean: docker-stop
	 docker container rm api-template-go || true

docker-stop:
	docker container stop api-template-go || true


