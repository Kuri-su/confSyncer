GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build

gen:
	${GO} run cmd/gen/main.go

cs:
	${GOBUILD} -o _output/confsyncer ./cmd/confsyncer