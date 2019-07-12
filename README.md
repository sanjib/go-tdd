# go-tdd
TDD exercises based on the excellent learn-go-with-tests by quii: https://quii.gitbook.io/learn-go-with-tests/

The 3 laws of TDD

1. Write test first
2. Fail, then make it work
3. Refactor

## 1. hello

Key points

- test file format: xxx_test.go where xxx = go file name
- test function format: TestXxx() where Xxx = function name
- test variable naming convention: got, want
- t *testing.T is the hook into the testing framework
- constants can improve readability and performance 
- use t.Helper() to mark helper test function
- named returns doesn't look so bad [[...]](#quote-named-return)

## 2. integers

Key points

- example functions can be added in xxx_test.go
- ExampleXxx function will be shown in documentation
- output syntax: // Output: xxx 

## 3. iteration

Key points

- benchmark functions can be placed in xxx_test.go
- BenchmarkXxx where Xxx = function to benchmark

## Quotes 

The following quotes are from https://quii.gitbook.io/learn-go-with-tests/

>By not writing tests you are committing to ***manually checking your code by running your software which breaks your state of flow*** and you won't be saving yourself any time, especially in the long run.

>Write enough code to satisfy the compiler and _that's all_ - remember ***we want to check that our tests fail for the correct reason***.

>Keep the discipline! You don't need to know anything new right now to make the test fail properly.

The following quotes are from the blue book http://www.gopl.io/

<a name="quote-named-return"></a>
>[…] with many return statements and several results, bare returns can reduce code duplication, but they rarely make code easier to understand. […] Bare returns are best used sparingly

