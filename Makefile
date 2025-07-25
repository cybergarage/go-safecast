# Copyright (C) 2022 The go-safecast Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

MODULE_ROOT=github.com/cybergarage/go-safecast

PKG_NAME=safecast

PKG=${MODULE_ROOT}/${PKG_NAME}
PKG_DIR=${PKG_NAME}

TEST_PKG_NAME=test
TEST_PKG=${MODULE_ROOT}/${TEST_PKG_NAME}
TEST_PKG_DIR=${TEST_PKG_NAME}

COVER_PROF=coverage

.PHONY: format vet lint cover clean
.IGNORE: lint

all: test

format:
	gofmt -s -w ${PKG_DIR} ${TEST_PKG_DIR}

vet: format
	go vet ${PKG} ./${TEST_PKG_NAME}/...

lint: vet
	golangci-lint run ${PKG_DIR}/... ${TEST_PKG_DIR}/...

test: lint
	go test -v -cover -coverpkg=${PKG} -coverprofile=${COVER_PROF}.out -timeout 60s ${PKG}/... ./${TEST_PKG_NAME}/...
	go tool cover -html=${COVER_PROF}.out -o ${COVER_PROF}.html

cover: test
	open ${COVER_PROF}.html || xdg-open ${COVER_PROF}.html || gnome-open ${COVER_PROF}.html

fuzz: test
	pushd ${TEST_PKG_DIR} && ./fuzz && popd

clean:
	go clean -i ${PKG}

watchtest:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make test

watchlint:
	fswatch -o . -e ".*" -i "\\.go$$" | xargs -n1 -I{} make lint
