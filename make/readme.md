# make

The `make` is a built-in function that allocates and initializes an object of type 1. slice, 2. map, or 3. chan (only).

-   It returns an **initialized** (not _zeroed_) value of type `T` (not `*T`).
-   These tree types represent, under the covers, references to data structures that **must be initialized** before use.

1. creating slices

    length must be passed, capacity can be omitted. if omitted, than equal to the length.

    ```go
    slice1 := make([]string, 3)    // Equivalent to make([]string, 3, 3)
    slice2 := make([]string, 3, 5) // length is 3, capacity is 5. capacity cannot be smaller than the length.
    ```

2. creating maps

    The size may be omitted.

    ```go
    m := make(map[string]int) // same as m := map[string]int{}
    ```

3. creating channels

    The size can be omitted. If omitted, the channel is unbuffered.
    The channel's buffer is initialized with the specified buffer capacity.

    ```go

    ```
