package main
 
import (
//  "log"
"fmt"
)

 
type Observable struct {
	Name string
}
 
type TestCallBack struct {
	Name string
  
}
 
func (c *TestCallBack) Exec(o *Observable) {
//	log.Printf(c.name+ ":")
//  log.Println(o.Name)
fmt.Printf(c.Name + ": ")
fmt.Printf(o.Name + "\n")
}
 
type Callback interface {
	Exec(h *Observable)
}
 
type Observer struct {
	callbacks []Callback
}
 
func (o *Observer) Add(c Callback) {
	o.callbacks = append(o.callbacks, c)
}

 func (o *Observer) Delete(c Callback) {
 	var i = 0
	o.callbacks, o.callbacks[len(o.callbacks)-1] = append(o.callbacks[:i], o.callbacks[i+1:]...), nil
}
func (o *Observer) Process(oe *Observable) {
	for _, c := range o.callbacks {
		c.Exec(oe)
	}
}
 
func main() {
  oe := Observable{Name: "Hello World"}
  testCall := []*TestCallBack{&TestCallBack{Name: "test"},&TestCallBack{Name: "test1"}}
  o := Observer{}
//  o.Add(&TestCallBack{name: "test"})
  o.Add(testCall[0])
  o.Add(testCall[1])
  o.Process(&oe)
  o.Delete(testCall[0])
  eo := Observable{Name: "Test1"}
  o.Process(&eo)
  
}