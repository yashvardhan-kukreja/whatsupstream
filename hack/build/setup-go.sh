#!/bin/bash
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

# read go-version file unless GO_VERSION is set
GO_VERSION="${GO_VERSION:-"$(cat .go-version)"}"

# setting up GIMME_ENV_PREFIX because .env might just be present already in the repo
# and if it does already exist, then gimme won't trigger re-generating a new .env file
# hence, prefixing the GIMME-generated env file with ./bin/.gimme so as to save the .env in ./bin/.gimme/
export GIMME_ENV_PREFIX=./bin/.gimme/
export GIMME_SILENT_ENV=y

# setup go if `go version` doesn't match
# go version output looks like:
# go version go1.14.5 darwin/amd64
if ! ([ -n "${FORCE_HOST_GO:-}" ] || \
      (command -v go >/dev/null && [ "$(go version | cut -d' ' -f3)" = "go${GO_VERSION}" ])); then
    # eval because the output of this is shell to set PATH etc.
    eval "$(hack/third_party/gimme/gimme "${GO_VERSION}")"
fi

# force go modules
export GO111MODULE=on
