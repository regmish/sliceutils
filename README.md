# Slice Utils: Useful Array Methods (Inspired from JavaScript)
- [x] Length
- [x] Push
- [x] Pop
- [x] ForEach
- [x] Map
- [x] Filter
- [x] IndexOf
- [x] Includes
- [x] Sort
- [x] Concat
- [x] Every
- [x] Join
- [x] Reduce
- [ ] Find
- [ ] FindIndex
- [ ] Fill
- [ ] Slice
- [ ] Reverse
- [ ] Shift
- [ ] Unshift


```go
myArr := Array[int]{1, 2, 3, 4, 5}

myArr.Length()
myArr.Push(6)
el := myArr.Pop()
myArr.ForEach(func(el int, idx int, arr Array) {
    fmt.PrintLn(el)
})

myArr.Filter(func(el int, idx int, arr Array) {
    return el > 3
})
..
..
and many many more...
```
