# go-tdd
TDD exercises based on the excellent learn-go-with-tests by quii: https://quii.gitbook.io/learn-go-with-tests/
The following are my personal notes derived from following the book. I have tried to be as
brief as possible for my own benefit and include only what was interesting to
me or thought that I should remember.

You could use it as a review as you finish each section, or to reinforce your 
learning later. The exercises are almost verbatim but not always. It might be 
interesting to you that all exercises are in one repo as opposed to multiple
repos for each section as suggested in the course.

The 3 laws of TDD

1. Write test first
2. Fail, then make it work
3. Refactor

The following are notes from the different sections of learn-go-with-tests.

If the content does not correspond with the section topic, it's
only because they were introduced in that section.

Quotes are relevant to each section under which they appear. Where not specified, quotes are from learn-go-with-tests.

## 1. hello

Key points

- test file format: xxx_test.go where xxx = go file name
- test function format: TestXxx() where Xxx = function name
- test variable naming convention: got, want
- t *testing.T is the hook into the testing framework
- constants can improve readability and performance 
- use t.Helper() to mark helper test function
- named returns doesn't look so bad

Quotes:

>By not writing tests you are committing to ***manually checking your code by running your software which breaks your state of flow*** and you won't be saving yourself any time, especially in the long run.

Quotes from http://www.gopl.io/:

>[…] with many return statements and several results, bare returns can reduce code duplication, but they rarely make code easier to understand. […] Bare returns are best used sparingly

## 2. integers

Key points

- example functions can be added in xxx_test.go
- ExampleXxx function will be shown in documentation
- output syntax: // Output: xxx 

Quotes:

>Write enough code to satisfy the compiler and _that's all_ - remember ***we want to check that our tests fail for the correct reason***.

## 3. iteration

- benchmark functions can be placed in xxx_test.go
- BenchmarkXxx where Xxx = function to benchmark
- benchmark test: ```go test -bench=.```

Quotes:

>Keep the discipline! You don't need to know anything new right now to make the test fail properly.

## 4. arrays and slices

- coverage: ```go test -cover```
- create a slice from array:
```
x := [3]string{"Лайка", "Белка", "Стрелка"} // type of x [3]string
s := x[:] // a slice referencing the storage of x, type of s []string
```
- slice vs slice & copy: https://play.golang.org/p/bTrRmYfNYCp
- after slicing, if the new slice length is less than the original, 
then it's better to copy it so that the original can be garbage
 collected: https://play.golang.org/p/Poth8JS28sc 

Quotes:

> It should not be a goal to have as many tests as possible, but rather to have as much confidence as possible in your code base. [...] Every test has a cost.

Quotes from https://blog.golang.org/go-slices-usage-and-internals:

> An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types).

> One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.

In the context of slices:
> make allocates an array and returns a slice that refers to that array

> most programs don't need complete control, so Go provides a built-in append function that's good for most purposes

> re-slicing a slice doesn't make a copy of the underlying array

## 5. structs, methods and interfaces

- consider table-driven tests where the data can be used to drive the tests
- interfaces with possible tests of different data are good candidates for 
table-driven tests 

Quotes:

> In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

Quotes from Test-Driven Development by Example, by Kent Beck:

> The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations

## 6. pointers, errors

- it's important to fail correctly with the right error message
- write minimal amount of code for test to run, then check failing test output
- micro incremental test iteration needs to be followed to create good reliable tests
- methods are declared on types: whether structs or types created from existing types
- errors.New creates a new error with a message of your choosing

Quotes:

> Remember to only do enough to make the tests run. We need to make sure our test fails correctly with a clear error message.

## 7. maps

- maps are reference types
- never initialize empty map which is a nil
- trying to write to nil will cause panic: assignment to entry in nil map
- if key doesn't exist, we get the type's "zero" value https://play.golang.org/p/Btz1HnnnDWg

Quotes:

> Maps being a reference is really good, because no matter how big a map gets there will only be one copy.

> attempts to write to a nil map will cause a runtime panic

## 8. dependency injection

- implementation uses an interface so that different calling 
functions depending on use can pass different objects that confirm
to the interface 
- in di_test.go a buffer is passed as io.Writer to Greet
- uncomment the code in main.go to see the same Greet function in 
action in two other scenarios: output to console, output to browser
- 3 examples of io.Writer:
    - bytes.Buffer
    - os.Stdout
    - http.ResponseWriter

Quotes:

> The more familiar you are with the standard library the more you'll see these general purpose interfaces which you can then re-use in your own code to make your software reusable in a number of contexts

## 9. mocking

- implementation has 
```
type Sleeper interface {
	Sleep()
}
```
- test has
```
type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}
```
- main has
```
type DefaultSleeper struct{}
func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}
```

Quotes:

> let's now wire up our function into a main so we have some working software to reassure ourselves we're making progress

> Take a thin slice of functionality and make it work end-to-end, backed by tests

> By investing in getting the overall plumbing working right, we can iterate on our solution safely and easily

> Let's define our dependency as an interface. This lets us then use a real Sleeper in main and a spy sleeper in our tests.


## 10. concurrency

- race conditions test: ```go test -race```

Quotes:

> Everything you need to know is printed to your terminal - all you have to do is be patient enough to read it.

> Make it work, make it right, make it fast

## 11. select

- mock http testing server: 
```
httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
}))
```

- select case to run after timeout: ```case <-time.After(timeout)```

> ideally we don't want to be relying on external services to test our code

## 12. reflection

Quotes: 

> We use an anonymous struct with a Name field of type string to ***go for the simplest "happy" path***.

