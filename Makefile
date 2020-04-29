GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build


gen:
	${GOBUILD} -o bin/gen ./cmd/gen

confSyncer:
	${GOBUILD} -o bin/confSyncer ./cmd/confSyncer