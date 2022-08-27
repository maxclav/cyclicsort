# Cyclic Sort

Simple cyclic sort of slice of ints in [Go](https://go.dev/) (GoLang).  
For learning purpose.

## Usage

### Sort a slice if it's valid

```go
var err error

s1 := []int{4, 1, 3, 0, 5, 2}
err = Sort(s1)
fmt.Println(s1)   // Prints []int{0, 1, 2, 3, 4, 5}
fmt.Println(err)  // Prints nil

s2 := []int{40, 123, 333, 0, -5, 321}
err = Sort(s2)
fmt.Println(s2)   // Prints []int{40, 123, 333, 0, -5, 321}
fmt.Println(err)  // Prints error
```

### Check if a slice is valid

```go
s1 := []int{0, 1, 2, 3, 4, 5}
fmt.Println(Validate(s1)) // Prints nil

s2 := []int{0, 1, 2, 3, 4, 4}
fmt.Println(Validate(s2)) // Prints error

s3 := []int{1, 2, 3, 4, 5, 6}
fmt.Println(Validate(s3)) // Prints error

s4 := []int{0, 1, 2, 3, 4, 50}
fmt.Println(Validate(s4)) // Prints error
```

### Sort a valid slice

```go
// Sort a valid slice
s := []int{4, 1, 3, 0, 5, 2}
SortValid(s)
fmt.Println(s) // Prints []int{0, 1, 2, 3, 4, 5}
```

### Must sort a slice (or panic)

```go
s1 := []int{4, 1, 3, 0, 5, 2}
MustSort(s1)    // Doesn't panic
fmt.Println(s1) // Prints []int{0, 1, 2, 3, 4, 5}

s2 := []int{40, 123, 333, 0, -5, 321}
MustSort(s2) // PANIC
```

## Commands

- `make` or `make help`: display all commands
- `make run-lint`: run linter
- `make run-tests`: run all tests
  - `make run-unit-tests`: run unit tests
