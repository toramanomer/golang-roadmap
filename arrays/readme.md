# Arrays

-   An **array** is a numbered sequence of elements of a **signel type**, called the **element type**.
-   The number of elements is called the **length** of the array and is **never** negative.
-   The **length** is part of the array's type; it must evaluate to a non-negative constant representable by a value of type `int`.

-   Arrays are **values**. Assigning one array to another copies all the elements.
-   In particular, if you pass an array to a function, it will receive a **copy** of the array, not a pointer to it.
-   The size of an array is part of its type. The types [10]int and [20]int are distinct.

## Array Declaration

1. `var` keyword

    ```go
    var array_name1 = [3]int{1, 2, 3}   // length is defined
    var array_name2 = [...]int{1, 2, 3} // length is inferred
    ```

2. `:=` sign

    ```go
    array_name1 = [3]int{1, 2, 3}   // length is defined
    array_name2 = [...]int{1, 2, 3} // length is inferred
    ```

## Array Initialization

If an array or one of its elements has not been initialized in the code, it is assigned the zero value of its type.

```go
arr1 := [3]int{}        // not initialized, equivalent is [3]int{0, 0, 0}
arr2 := [3]int{1, 2}    // partially initialized, equivalent is [3]{1, 2, 0}
arr3 := [3]int{1, 2, 3} // fully initialized
arr4 := [3]int{2: 1}	// partially initialized, equivalent is [3]{0, 0, 1}
```
