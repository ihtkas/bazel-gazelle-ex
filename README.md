###### Repository to test bazel and gazelle tool to build go project with proto and gRPC dependencies.

**Usecase:**

Use bazel and gazelle to build go repository with pre generated proto files. There are issues in using [validate rules](https://github.com/envoyproxy/protoc-gen-validate) from envoyproxy in proto definition and the bazel tool fails to build repository with this dependency. To reproduce the issue, a proto file with a sample gRPC service is defined in the same repository in protos/service.proto. 

Repository is cloned into a path ~/go/src/github.com/ihtkas/bazel-gazelle-ex


**Commands used to build the repository:**

- Generate proto files before building repository.


`protoc -I=. -I ~/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 --go_out=./../../../  --validate_out="lang=go:./../../../" --go-grpc_out=./../../../ ./protos/pinger/service.proto`

- Use gazelle to generate dependencies

`bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies -build_file_proto_mode=disable_global -prune`

- Use gazelle to generate BUILD.bazel files

`bazel run //:gazelle`

- Build pinger go binary

`bazel build //cmd/pinger`

The last build command throws an error

`go/src/github.com/ihtkas/bazel-gazelle-ex/api/pinger/BUILD.bazel:3:11: error loading package '@com_github_envoyproxy_protoc_gen_validate//validate': Unable to find package for @com_google_protobuf//:protobuf.bzl: The repository '@com_google_protobuf' could not be resolved. and referenced by '//api/pinger:pinger'
ERROR: Analysis of target '//cmd/pinger:pinger' failed; build aborted: Analysis failed`

There is a [github issue](https://github.com/bazelbuild/bazel-gazelle/issues/988) for the same.  
