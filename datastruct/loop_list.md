```go
type a struct {
   val  int
   next *a
}
```
定义两个指针p1,p2

从链表头开始，p1一次走一步，p2一次走两步

p1,p2逐次遍历，当p1和p2相遇时表明该列表为循环列表
