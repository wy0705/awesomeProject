package design_mode
//工厂模式
type Pizz interface {
	Eat() string
}
type CheesePizza struct {

}
func (cheesepize CheesePizza) Eat() string{
	return "做好了芝士批萨"
}
type PepperPizza struct {

}

func (pepperPizza PepperPizza) Eat() string {
	return "做好了Prepper批萨"
}
type GreekPizze struct {

}
func (greekPizze GreekPizze) Eat() string{
	return "做好了GreekPizze批萨"
}
type SimplePizzaFactory struct {
	pizza Pizz
	ordertype string
}

func (simplePizzaFactory SimplePizzaFactory) CreatPizze()  Pizz{
	if simplePizzaFactory.ordertype=="cheese" {
		simplePizzaFactory.pizza=new(CheesePizza)
	}else if simplePizzaFactory.ordertype=="greek" {
		simplePizzaFactory.pizza=new(GreekPizze)
	}else if simplePizzaFactory.ordertype=="pepper" {
		simplePizzaFactory.pizza=new(PepperPizza)
	}
	return simplePizzaFactory.pizza
}

