# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Build the Database Daemon binary.
FROM docker.io/golang:1.15 as builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY oracle/pkg/ oracle/pkg/
COPY oracle/cmd/ oracle/cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o logging_main oracle/cmd/logging/logging_main.go

# Build the container image
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /build/logging_main /

USER nonroot:nonroot
