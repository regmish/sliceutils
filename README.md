# Slice Utils
- [x] Length
- [x] Push
- [x] Pop
- [x] ForEach
- [x] Map
- [x] Filter
- [x] IndexOf
- [x] Includes
- [] Sort
- [] Concat
- [] Every
- [] Join
- [] Reduce
- [] Find
- [] FindIndex
- [] Fill
- [] Slice
- [] Reverse
- [] Shift
- [] Unshift


Frustration of not having Cool Array methods in GO

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