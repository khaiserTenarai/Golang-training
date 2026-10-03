package main

import "fmt"

func add(a, b int) int {
	sum := a - b // Bug
	return sum
}

func main() {
	x := 10
	y := 5

	result := add(x, y)

	fmt.Println("Result:", result)
}

/*
DEBUGGING STEPS USING DELVE

1. Run the program:
   go run .

   Output:
   Result: 5

2. Start Delve:
   dlv debug

3. Set breakpoint:
   break main.add

4. Continue the program:
   continue

5. Check variables:
   print a
   Output: 10

   print b
   Output: 5

6. Move to the next line:
   next

7. Check sum:
   print sum

   First it showed:
   Command failed: could not find symbol value for sum

   After using next again:

   print sum
   Output: 5

8. The returned value was:
   Values returned:
       ~r0: 5

9. Bug found:
   sum := a - b

10. Correct code:
    sum := a + b

11. Run the program again:
    go run .

12. Correct output:
    Result: 15

CONCLUSION:
The bug was caused by using '-' instead of '+'.
The program was subtracting 5 from 10 instead of adding them.
*/
