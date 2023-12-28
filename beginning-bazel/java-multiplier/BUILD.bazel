java_library(
    name = "IntMultiplierLib",
    srcs = ["src/main/java/com/raejz/multi/IntMultiplier.java"],
    visibility = ["//visibility:public"],
)

java_test(
    name = "IntMultilierLibTest",
    srcs = ["src/test/java/com/raejz/multi/IntMultiplierTest.java"],
    test_class = "com.raejz.multi.IntMultiplierTest",
    deps = [
        ":IntMultiplierLib",
    ],
)
