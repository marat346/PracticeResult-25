build:
	go build -o bin/main main.go

copy:
	cp cities.csv bin/linux_amd64
	cp cities.csv bin/windows_amd64
	cp cities.csv bin/mac_amd64
	cp cities.csv bin/mac_arm64
	cp -r configs bin/linux_amd64
	cp -r configs bin/windows_amd64	
	cp -r configs bin/mac_amd64	
	cp -r configs bin/mac_arm64		

compile:
	GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/cities_server cmd/main.go	
	GOOS=darwin GOARCH=amd64 go build -o bin/mac_amd64/cities_server cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/mac_arm64/cities_server cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows_amd64/cities_server.exe cmd/main.go	

all: compile copy