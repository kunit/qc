PKG = github.com/kunit/qc
COMMIT = $$(git describe --tags --always)

GO ?= GO111MODULE=on go

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT)
RELEASE_BUILD_LDFLAGS = -s -w $(BUILD_LDFLAGS)

default: build
ci: depsdev test

depsdev:
	GO111MODULE=off go get golang.org/x/tools/cmd/cover
	GO111MODULE=off go get golang.org/x/lint/golint
	GO111MODULE=off go get github.com/motemen/gobump/cmd/gobump
	GO111MODULE=off go get github.com/Songmu/goxz/cmd/goxz
	GO111MODULE=off go get github.com/tcnksm/ghr
	GO111MODULE=off go get github.com/Songmu/ghch

lint:
	golint -set_exit_status $(TEST)
	$(GO) vet $(shell go list ./... | grep -v misc)
	$(GO) fmt $(shell go list ./... | grep -v misc)

test:
	$(GO) test -v $(shell go list ./... | grep -v misc) -coverprofile=coverage.txt -covermode=count

build:
	$(GO) build -ldflags="$(BUILD_LDFLAGS)" ./cmd/qc

crossbuild: depsdev
	$(eval ver = v$(shell gobump show -r version/))
	GO111MODULE=on goxz -pv=$(ver) -os=linux,darwin -arch=386,amd64 -build-ldflags="$(RELEASE_BUILD_LDFLAGS)" \
	  -d=./dist/$(ver) ./cmd/qc

prerelease:
	$(eval ver = v$(shell gobump show -r version/))
	ghch -w -N ${ver}

release:
	$(eval ver = v$(shell gobump show -r version/))
	ghr -username kunit -replace ${ver} dist/${ver}

.PHONY: default test
