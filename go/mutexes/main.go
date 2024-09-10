package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}

/*
Generics in Go
As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many cases. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't care about the types of values stored in the slice. Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
In Go 1.18 however, support for generics was released, effectively solving this problem!

Type Parameters
Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
In the example above, T is the name of the type parameter for the splitAnySlice function, and we've said that it must match the any constraint, which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.

firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)
Copy icon
Assignment
At Mailio we store all the emails for a campaign in memory as a slice. We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function that returns the last element from a slice, no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

Tip: Zero value of a type
Creating a variable that's the zero value of a type is easy:

var myZeroInt int
Copy icon
It's the same with generics, we just have a variable that represents the type:

var myZero T
*/

package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	m := map[int]int{}

	mu := &sync.RWMutex{}

	go writeLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)
	go readLoop(m, mu)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoop(m map[int]int, mu *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}
}

func readLoop(m map[int]int, mu *sync.RWMutex) {
	for {
		mu.RLock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mu.RUnlock()
	}
}

/// Generics

/*

Generics in Go
As we've mentioned, Go does not support classes. For a long time, that meant that Go code couldn't easily be reused in many cases. For example, imagine some code that splits a slice into 2 equal parts. The code that splits the slice doesn't care about the types of values stored in the slice. Before generics, we needed to write the same code for each type, which is a very un-DRY thing to do.

func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
In Go 1.18 however, support for generics was released, effectively solving this problem!

Type Parameters
Put simply, generics allow us to use variables to refer to specific types. This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.

func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}
Copy icon
In the example above, T is the name of the type parameter for the splitAnySlice function, and we've said that it must match the any constraint, which means it can be anything. This makes sense because the body of the function doesn't care about the types of things stored in the slice.

firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
fmt.Println(firstInts, secondInts)
Copy icon
Assignment
At Mailio we store all the emails for a campaign in memory as a slice. We store payments for a single user in the same way.

Complete the getLast() function. It should be a generic function that returns the last element from a slice, no matter the types stored in the slice. If the slice is empty, it should return the zero value of the type.

Tip: Zero value of a type
Creating a variable that's the zero value of a type is easy:

var myZeroInt int
Copy icon
It's the same with generics, we just have a variable that represents the type:

var myZero T*/


package main

func getLast[T any](s []T) T {
	// ?
}

/*
Why Generics?
Generics reduce repetitive code
You should care about generics because they mean you don’t have to write as much code! It can be frustrating to write the same logic over and over again, just because you have some underlying data types that are slightly different.

Generics are used more often in libraries and packages
Generics give Go developers an elegant way to write amazing utility packages. While you will see and use generics in application code, I think it will be much more common to see generics used in libraries and packages. Libraries and packages contain importable code intended to be used in many applications, so it makes sense to write them in a more abstract way. Generics are often the way to do just that!

Why did it take so long to get generics?
Go places an emphasis on simplicity. In other words, Go has purposefully left out many features to provide its best feature: being simple and easy to work with.

According to historical data from Go surveys, Go’s lack of generics has always been listed as one of the top three biggest issues with the language. At a certain point, the drawbacks associated with the lack of a feature like generics justify adding complexity to the language.
*/