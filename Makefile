build:
	@go build -o ./bin/registrar-student-insertor

run: build 
	@./bin/csv-to-mysql 
	 
win:
	@env GOOS=windows GOARCH=amd64 go build -o ./bin/registrar-student-insertor.exe

