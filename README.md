# go-tdd
TDD exercises based on the excellent learn-go-with-tests by quii: https://quii.gitbook.io/learn-go-with-tests/

The 3 laws of TDD

1. Write test first
2. Fail, then make it work
3. Refactor

The following are notes from the different sections of learn-go-with-tests.

If the content does not correspond with the section topic, it's
only because they were introduced in that section.

## 1. hello

- test file format: xxx_test.go where xxx = go file name
- test function format: TestXxx() where Xxx = function name
- test variable naming convention: got, want
- t *testing.T is the hook into the testing framework
- constants can improve readability and performance 
- use t.Helper() to mark helper test function
- named returns doesn't look so bad [[...]](#quote-named-return)

## 2. integers

- example functions can be added in xxx_test.go
- ExampleXxx function will be shown in documentation
- output syntax: // Output: xxx 

## 3. iteration

- benchmark functions can be placed in xxx_test.go
- BenchmarkXxx where Xxx = function to benchmark
- benchmark test: ```go test -bench=.```

## 4. arrays and slices

- coverage: ```go test -cover```

## Quotes 

The following quotes are from https://quii.gitbook.io/learn-go-with-tests/

>By not writing tests you are committing to ***manually checking your code by running your software which breaks your state of flow*** and you won't be saving yourself any time, especially in the long run.

>Write enough code to satisfy the compiler and _that's all_ - remember ***we want to check that our tests fail for the correct reason***.

>Keep the discipline! You don't need to know anything new right now to make the test fail properly.

>It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. [...] Every test has a cost.

The following quotes are from the blue book http://www.gopl.io/

<a name="quote-named-return"></a>
>[…] with many return statements and several results, bare returns can reduce code duplication, but they rarely make code easier to understand. […] Bare returns are best used sparingly

