[![Go Reference](https://pkg.go.dev/badge/github.com/granddaifuku/generic-ordered-map.svg)](https://pkg.go.dev/github.com/granddaifuku/generic-ordered-map)
# generic-ordered-map
`generic-ordered-map` is a Linked Hash Map with Go1.18's generics features.
By internally using `container/list` and `map`, `generic-ordered-map` allows users to maintain the insertion order.

Note: This package requires [Go 1.18](https://go.dev/blog/go1.18).

## Feature

### Map
| Operation   | Order                              | Notes                                                 | 
| :---------: | :--------------------------------: | :---------------------------------------------------: | 
| Set         | O(1)                               |                                                       | 
| Get         | O(1)                               |                                                       | 
| Delete      | O(1)                               |                                                       | 
| Keys        | O(N)                               | Earn all the keys<br>(Insertion-order)                | 
| Values      | O(N)                               | Earn all the values<br>(Insertion-order)              | 
| Entries     | O(N)                               | Earn all the key and value pairs<br>(Insertion-order) | 
| FromEntries | O(M)<br>M : The length of the args | Add key and value pairs to map                        | 
| Merge       | O(M)                               | Merge the other map                                   | 
| Front       | O(1)                               | Earn the first element                                | 
| Back        | O(1)                               | Earn the last element                                 | 

### Set
`generic-ordered-map` provides `Set` feature by wrapping its `Map` implementation.
That is `Set` is defined as below.
```go
type Set[T comparable] struct {
	mp *Map[T, struct{}]
}
```

| Operation   | Order                              | Notes                                                 | 
| :---------: | :--------------------------------: | :---------------------------------------------------: | 
| Set         | O(1)                               |                                                       | 
| Contains    | O(1)                               |                                                       | 
| Delete      | O(1)                               |                                                       | 
| Values      | O(N)                               | Earn all the values<br>(Insertion-order)              | 
| FromValues  | O(M)                               |                                                       | 
| Merge       | O(M)                               | Merge the other map                                   | 
| Front       | O(1)                               | Earn the first element                                | 
| Back        | O(1)                               | Earn the last element                                 | 
| Intersect   | O(min(N, M))                       | Return the result of the intersection                 |
| Union       | O(min(N, M))                       | Return the result of the union                        |


### Iterating
```go
for ele := mp.Front(); ele != nil; ele = ele.Next() {
    // Iterate the elements from oldest to latest
}

for ele := mp.Back(); ele != nil; ele = ele.Prev() {
    // Iterate the elements from latest to oldest
}
```

## Install
```sh
$ go get -u github.com/granddaifuku/generic-ordered-map
```

## Usage
See `examples` directory.(TBD)

## Benchmarks
TBD
