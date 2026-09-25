package main
import "fmt"
func main(){
	fact(5)
	fib(5)
	prime(5)
	rev(123)
	palindrome(121)
}

func fact(n int){
	factorial :=1
	for i:=1;i<=n;i++{
		factorial *= i;
	}
	fmt.Println("Factorial",factorial)
}

func fib(num int){
	a:=0
	b:=1
	fmt.Println("Fibonacci:")
	for i:=0;i<=num;i++{
		fmt.Println(a," ")
		temp:=a
		a=b
		b=temp
	}
	fmt.Println()
	
}
func prime(n int){
	count:=0
	for i:=1;i<=n;i++{
		if n%i==0{
			count++
		}
	}
	if count ==2{
		fmt.Println(n,"is Prime Number")
	}else{
		fmt.Println(n,"is not Prime Number")
	}
}

func rev(n int){
	reverse :=0
	for n>0{
		x :=n%10
		reverse=reverse*10+x
		n=n/10
	}
	fmt.Println("Reverse of the given number is:",reverse)

}

func palindrome(n int){
	temp :=n
	reverse :=0
	for n>0{
		x :=n%10
		reverse=reverse*10+x
		n=n/10
	}
	if temp==reverse{
		fmt.Println(temp," is Palindrome")
	}else{
		fmt.Println(temp," is not a Palindrome")
	}


}