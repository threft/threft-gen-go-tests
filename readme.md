The threft-gen-go-tests repository contains tests to find wether the generated code works as expected.

Each folder in this repository contains at least four items:
 - golden/: Designed code that is expected to be generator output
 - *_test.go: A `go test` testing that the golden code works as expected
 - *.thrift: File(s) to be used by threft to generate code
 - generated/: This directory will contain code generated from the thrift file(s). The code in this folder should match the expected golden code.