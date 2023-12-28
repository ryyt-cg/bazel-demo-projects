# Beginning Bazel - Build and Test Java, Go and more

<!-- TOC -->

* [Beginning Bazel - Build and Test Java, Go and more](#beginning-bazel---build-and-test-java-go-and-more)
    * [Create First Bazel Project](#create-first-bazel-project)
        * [Configure Bazel Project](#configure-bazel-project)
        * [Adding Package & Source Code](#adding-package--source-code)
        * [Adding Source Code](#adding-source-code)
    * [WORKSPACE FIle Functionality](#workspace-file-functionality)

<!-- TOC -->

## Create First Bazel Project

### Configure Bazel Project

|   | Description                                                                                        |
|---|----------------------------------------------------------------------------------------------------|
| 1 | create a project directory, let's call it beginning-bazel <br/>mkdir beginning-bazel               |
| 2 | go into beginning-bazel and add an empty WORKSPACE file<br/>cd beginning-bazel<br/>touch WORKSPACE |
| 3 | optionally, add BUILD.bazel<br/>touch BUILD.bazel                                                  |

This is a minimum Bazel project and the WORKSPACE file location should always be at the root of your
Bazel project.

### Adding Package & Source Code

* Create a directory for Java src. Bazel does not require configuring maven-like project structure.
  By Java convention, let's create maven-like project.

```shell
mkdir -pv java-hello-world/src/main/java/com/raejz/greet
```

add HelloWorld.java and BUILD file<br/>
add this to BUILD file to build java binary

```
java_binary(
    name = "JavaHelloWorld",
    srcs = ["src/main/java/com/raejz/greet/HelloWorld.java"],
    main_class = "com.raejz.greet.HelloWorld",
    deps = [],
)
```

build the binary

```shell
bazel build java-hello-world:JavaHelloWorld
```

run the binary

```shell
bazel run java-hello-world:JavaHelloWorld
```

* Create a directory for maven package

```shell
mvn archetype:generate -DgroupId=com.raejz.hello -DartifactId=maven-hello-world -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false
```

add BUILD file<br/>
insert this into BUILD file to build java binary

```
java_binary(
    name = "MavenHelloWorld",
    srcs = ["src/main/java/com/raejz/hello/App.java"],
    main_class = "com.raejz.hello.App",
    deps = [],
)
```

build the binary

```shell
bazel build maven-hello-world:MavenHelloWorld
```

run the binary

```shell
bazel run maven-hello-world:MavenHelloWorld
```

* Creating and Using Dependencies

create java-multiplier directory<br/>
add IntMultiplier.java
add BUILD and insert this

```
java_library(
    name = "IntMultiplierLib",
    srcs = ["src/main/java/com/raejz/multi/IntMultiplier.java"],
    visibility = ["//visibility:public"],
)

java_test(
    name = "MavenHelloWorldTest",
    srcs = ["src/test/java/com/raejz/hello/AppTest.java"],
    test_class = "com.raejz.hello.AppTest",
    deps = [],
)
```

modify BUILD in java-hello-world and add a dependency

```
deps = ["//java-multiplier:IntMultiplierLib"],
```

## WORKSPACE FIle Functionality
add Golang configuration to WORKSPACE

```
## Golang Configuration & Setup
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6ab6b57e48c09523e93050f13698f708428cfd5e619252e369d377af6597707",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.43.0/rules_go-v0.43.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.43.0/rules_go-v0.43.0.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()
go_register_toolchains(version = "1.21.3")
```

## Go Echo Server
add echo_server.go<br/>
add BUILD and insert this
```
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_test", "go_library")

go_binary(
    name = "echo_server",
    srcs = ["echo_server.go"],
)
```
load go_binary, go_test, go_library, etc. to the package


## Demo API with Gazelle
add bazel and gazelle configuration to WORKSPACE file
```
## Golang Configuration & Setup
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6ab6b57e48c09523e93050f13698f708428cfd5e619252e369d377af6597707",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.43.0/rules_go-v0.43.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.43.0/rules_go-v0.43.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

go_rules_dependencies()
go_register_toolchains(version = "1.21.3")
gazelle_dependencies()
```
version=1.21.3 is Golang SDK version>br/>
add BUILD or BUILD.bazel file in root project

**Important**: For Go projects, replace the string after prefix with the portion of your import path that corresponds to your repository.
```
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix gitlab.com/aionx/bazel-demo-projects/beginning-bazel
gazelle(name = "gazelle")
```
After adding this code, you can run Gazelle with Bazel.
```shell
bazel run //:gazelle
```
This will generate new BUILD.bazel files for your project. You can run the same command in the future to update existing BUILD.bazel files to include new source files or options.

You can write other gazelle rules to run alternate commands like update-repos.
```
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```
You can also pass additional arguments to Gazelle after a -- argument.
```shell
bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```
this command will generate BUILD file in golang projects/packages<br/>
After running update-repos, you might want to run bazel run //:gazelle again, as the update-repos command can affect the output of a normal run of Gazelle.<br/>
For Bazel plugin in Intellij, click sync projects with BUILD file.


At the end, BUILD file may look like this
```
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix gitlab.com/aionx/bazel-demo-projects/beginning-bazel
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
```
run bazel commands
```shell
bazel run //:gazelle

```