all:
	mkdir -p bin
	cd bin && go build ../votebot && cd ..