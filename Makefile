run: clean build
	./api-template-go 

build:
	go build .

clean:
	rm -f ./api-template-go ./api-template-go.exe
