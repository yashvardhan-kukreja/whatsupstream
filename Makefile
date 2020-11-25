# Copyright © 2020 Yashvardhan <yash.kukreja.98@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
################################################################################
# ========================== Capture Environment ===============================
# get the repo root and output path
REPO_ROOT:=${CURDIR}
OUT_DIR=$(REPO_ROOT)/bin
################################################################################
# ========================= Setup Go With Gimme ================================
# go version to use for build etc.
# setup correct go version with gimme
PATH:=$(shell . hack/build/setup-go.sh && echo "$${PATH}")
# enable modules
GO111MODULE=on
export PATH GOROOT GO111MODULE
################################################################################
# ============================== OPTIONS =======================================
# install tool
INSTALL?=install
# install will place binaries here, by default attempts to mimic go install
INSTALL_DIR_DEV?=$(shell hack/build/goinstalldir.sh)
INSTALL_DIR?="/usr/local/bin"
# the output binary name, overridden when cross compiling
WHATSUPSTREAM_BINARY_NAME?=whatsupstream
################################################################################
# ================================= Building ===================================
# standard "make" target -> builds
all: build
# builds the whatsupstream and outputs it to $(OUT_DIR)
build:
	go build -v -o $(OUT_DIR)/$(WHATSUPSTREAM_BINARY_NAME)
# for devs contributing to whatsupstream
install-dev: build
	$(INSTALL) -d $(INSTALL_DIR_DEV)
	$(INSTALL) $(OUT_DIR)/$(WHATSUPSTREAM_BINARY_NAME) $(INSTALL_DIR_DEV)/$(WHATSUPSTREAM_BINARY_NAME)
# for normal users
install: build
	$(INSTALL) -d $(INSTALL_DIR)
	$(INSTALL) $(OUT_DIR)/$(WHATSUPSTREAM_BINARY_NAME) $(INSTALL_DIR)/$(WHATSUPSTREAM_BINARY_NAME)
################################################################################
# ================================= Cleanup ====================================
# cleans up the project
clean:
	rm -rf $(OUT_DIR)/
################################################################################
# ============================== Auto-Update ===================================
# update the latest code in rightful way -> gofmt
update: gofmt
# gofmt
gofmt:
	hack/make-rules/update/gofmt.sh
################################################################################
# ================================== Linting ===================================
# linter checks
lint:
	hack/make-rules/verify/lint.sh
#################################################################################
################################################################################
# ================================== Run Whatsupstream ===================================
# runs whatsupstream with the provided config as a background process
notify: all
	hack/run/notify.sh $1
stop:
	hack/run/stop.sh

.PHONY: all build install clean update gofmt lint

