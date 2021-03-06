package template

const ConfTemplate = `output_dir: ./.xxx_vendor.proto
external_dir: ./.xxx_externalservice

local_proto_dirs:
  - proto

build:
  - file_path: proto/api/api.proto

    plugins:
      - name: "go"
        path: "bin/protoc-gen-go"
        options:
          - paths=source_relative

      - name: go-grpc
        path: bin/protoc-gen-go-grpc
        options:
          - paths=source_relative

rules:
  - prefix: google/api
    repo: googleapis/googleapis
    host: github.com

  - prefix: google/protobuf
    repo: protocolbuffers/protobuf
    subpath: src
    host: github.com

  - prefix: google/rpc
    repo: googleapis/googleapis
    host: github.com

  - prefix: protoc-gen-openapiv2
    repo: grpc-ecosystem/grpc-gateway
    host: github.com

  - prefix: google/datastore
    repo: googleapis/googleapis
    host: github.com

  - prefix: google/type
    repo: googleapis/googleapis
    host: github.com

  - prefix: google/logging
    repo: googleapis/googleapis
    host: github.com

  - prefix: google/analytics
    repo: googleapis/googleapis
    host: github.com

  - prefix: google/appengine
    repo: googleapis/googleapis
    host: github.com

  - prefix: google/longrunning
    repo: googleapis/googleapis
    host: github.com
external:
  - host: github.com
    repo: googleapis/googleapis
    file_path: google/datastore/v1/datastore.proto
    plugins:
      - name: go
        path: bin/protoc-gen-go
        options:
          - paths=source_relative

      - name: go-grpc
        path: bin/protoc-gen-go-grpc
        options:
          - paths=source_relative

  - host: github.com
    repo: googleapis/googleapis
    file_path: google/logging/v2/logging.proto
    plugins:
      - name: go
        path: bin/protoc-gen-go
        options:
          - paths=source_relative

      - name: go-grpc
        path: bin/protoc-gen-go-grpc
        options:
          - paths=source_relative

  - host: github.com
    repo: googleapis/googleapis
    file_path: google/analytics/admin/v1alpha/analytics_admin.proto
    plugins:
      - name: go
        path: bin/protoc-gen-go
        options:
          - paths=source_relative

      - name: go-grpc
        path: bin/protoc-gen-go-grpc
        options:
          - paths=source_relative

  - host: github.com
    repo: googleapis/googleapis
    file_path: google/appengine/v1/appengine.proto
    plugins:
      - name: go
        path: bin/protoc-gen-go
        options:
          - paths=source_relative

      - name: go-grpc
        path: bin/protoc-gen-go-grpc
        options:
          - paths=source_relative
`
