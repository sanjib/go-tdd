package main

import (
	// Testing for 8.dependency-injection
	//di "./8.dependency-injection"
	//"net/http"
	//"os"

	// Testing for 9.mocking
	mo "./9.mocking"
	"os"
	"time"
)

func main() {
	// Uncomment relevant sections to run

	// Testing for 8.dependency-injection
	//di.Greet(os.Stdout, "Gopher")
	//_ = http.ListenAndServe(":5000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	di.Greet(w, "Beastie")
	//}))

	// Testing for 9.mocking
	sleeper := &mo.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mo.Countdown(os.Stdout, sleeper)
}
