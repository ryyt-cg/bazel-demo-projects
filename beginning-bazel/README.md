
# Beginning Bazel - Build and Test Java, Go and more

<!-- TOC -->
* [Beginning Bazel - Build and Test Java, Go and more](#beginning-bazel---build-and-test-java-go-and-more)
  * [Create First Bazel Project](#create-first-bazel-project)
    * [Configure Bazel Project](#configure-bazel-project)
    * [Adding Source Code](#adding-source-code)
<!-- TOC -->





## Create First Bazel Project

### Configure Bazel Project
|   | Description                                                                                        |
|---|----------------------------------------------------------------------------------------------------|
| 1 | create a project directory, let's call it beginning-bazel <br/>mkdir beginning-bazel               |
| 2 | go into beginning-bazel and add an empty WORKSPACE file<br/>cd beginning-bazel<br/>touch WORKSPACE |
| 3 | optionally, add BUILD.bazel<br/>touch BUILD.bazel                                                  |

This is a minimum Bazel project and the WORKSPACE file location should always be at the root of your Bazel project.

### Adding Source Code
Create a directory for Java src.  Bazel does not require configuring maven-like project structure.  However, Java developers tend to create maven-like project.

```shell
mkdir -pv java-hello-world/src/main/java/com/raejz/greet
```
add HelloWorld.java and BUILD
add this to BUILD file
```
java_binary(
    name = "JavaHelloWorld",
    srcs = ["src/main/java/com/raejz/greet/HelloWorld.java"],
    main_class = "com.raejz.greet.HelloWorld",
    deps = [],
)
```


Create a directory for maven package
```shell
mvn archetype:generate -DgroupId=com.raejz.hello -DartifactId=maven-hello-world -DarchetypeArtifactId=maven-archetype-quickstart -DarchetypeVersion=1.4 -DinteractiveMode=false
```
add BUILD file
insert this into BUILD file
```
java_binary(
    name = "MavenHelloWorld",
    srcs = ["src/main/java/com/raejz/hello/App.java"],
    main_class = "com.raejz.hello.App",
    deps = [],
)
```