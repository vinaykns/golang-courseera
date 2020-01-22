PACKAGE=github.com/vinaykns/golang-courseera

GO_BUILD_ARGS=CGO_ENABLED=0 GO111MODULE=on


.PHONY: findian
findian:
	cd ./week2/me/findian && $(GO_BUILD_ARGS) go build -o findian $(PACKAGE)/week2/me/findian
