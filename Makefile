default:
	@go build -o ./bin/registrar-student-insertor

build:
	@go build -o ./bin/registrar-student-insertor

run: build 
	@./bin/registrar-student-insertor
	 
win:
	@env GOOS=windows GOARCH=amd64 go build -o ./bin/registrar-student-insertor.exe

clean:
	@rm bin/*
	@rm registrar-student-insertor*
	@rm *.txt
