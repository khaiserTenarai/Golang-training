package view
import("bufio";"fmt";"os";"strconv";"strings";"q08-money-transfer/controller")
type AccountView struct{c controller.AccountController;r *bufio.Reader}
func NewAccountView(c controller.AccountController)*AccountView{return &AccountView{c,bufio.NewReader(os.Stdin)}}
func(v *AccountView)Start(){for{fmt.Println("1.Balance 2.Transfer 3.Exit");fmt.Print("Choice: ");switch v.i(){case 1:fmt.Print("Account ID: ");a,e:=v.c.GetAccount(v.i());if e!=nil{fmt.Println(e)}else{fmt.Println(a)};case 2:fmt.Print("From: ");a:=v.i();fmt.Print("To: ");b:=v.i();fmt.Print("Amount: ");x:=v.f();if e:=v.c.Transfer(a,b,x);e!=nil{fmt.Println(e)};case 3:return}}}
func(v *AccountView)read()string{x,_:=v.r.ReadString('\n');return strings.TrimSpace(x)}
func(v *AccountView)i()int{x,_:=strconv.Atoi(v.read());return x}
func(v *AccountView)f()float64{x,_:=strconv.ParseFloat(v.read(),64);return x}
