上台阶多少中方法的问题，一次上1个台阶或一次上两个台阶，归纳问题，最后是求剩余1个台阶+剩余2个台阶的走法，使用递归求解。

```go
func f1(n int) int {
  if n == 1 {
    return 1
  }
  if n == 2 {
    return 2
  }
  return f1(n-1) + f1(n-2)
}

// 优化：
func f2(n int) int {
  if n == 1 {
    return 1
  }
  x := 1
  y := 2
  for(i:=3; i<=n; i++) {
    z := x+y
    x := y
    y := z
  }
  return y
}
```
