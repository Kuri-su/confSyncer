GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build

build:  cs

cs:
	bash -c "mkdir _output ; echo "
	${GOBUILD} -o _output/confsyncer ./cmd/confsyncer

docker:
	docker rm -f confsycner; echo
	docker build -t confsyncer .