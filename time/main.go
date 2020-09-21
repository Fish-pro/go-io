package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main(){
	/*
	一天=24小时 hour
	一小时=60分钟 minute
	一分钟=60秒 second
	一秒=1000毫秒 millisecond
	一毫秒=1000微秒 microsecond
	一微妙=1000纳秒 nanosecond
	一纳秒=1000皮秒 picosecond
	*/
	// 获取当前事件
	t1 := time.Now()
	fmt.Printf("%T\n",t1)
	fmt.Println(t1)

	t2 := time.Date(2020,10,1,12,12,40,0,time.Local)
	fmt.Println(t2)

	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	s2 := t1.Format("2006/1/2")
	fmt.Println(s2)

	s3 := "1999年7月29日"
	t3, err := time.Parse("2006年1月2日",s3)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(t3)

	year,month,day := t1.Date()
	fmt.Println(year,month,day)

	hour,min,sec := t1.Clock()
	fmt.Println(hour,min,sec)

	fmt.Println(t1.Year())
	fmt.Println(t1.YearDay())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())
	fmt.Println(t1.Hour())
	fmt.Println(t1.Minute())
	fmt.Println(t1.Second())
	fmt.Println(t1.Weekday())

	fmt.Println(t1.Add(24*time.Hour))
	fmt.Println(t1.AddDate(1,1,1))
	fmt.Println(t1.AddDate(1,1,1).Sub(t1))

	time.Sleep(3*time.Second)
	fmt.Println("end main")

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(9)+1)
	fmt.Println(time.Duration(rand.Intn(9)+1)*time.Second)
}
