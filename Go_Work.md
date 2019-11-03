Go 语言并发，使用无缓冲通道构建并发池。

首先构建一个并发池，有两个字段一个是work  通道类型，一个是sync.WaitGroup。

```
type Pool struct {   
	work chan Worker   
	wg sync.WaitGroup
}
```

为工作并发池 提供一个接口

```
type Worker interface {
	Task()
}
```

接着创建一个新的工作池，使用工厂函数来创建,返回值为Pool

```
func New(maxGoroutines int) *Pool  {
	p := Pool{
		work : make(chan Worker),
	}
	p.wg.Add(maxGorotuine) //往并发池添加并发的数量
	for i:= 0 ; i < maxGoroutines; i ++{
		go func() {
			for w := range p.work{   // 这里会发生阻塞，在创建的时候，异步阻塞，返回一个池
				w.Task()
			}
		}()
		p.wg.Done()
	}
	return &p
	
}
```

将工作提交到工作池

```
func (p *Pool)Run(w Worker)  {
	p.work <- w
}
```

关闭工作池

```
func (p *Pool)Shutdown(){   
	close(p.work)   
	p.wg.Wait()
}
```

接下来是对上面的使用

```
package main

import (
	"fmt"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter)Task()  {
	fmt.Println(m.name)
	time.Sleep(time.Second )
}


func main(){
	p := New(10) //创建10个并发池子
	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	// time.Sleep(3 * time.Second)

	for i :=0 ; i < 10 ; i ++{
		for _ , name := range names{
			np := namePrinter{
				name : name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
```

