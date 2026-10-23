package graphql

import (
	"net/http"
	"strings"
	"strconv"

	"employee-platform/model"
	"employee-platform/service"
	gql "github.com/graphql-go/graphql"
	"github.com/gin-gonic/gin"
)

const PlaygroundHTML = `<!doctype html><html><head><meta charset="utf-8"><title>Employee GraphQL Playground</title><style>body{font:16px Arial;margin:35px;max-width:1000px}textarea{width:100%;height:240px;font:14px monospace}button{padding:10px;margin:8px 0}pre{background:#f3f3f3;padding:14px;white-space:pre-wrap}</style></head><body><h2>Employee GraphQL Playground</h2><p>Endpoint: <code>/graphql</code>. For protected operations, add an Authorization: Bearer token header.</p><textarea id="q">query { employees(page: 1, limit: 10) { total data { id name email age salary department { id name } } } }</textarea><p><textarea id="v" style="height:70px">{"Authorization":"Bearer YOUR_TOKEN"}</textarea></p><button onclick="run()">Run query</button><pre id="out"></pre><script>async function run(){try{let h=JSON.parse(document.getElementById('v').value||'{}');let r=await fetch('/graphql',{method:'POST',headers:{'Content-Type':'application/json',...h},body:JSON.stringify({query:document.getElementById('q').value})});document.getElementById('out').textContent=JSON.stringify(await r.json(),null,2)}catch(e){document.getElementById('out').textContent=String(e)}}</script></body></html>`

type API struct { Service *service.Service; Schema gql.Schema }
func New(s *service.Service)*API{
	departmentType:=gql.NewObject(gql.ObjectConfig{Name:"Department",Fields:gql.Fields{
		"id":&gql.Field{Type:gql.NewNonNull(gql.ID)},"name":&gql.Field{Type:gql.NewNonNull(gql.String)},
	}})
	employeeType:=gql.NewObject(gql.ObjectConfig{Name:"Employee",Fields:gql.Fields{
		"id":&gql.Field{Type:gql.NewNonNull(gql.ID)},"name":&gql.Field{Type:gql.NewNonNull(gql.String)},"email":&gql.Field{Type:gql.NewNonNull(gql.String)},
		"age":&gql.Field{Type:gql.NewNonNull(gql.Int)},"salary":&gql.Field{Type:gql.NewNonNull(gql.Float)},
		"department":&gql.Field{Type:departmentType,Resolve:func(p gql.ResolveParams)(interface{},error){e,ok:=p.Source.(model.Employee);if !ok{return nil,nil};if e.Department!=nil{return e.Department,nil};return s.GetDepartment(p.Context,e.DepartmentID)}},
	}})
	inputType:=gql.NewInputObject(gql.InputObjectConfig{Name:"EmployeeInput",Fields:gql.InputObjectConfigFieldMap{
		"name":&gql.InputObjectFieldConfig{Type:gql.NewNonNull(gql.String)},"email":&gql.InputObjectFieldConfig{Type:gql.NewNonNull(gql.String)},
		"age":&gql.InputObjectFieldConfig{Type:gql.NewNonNull(gql.Int)},"salary":&gql.InputObjectFieldConfig{Type:gql.NewNonNull(gql.Float)},"department_id":&gql.InputObjectFieldConfig{Type:gql.NewNonNull(gql.ID)},
	}})
	pageType:=gql.NewObject(gql.ObjectConfig{Name:"EmployeePage",Fields:gql.Fields{"total":&gql.Field{Type:gql.NewNonNull(gql.Int)},"page":&gql.Field{Type:gql.NewNonNull(gql.Int)},"limit":&gql.Field{Type:gql.NewNonNull(gql.Int)},"data":&gql.Field{Type:gql.NewNonNull(gql.NewList(gql.NewNonNull(employeeType)))}}})
	query:=gql.NewObject(gql.ObjectConfig{Name:"Query",Fields:gql.Fields{
		"employees":&gql.Field{Type:pageType,Args:gql.FieldConfigArgument{"page":&gql.ArgumentConfig{Type:gql.Int},"limit":&gql.ArgumentConfig{Type:gql.Int},"name":&gql.ArgumentConfig{Type:gql.String},"department":&gql.ArgumentConfig{Type:gql.String},"sortBy":&gql.ArgumentConfig{Type:gql.String},"order":&gql.ArgumentConfig{Type:gql.String}},Resolve:func(p gql.ResolveParams)(interface{},error){
			page:=argInt(p.Args,"page",1);limit:=argInt(p.Args,"limit",10)
			data,total,err:=s.ListEmployees(p.Context,page,limit,argString(p.Args,"name"),argString(p.Args,"department"),defaultString(argString(p.Args,"sortBy"),"id"),defaultString(argString(p.Args,"order"),"ASC"))
			return map[string]interface{}{"data":data,"total":total,"page":page,"limit":limit},err
		}},
		"employee":&gql.Field{Type:employeeType,Args:gql.FieldConfigArgument{"id":&gql.ArgumentConfig{Type:gql.NewNonNull(gql.ID)}},Resolve:func(p gql.ResolveParams)(interface{},error){id,e:=strconv.ParseInt(p.Args["id"].(string),10,64);if e!=nil{return nil,e};v,e:=s.GetEmployee(p.Context,id);if e!=nil{return nil,nil};return v,nil}},
		"departments":&gql.Field{Type:gql.NewList(departmentType),Resolve:func(p gql.ResolveParams)(interface{},error){return s.ListDepartments(p.Context)}},
	}})
	mutation:=gql.NewObject(gql.ObjectConfig{Name:"Mutation",Fields:gql.Fields{
		"createEmployee":&gql.Field{Type:employeeType,Args:gql.FieldConfigArgument{"input":&gql.ArgumentConfig{Type:gql.NewNonNull(inputType)}},Resolve:func(p gql.ResolveParams)(interface{},error){in,e:=inputFrom(p.Args["input"]);if e!=nil{return nil,e};return s.CreateEmployee(p.Context,in)}},
		"updateEmployee":&gql.Field{Type:employeeType,Args:gql.FieldConfigArgument{"id":&gql.ArgumentConfig{Type:gql.NewNonNull(gql.ID)},"input":&gql.ArgumentConfig{Type:gql.NewNonNull(inputType)}},Resolve:func(p gql.ResolveParams)(interface{},error){id,e:=strconv.ParseInt(p.Args["id"].(string),10,64);if e!=nil{return nil,e};in,e:=inputFrom(p.Args["input"]);if e!=nil{return nil,e};return s.UpdateEmployee(p.Context,id,in)}},
		"deleteEmployee":&gql.Field{Type:gql.Boolean,Args:gql.FieldConfigArgument{"id":&gql.ArgumentConfig{Type:gql.NewNonNull(gql.ID)}},Resolve:func(p gql.ResolveParams)(interface{},error){id,e:=strconv.ParseInt(p.Args["id"].(string),10,64);if e!=nil{return false,e};if e=s.DeleteEmployee(p.Context,id);e!=nil{return false,e};return true,nil}},
	}})
	schema,err:=gql.NewSchema(gql.SchemaConfig{Query:query,Mutation:mutation});if err!=nil{panic(err)}
	return &API{Service:s,Schema:schema}
}
func(a *API)Serve(c *gin.Context){
	var req struct{Query string `json:"query"`;Variables map[string]interface{} `json:"variables"`}
	if err:=c.ShouldBindJSON(&req);err!=nil{c.JSON(400,gin.H{"errors":[]interface{}{gin.H{"message":"invalid JSON request"}}});return}
	// All GraphQL requests require JWT. Mutations require the admin role.
	if strings.Contains(strings.ToLower(req.Query), "mutation") {
		role, _ := c.Get("role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"errors": []interface{}{gin.H{"message":"admin role required for mutations"}}})
			return
		}
	}
	result:=gql.Do(gql.Params{Schema:a.Schema,RequestString:req.Query,VariableValues:req.Variables,Context:c.Request.Context()})
	c.JSON(http.StatusOK,result)
}
func inputFrom(v interface{})(model.EmployeeInput,error){
	m,ok:=v.(map[string]interface{});if !ok{return model.EmployeeInput{},http.ErrBodyNotAllowed}
	age,_:=number(m["age"]);salary,_:=number(m["salary"]);dept,_:=number(m["department_id"])
	return model.EmployeeInput{Name:asString(m["name"]),Email:asString(m["email"]),Age:int(age),Salary:salary,DepartmentID:int64(dept)},nil
}
func number(v interface{})(float64,error){switch n:=v.(type){case float64:return n,nil;case int:return float64(n),nil;case string:f,e:=strconv.ParseFloat(n,64);return f,e};return 0,nil}
func asString(v interface{})string{s,_:=v.(string);return s}
func argInt(m map[string]interface{},k string,d int)int{if v,ok:=m[k];ok{if n,e:=number(v);e==nil{return int(n)}};return d}
func argString(m map[string]interface{},k string)string{s,_:=m[k].(string);return s}
func defaultString(s,d string)string{if s==""{return d};return s}
