package main

import "fmt"

type item struct {
	p *item // previous
	n *item // next
	k int   // key
	v int   // value
}

type LRUCache struct {
	c   map[int]*item // cache
	h   *item         // head
	t   *item         // tail
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		c:   make(map[int]*item, capacity),
	}
}

func (this *LRUCache) rebalance(i *item) {
	switch {
	case i == this.t:
		// do nothing, this is the last used key anyway
	case i == this.h:
		if i.n != nil {
			// tOrig := this.t

			// this.h = i.n
			// tOrig.n = i
			// i.p = tOrig
			// this.t = i
			// i.n = nil

			this.h, this.t.n, i.p, this.t, i.n = i.n, i, this.t, i, nil
		}
	default:
		// 	tOrig := this.t
		// 	prev := i.p
		// 	next := i.n

		// 	i.p = tOrig
		// 	tOrig.n = i
		// 	i.n = nil
		// 	prev.n = next
		// 	next.p = prev
		// 	this.t = i

		i.p, this.t.n, i.n, i.p.n, i.n.p, this.t = this.t, i, nil, i.n, i.p, i
	}

}
func (this *LRUCache) Get(key int) int {
	if i, ok := this.c[key]; ok {
		this.rebalance(i)
		return i.v
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	i, ok := this.c[key]

	if ok {
		i.v = value
		this.rebalance(i)
		return
	}

	if len(this.c) == this.cap {
		delete(this.c, this.h.k)
		i = this.h
		this.rebalance(i)
	} else {
		i = &item{
			p: this.t,
		}

		// the first append
		if len(this.c) == 0 {
			this.h = i
		}

		if this.t != nil {
			this.t.n = i
		}

		this.t = i

	}

	i.k = key
	i.v = value
	this.c[key] = i
}

func (this *LRUCache) String() string {
	res := fmt.Sprintf("LRUCAche: len = %v, head = %v, tail = %v\n", len(this.c), this.h.v, this.t.v)
	cnt := 0
	head := this.h
	for head != nil {
		pv, nv := 0, 0
		if head.n != nil {
			nv = head.n.v
		}
		if head.p != nil {
			pv = head.p.v
		}

		res += fmt.Sprintf("%v[%v:%v] -> ", head.k, pv, nv)
		head = head.n
		cnt++
		if cnt > this.cap+1 {
			break
		}
	}
	return res
}

func main() {
	//	cmds := []string{"LRUCache", "put", "put", "put", "put", "get", "get", "get", "get", "put", "get", "get", "get", "get", "get"}
	//	vals := [][]int{{3}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {4}, {3}, {2}, {1}, {5, 5}, {1}, {2}, {3}, {4}, {5}}

	cmds := []string{"LRUCache", "put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"}

	vals := [][]int{{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}}

	c := Constructor(1)
	for i := 0; i < len(cmds); i++ {
		fmt.Println("------------------------------")
		cmd := cmds[i]
		v := vals[i]
		switch cmd {
		case "LRUCache":
			c = Constructor(v[0])
		case "put":
			fmt.Printf("put %v\n", v[0])
			c.Put(v[0], v[1])
			fmt.Println(c.String())
		case "get":
			fmt.Println(c.String())
			fmt.Printf("get %v = ", v[0])
			res := c.Get(v[0])
			fmt.Printf("%v\n", res)
		}
	}

	/*
		lRUCache := Constructor(2)
		lRUCache.Put(1, 1) // cache is {1=1}
		fmt.Println(lRUCache.String())
		lRUCache.Put(2, 2) // cache is {1=1, 2=2}
		fmt.Println(lRUCache.String())
		fmt.Println(lRUCache.Get(1)) // return 1
		lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
		fmt.Println(lRUCache.String())
		fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
		lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
		fmt.Println(lRUCache.String())
		fmt.Println(lRUCache.Get(1)) // return -1 (not found)
		fmt.Println(lRUCache.Get(3)) // return 3
		fmt.Println(lRUCache.Get(4)) // return 4
	*/
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
