new:
	mkdir -p bin
	cd bin && go build ../votebot && cd ..
	cp docs/exampleConfig.yml bin/config.yml
run:
	cd bin && go build ../votebot
	cd bin && ./votebot