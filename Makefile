BINARY_NAME= SpotifyTracker

build:
	cd src/ && go build -o ../${BINARY_NAME}

run: build
	./${BINARY_NAME}

clean:
	go clean

fclean: clean
	rm ${BINARY_NAME}

re: fclean run