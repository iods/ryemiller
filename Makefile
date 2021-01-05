
all: clean build

clean:
	@docker-compose down

build: clean
	@docker-compose up -d

.PHONY: all clean build