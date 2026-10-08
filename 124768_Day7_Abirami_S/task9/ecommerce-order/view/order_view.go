package view
import("bufio";"fmt";"os";"q09-ecommerce-order/controller";"strconv";"strings")
type OrderView struct{c controller.OrderController;r *bufio.Reader}
func NewOrderView(c controller.OrderController)*OrderView{return &OrderView{c,bufio.NewReader(os.Stdin)}}
func(v *OrderView)Start(){fmt.Print("Order ID: ");id,_:=strconv.Atoi(v.read());xs,e:=v.c.GetOrderDetails(id);if e!=nil{fmt.Println(e);return};for _,x:=range xs{fmt.Println(x)}}
func(v *OrderView)read()string{x,_:=v.r.ReadString('\n');return strings.TrimSpace(x)}
