run: clean build
	./api-template-go.exe

build:
	go build -o ./api-template-go.exe

clean:
	rm -f ./api-template-go.exe 
