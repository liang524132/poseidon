load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "k8sclient.go",
        "keyed_queue.go",
        "nodewatcher.go",
        "podwatcher.go",
        "types.go",
        "utils.go",
    ],
    importpath = "github.com/kubernetes-sigs/poseidon/pkg/k8sclient",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/firmament:go_default_library",
        "//pkg/metrics:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/google/uuid:go_default_library",
        "//vendor/github.com/jinzhu/copier:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/fields:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/labels:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/tools/cache:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "keyed_queue_test.go",
        "nodewatcher_test.go",
        "podwatcher_test.go",
    ],
    embed = [":go_default_library"],
    deps = [
        "//pkg/firmament:go_default_library",
        "//vendor/github.com/golang/mock/gomock:go_default_library",
        "//vendor/k8s.io/api/core/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/api/resource:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/types:go_default_library",
        "//vendor/k8s.io/client-go/kubernetes/fake:go_default_library",
        "//vendor/k8s.io/client-go/tools/cache:go_default_library",
    ],
)
