# Structs

## Struct

-   A `struct` is a typed collection of `fields`.
-   Each field has a name and a type.
-   Non-blank field names must be unique.
-   A struct can be empty. `type emptyStruct struct{}`

## Struct Fields

-   Struct fields are accessed using a dot.
-   Struct fields can be accessed through a struct pointer.

    -   To access the field `X` of a struct when we have the struct pointer `p` we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

    -   `p.X` is same as `(*p).X`

-   A struct literal denotes a newly allocated struct value by listing the values of its fields.

    You can list just a subset of fields by using the `Name:` syntax, and the order of the named fields is irrelevant. If the field is not present, it is set to the zero value for its type.

    ```go
    var (
    	v1 = Vertex{1, 2}  // has type Vertex
    	v2 = Vertex{X: 1}  // Y:0 is implicit
    	v3 = Vertex{}      // X:0 and Y:0
    	p  = &Vertex{1, 2} // has type *Vertex
    )
    ```

-   A struct is mutable, meaning that its fields can be changed after the struct is created.

    ```go
    p := person{name: "Bob", age: 20}
    p.age = 40
    ```

-   A struct can be anonymous.

    ```go
    p := struct {
        name string
        age  int
    }{name: "Bob", age: 20}
    ```

## Relationship between structs and functions

-   Structs are passed by value.

    ```go
    func growOld(p person) {
    	p.age = p.age + 1
    }
    john := person{name: "john", age: 41}
    growOld(john) // john is still 41
    ```

-   If the struct is large → Return a pointer (\*person) to avoid unnecessary copying.
-   If the struct is small and copying isn’t an issue → Return a value (person) for better memory locality.
-   If you need mutability (modify after creation) → Use a pointer.
-   You can't accidentally return `nil` or pass a `nil`where a structure is expected, but you can with a pointer.
-   Prefer accepting a struct over a pointer unless you need to make side-effects that the caller (or others) can see.

    ```go
    type person struct {
    	name string
    	age  int
    }
    // Accept struct as value instead of pointer
    func printName(p person) {
    	fmt.Println(p.name)
    }
    ```
