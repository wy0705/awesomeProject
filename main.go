package main

/*type Person struct{
	id int
	name string
	age int
}
type Cou struct {
	name string
	stus [10]*Stu
}
type Stu struct{
	Person
	teas [3]*Tea

}
type Tea struct {
	Person
	coursename string

}

func byTeaFindStu(t Tea,allCourse [3]*Cou)  {
	fmt.Println(t.name+" have students ")
	for i := 0; i < 3; i++ {
		if  allCourse[i].name==t.coursename{
			for i := 0; i < 10; i++ {
				if allCourse[i].stus[i]==nil {
					return
				}
				fmt.Println(allCourse[i].stus[i].name);
			}
		}

	}
}

func byStuFindTeaAndCou(s Stu)  {
	fmt.Println(s.name+" have teacher and course ")
	for i := 0; i < 3; i++ {
		fmt.Println(s.teas[i].name+s.teas[i].coursename)
	}
}
func byCouFindStu(c Cou){
	fmt.Println(c.name+"have students ")
	for i := 0; i < 10; i++ {
		if c.stus[i]==nil {
			return
		}
		fmt.Println(c.stus[i].name);
	}
}
func main()  {
	c1:=new(Cou)
	c2:=new(Cou)
	c3:=new(Cou)

	t1:=new(Tea)
	t2:=new(Tea)
	t3:=new(Tea)
	t4:=new(Tea)
	t1=&Tea{Person{1,"李老师",33},"语文"}
	t2=&Tea{Person{2,"王老师",32},"语文"}
	t3=&Tea{Person{3,"郭老师",33},"数学"}
	t4=&Tea{Person{4,"陈老师",31},"英语"}

	s1:=new(Stu)
	s2:=new(Stu)
	s3:=new(Stu)
	s4:=new(Stu)
	s5:=new(Stu)
	s6:=new(Stu)
	s7:=new(Stu)
	s8:=new(Stu)
	s9:=new(Stu)

	//建立三个分组
	m1:=[10]*Stu{new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu)}
	*m1[0]=*s1
	*m1[1]=*s2
	*m1[2]=*s3
	*m1[3]=*s4
	*m1[4]=*s5
	*m1[5]=*s6
	m2:=[10]*Stu{new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu)}
	*m2[0]=*s4
	*m2[1]=*s5
	*m2[2]=*s6
	*m2[3]=*s7
	*m2[4]=*s8
	*m2[5]=*s9
	m3:=[10]*Stu{new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu),new(Stu)}
	*m3[0]=*s3
	*m3[1]=*s4
	*m3[2]=*s5
	*m3[3]=*s6
	*m3[4]=*s7
	*m3[5]=*s8

	c1=&Cou{"语文",m1}
	c2=&Cou{"数学",m2}
	c3=&Cou{"英语",m3}

	v1:=[3]*Tea{new(Tea),new(Tea),new(Tea)}
	*v1[0]=*t1
	*v1[1]=*t3
	*v1[2]=*t4
	v2:=[3]*Tea{new(Tea),new(Tea),new(Tea)}
	*v2[0]=*t2
	*v2[1]=*t3
	*v2[2]=*t4


	s1=&Stu{Person{1,"小a",13},v1}
	s2=&Stu{Person{1,"小b",13},v1}
	s3=&Stu{Person{1,"小c",13},v1}
	s4=&Stu{Person{1,"小d",13},v1}
	s5=&Stu{Person{1,"小e",13},v2}
	s6=&Stu{Person{1,"小f",13},v2}
	s7=&Stu{Person{1,"小g",13},v2}
	s8=&Stu{Person{1,"小h",13},v2}
	s9=&Stu{Person{1,"小q",13},v2}

	allCourse:=[3]*Cou{new(Cou),new(Cou),new(Cou)}
	*allCourse[0]=*c1
	*allCourse[1]=*c2
	*allCourse[2]=*c3

    //byTeaFindStu(*t1,allCourse)
	//byCouFindStu(*c1)
	byStuFindTeaAndCou(*s1)



}*/
func main()  {
	balance:= [20]byte{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t'}
}
