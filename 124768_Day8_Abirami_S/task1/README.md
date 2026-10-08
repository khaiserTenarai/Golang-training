Commands Used:
1. dlv debug - Starts the Go program in the delve debugger.
2. break calculateSalary - Sets a breakpoint at the calculateSalary function
3. continue - Runs the program until it reaches the breakpoint.
4. print salary - Displays the current value of the salary
5. print bonus - Displays the current value of the bonus variable.
Identified the error in the salary bonus calculation. The program was using fixed bonus instead of calculating 10% bonus.

6. next - Executes the next line of code step by step.
7. exit - Exits the delve debugger.