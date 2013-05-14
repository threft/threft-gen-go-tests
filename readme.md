The threft-gen-go-tests repository contains tests to find wether the generated code works as expected.

Each folder in this repository tests a certain part of threft.
Each folder contains a 'golden' test, comparing generated code to expected (designed) code, and an actual `go test`, which tests wether the code works as expected (actually testing the golden/designed code).