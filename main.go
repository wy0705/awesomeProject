package main

func main(){
	type Person struct {
		a int//8
		name string//16=8+8
	}
	//继承
	type Stu struct{
		/*p Person*/
		Person
	}
	type stu Stu
	m:=make(stu,3)
	s1:=new(Stu)
	m[0]=s1
	/*s:=new(Stu)


	s=&Stu{Person{a:1,name:"xxx"}}
	fmt.Println(s.a,s.name)*/
	/*p:=new(Person)
	p=&Person{a:1,name:"xxx"}
	fmt.Println(unsafe.Sizeof(*p))//24*/

	/*m1:=make(map[string] func() int,2)
	m1[">1"]= func() int {
		return 10
	}
	m1["=="]= func() int {
		return 1
	}
	var j int
	switch i:=1; {
	    case i==1:
	        j=m1["=="]()
	    case i>1:
	    	j=m1[">1"]()
	}
	fmt.Println(j)*/
}
