.PHONY: test build dist release release-dryrun clean

NAME = covercmp

DIST_OPTS = -a -tags netgo -installsuffix netgo
BUILD_CMD = go build $(DIST_OPTS)
SRC_FILES = main.go
DIST_DIR = ./dist

test:
	go test -cover ./...

build:
	go build -o $(NAME) $(SRC_FILES)

$(NAME)-%:
	GOOS=$* $(BUILD_CMD) -o $(DIST_DIR)/$@ $(SRC_FILES)

dist: $(NAME)-darwin $(NAME)-linux $(NAME)-windows

release:
	semantic-release -noci -ghr -vf -slug arithran/covercmp
	ghr $(cat .ghr) dist

release-dryrun:
	semantic-release -noci -ghr -vf -slug arithran/covercmp -dry

clean:
	rm -f $(NAME)
	rm -rf $(DIST_DIR)

