#!/bin/bash
# Copyright 2023 XigXog
#
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at https://mozilla.org/MPL/2.0/.
#
# SPDX-License-Identifier: MPL-2.0

source "$(dirname "${BASH_SOURCE[0]}")/setup.sh"

${SCRIPTS}/clean.sh
${SCRIPTS}/hello-world.sh
${SCRIPTS}/docs.sh
${SCRIPTS}/addlicense.sh

go mod tidy
gofmt -l -s -w .
go vet ./...

git add .
git commit
