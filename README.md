
#### introduce
Priority Queue that implements by golang based on heap sort.<br>
Queue interface has two implements:<br>
    
```
type Queue interface {
	Add(data interface{})
	Remove() interface{}
	Top() interface{}
	Size() int
}

```
    
implement|goroutine secure|blocked
---|---|---
PriorityQueue|No|No
BlockingPriorityQueue|Yes|Yes
