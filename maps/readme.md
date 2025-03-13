# Maps

A map is an **unordered** group of elements of **one** type, called the **element type**, indexed by a set of unique keys of another type, called the **key type**.

The value of an unititialized map is `nil`, that is to say, the zero value of a map is `nil`. A `nil` map has no keys, nor can keys can be added:

```go
var nilMap map[string]int
```

## Map Keys

The comparison operators == and != must be fully defined for operands of the **key type**,; thus the key type must not be:

| Key Type       | Allowed in Map? | Reason                         |
| -------------- | --------------- | ------------------------------ |
| `bool`         | ✅ Yes          | Comparable                     |
| `int`          | ✅ Yes          | Comparable                     |
| `float64`      | ✅ Yes          | Comparable                     |
| `string`       | ✅ Yes          | Comparable                     |
| `struct{}`     | ✅ Yes          | Comparable (if all fields are) |
| `[N]T` (array) | ✅ Yes          | Comparable (if `T` is)         |
| `interface{}`  | ❌ No           | Not inherently comparable      |
| `slice`        | ❌ No           | Not comparable                 |
| `map`          | ❌ No           | Not comparable                 |
| `func`         | ❌ No           | Not comparable                 |

-   function
-   map
-   slice

-   If the top-level type is just a type name, you can omit it from the elements of the literal.

## Map Values

The map values can be any type.

## Empty Map

There are two ways to create an empty map:

1. Using map literal:

    ```go
    m := map[string]int{}
    ```

2. Using `make`:
   A new, empty map value si made using the built-in function `make`, which takes the map type and an optional capacity hint as arguments:

    ```go
    make(map[string]int)
    make(map[string]int, 100) // Capacity is 100
    ```

    The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of `nil` maps. A `nil` map is equivalent to an empty map except that no elements may be added.

## Map Length

The number of map elements is called its length.

For a map m, it can be discovered using the built-in function `len` and may change during execution.

## Maps are References

If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

```go
map1 := map[string]string{
	"brand": "Ford",
	"model": "Mustang",
	"year":  "1964",
}
map2 := map1

map2["brand"] = "mercedes"	// changes map1["brand"] too
delete(map2, "year")		// deletes "year" from map1 too
delete(map1, "model") 		// deletes "model" from map2 too
```

## Iteration

`range` can be used to iterate over maps.

```go
for key, value := range map[string]string{} {
	fmt.Printf("%v : %v, ", key, value)
}
```
