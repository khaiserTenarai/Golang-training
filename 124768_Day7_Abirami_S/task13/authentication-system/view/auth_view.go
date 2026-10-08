package view
import("bufio";"fmt";"os";"q13-authentication-system/controller";"strconv";"strings")
type AuthView struct{c controller.AuthController;r *bufio.Reader}
func NewAuthView(c controller.AuthController)*AuthView{return &AuthView{c,bufio.NewReader(os.Stdin)}}
func(v *AuthView)Start(){for{fmt.Println("1.Register 2.Login 3.Exit");fmt.Print("Choice: ");switch v.i(){case 1:fmt.Print("Username: ");u:=v.s();fmt.Print("Password: ");p:=v.s();fmt.Print("Role: ");fmt.Println(v.c.Register(u,p,v.s()));case 2:fmt.Print("Username: ");u:=v.s();fmt.Print("Password: ");p:=v.s();x,e:=v.c.Login(u,p);if e!=nil{fmt.Println(e)}else{fmt.Println("Login successful:",x.Username,x.Role)};case 3:return}}}
func(v *AuthView)s()string{x,_:=v.r.ReadString('\n');return strings.TrimSpace(x)}
func(v *AuthView)i()int{x,_:=strconv.Atoi(v.s());return x}
