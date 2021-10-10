remote_file(
    name = "godeps",
    binary = True,
    exported_files = ["godeps"],
    extract = True,
    url = f"https://github.com/sagikazarmark/please-go-modules/releases/latest/download/godeps_{CONFIG.HOSTOS}_{CONFIG.HOSTARCH}.tar.gz",
)
go_toolchain(
    name = "go_toolchain",
    version = "1.14.13",
)

WOLLEMI_VERSION = "v0.7.0"
remote_file(
    name = "wollemi",
    url = f"https://github.com/tcncloud/wollemi/releases/download/{WOLLEMI_VERSION}/wollemi-{WOLLEMI_VERSION}-{CONFIG.HOSTOS}-{CONFIG.HOSTARCH}.tar.gz",
    extract = True,
    exported_files = ["wollemi"],
    binary = True,
)

sh_cmd(
    name = "plz-tidy",
    cmd = [
        "export GOROOT=\\\\$($(out_exe :go_toolchain|go) env GOROOT)",
        "$(out_exe :godeps) -dir third_party/go -clean -builtin -wollemi",
        "$(out_exe :wollemi) gofmt ./...",
    ],
    deps = [
        ":godeps",
        ":wollemi",
        ":go_toolchain",
    ],
)