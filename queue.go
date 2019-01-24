package main

import (
	"fmt"
	"sync"
)

type Queue interface {
	Add(frame *Frame)
	Remove()*Frame
	Top()*Frame
	Less(func(data []*Frame,i,j int)bool)
	Size()int
}

type Frame struct {
	Priority int
	data     interface{}

}

type PriorityQueue struct {
	lesss func(data []*Frame,i,j int)bool
	ExpandSize int
	nodes []*Frame
	N int
}

func NewPrioQueue() Queue {
	p:= &PriorityQueue{N:1,ExpandSize:5}
	p.N=1
	return p
}

func (p *PriorityQueue)swim(k int)  {
	for k>1&&p.less(k, k/2) {
		p.exch(k/2, k)
		k /= 2
	}
}

func (p *PriorityQueue)sink(k int)  {
	if p.N==1{
		return
	}
	for 2*k < p.N {
		j := 2*k
		if j+1<p.N && p.less(j+1,j){
			j=j+1
		}

		if p.less(j,k){
			p.exch(k,j)
			k = j
		}else{
			break
		}
	}
}

func (p *PriorityQueue)Add(f *Frame)  {
	if len(p.nodes)<=p.N{
		apd:=make([]*Frame,p.N-len(p.nodes)+p.ExpandSize)
		p.nodes = append(p.nodes,apd...)
	}
	p.nodes[p.N]=f
	p.swim(p.N)
	p.N++
}

func (p *PriorityQueue)Remove() *Frame {
	if p.N==1{
		return nil
	}
	data:=p.nodes[1]
	p.nodes[1]=p.nodes[p.N-1]
	p.N--
	p.sink(1)
	return data
}

func (p *PriorityQueue)Top() *Frame  {
	if p.N<=1{
		return nil
	}
	return p.nodes[1]
}

func (p *PriorityQueue)Size()int  {
	return p.N-1
}

func (p *PriorityQueue)less(a,b int) bool {
	if p.lesss!=nil{
		return p.lesss(p.nodes,a,b)
	}
	return p.nodes[a].Priority <p.nodes[b].Priority
}

func (p *PriorityQueue)Less(f func(data []*Frame,i,j int)bool)  {
	p.lesss = f
}

func (p *PriorityQueue)show()  {

	for i:=1;i< len(p.nodes);i++{
		fmt.Print(p.nodes[i].Priority," ")
	}

	fmt.Println()
}

func (p *PriorityQueue)exch(i,j int)  {
	p.nodes[i],p.nodes[j]=p.nodes[j],p.nodes[i]
}

func NewBlockedPrioQueue()Queue  {
	pq:=&BlockedPrioQueue{
		pq:NewPrioQueue(),
		cond:sync.NewCond(&sync.Mutex{}),
	}
	return pq
}

type BlockedPrioQueue struct {
	cond *sync.Cond
	pq Queue
	hasdata bool
}

func (c *BlockedPrioQueue)Add(f *Frame)  {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	c.pq.Add(f)
	c.cond.Signal()
}

func (c *BlockedPrioQueue)Remove() *Frame {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()
	for c.Size()==0{
		c.cond.Wait()
	}
	return c.pq.Remove()
}

func (c *BlockedPrioQueue)Top()*Frame  {
	return c.Top()
}

func (c *BlockedPrioQueue)Less(f func(data []*Frame,i,j int)bool)  {
	 c.pq.Less(f)
}

func (c *BlockedPrioQueue)Size()int  {
	return c.pq.Size()
}