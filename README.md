# queue
A fast, thread safe, queue implementation

## Usage

```go
package main

import "github.com/mbcrocci/queue"

type MyStruct struct {
  i int
  s string 
}

func main() {
  // initialize the queue
  q := queue.New[MyStruct]()

  ms := &MyStruct{}

  // Enqueue to the queue. If there is a consumer blocked reading 
  // it will unlock.
  q.Enqueue(ms)
  
  // Dequeue from the queue (blocking)
  s := q.Dequeue()
}
```
