package priority_queue

import (
	"fmt"
	"testing"
	"time"
)

func TestNewPrioQueue(t *testing.T) {
	q:=NewPrioQueue(func(data []interface{}, i, j int) bool {
		return data[i].(int) <data[j].(int)
	});

	q.Add(3)
	q.Add(4)
	q.Add(1)
	q.Add(0)
	q.Add(7)
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	q.Add(8)
	q.Add(4)
	q.Add(5)
	q.Add(2)
	q.Add(1)
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
}

func TestNewBlockedPrioQueue(t *testing.T) {
	q:= NewBlockingPriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].(int) <data[j].(int)
	})
	go func() {
		for i:=0;i<100;i++{
			//go func(v int) {
			//	q.Add(v)
			//	fmt.Println(v)
			//}(i)
			//time.Sleep(5*time.Millisecond)
			q.Add(i)
			time.Sleep(50*time.Millisecond)
		}
		time.Sleep(10*time.Second)
		for i:=0;i<100;i++{
			//go func(v int) {
			//	q.Add(v)
			//	fmt.Println(v)
			//}(i)
			//time.Sleep(5*time.Millisecond)
			q.Add(i)
			time.Sleep(50*time.Millisecond)
		}
	}()


	//time.Sleep(1*time.Second)
	for i:=0;i<200;i++{

		fmt.Println("sorted:",q.Remove())
	}
}


func TestNewBlockedPrioQueue2(t *testing.T) {
	q:= NewBlockingPriorityQueue(func(data []interface{}, i, j int) bool {
		return data[i].(int) <data[j].(int)
	})
	go func() {
		for i:=0;i<100;i++{
			go func(v int) {
				q.Add(v)
				fmt.Println(v)
			}(i)
			//time.Sleep(5*time.Millisecond)
		//	q.Add(i)
		//	time.Sleep(50*time.Millisecond)
		}
		time.Sleep(10*time.Second)

	}()


	//time.Sleep(1*time.Second)
	for i:=0;i<100;i++{

		fmt.Println("sorted:",q.Remove())
		time.Sleep(50*time.Millisecond)
	}
}