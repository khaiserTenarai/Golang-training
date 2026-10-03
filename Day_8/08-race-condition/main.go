package main

import (
	"fmt"
	"time"
)

var counter int

/*
A race condition happens when multiple goroutines access and modify the same variable at the same time, and the final result depends on the timing of execution.
*/

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Not atomic
	}
}

func main() {
	go increment()
	go increment()

	time.Sleep(2 * time.Second)

	fmt.Println("Counter:", counter)
}


/*
Execution and Output :
-------------------------

PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\08-race-condition> where.exe go
C:\Program Files\Go\bin\go.exe
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\08-race-condition> $env:Path = "C:\Program Files\Go\bin;C:\msys64\ucrt64\bin;$env:Path"
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\08-race-condition> go version
go version go1.27.1 windows/amd64
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\08-race-condition> go env CGO_ENABLED
1
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\08-race-condition> go run -race main.go
==================
WARNING: DATA RACE
Read at 0x000140261a30 by goroutine 8:
  main.increment()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:16 +0x2c

Previous write at 0x000140261a30 by goroutine 9:
  main.increment()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:16 +0x44

Goroutine 8 (running) created at:
  main.main()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:21 +0x27

Goroutine 9 (finished) created at:
  main.main()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:22 +0x33
==================
==================
WARNING: DATA RACE
Read at 0x000140261a30 by main goroutine:
  main.main()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:26 +0x79

Previous write at 0x000140261a30 by goroutine 8:
  main.increment()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:16 +0x44

Goroutine 8 (finished) created at:
  main.main()
      C:/Training/Go Lang/Day_8/124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy/08-race-condition/main.go:21 +0x27
==================
Counter: 2000
Found 2 data race(s)
exit status 66
*/