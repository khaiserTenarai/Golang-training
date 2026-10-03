package view
import("bufio";"fmt";"os";"q10-order-inventory-transaction/controller";"q10-order-inventory-transaction/model";"strconv";"strings")
type OrderView struct{c controller.OrderController;r *bufio.Reader}
func NewOrderView(c controller.OrderController)*OrderView{return &OrderView{c,bufio.NewReader(os.Stdin)}}
func(v *OrderView)Start(){var o model.Order;fmt.Print("Product ID: ");o.ProductID=v.i();fmt.Print("Quantity: ");o.Quantity=v.i();if e:=v.c.CreateOrder(o);e!=nil{fmt.Println(e)}else{fmt.Println("Order created and stock reduced")}}
func(v *OrderView)read()string{x,_:=v.r.ReadString('\n');return strings.TrimSpace(x)}
func(v *OrderView)i()int{x,_:=strconv.Atoi(v.read());return x}
