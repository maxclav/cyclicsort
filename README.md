# Cyclic Sort

Simple cyclic sort of slice of ints in Go (GoLang).  
For learning purpose.

## Usage

```go
// Sort a valid slice
s1 := []int{4, 1, 3, 0, 5, 2}
SortValid(s1)
fmt.Println(s1) // Prints []int{0, 1, 2, 3, 4, 5}

// Check if a slice is valid
s2 := []int{0, 1, 2, 3, 4, 5}
fmt.Println(IsValid(s2)) // Prints true
s3 := []int{1, 2, 3, 4, 5, 6}
fmt.Println(IsValid(s3)) // Prints false
s4 := []int{0, 1, 2, 3, 4, 50}
fmt.Println(IsValid(s4)) // Prints false

// Sort a slice if it's valid
s5 := []int{4, 1, 3, 0, 5, 2}
Sort(s5)
fmt.Println(s5) // Prints []int{0, 1, 2, 3, 4, 5}
s6 := []int{40, 123, 333, 0, -5, 321}
Sort(s6)
fmt.Println(s6) // Prints []int{40, 123, 333, 0, -5, 321}
```

## Commands

- `make` or `make help`: display all commands
- `make run-lint`: run linter
- `make run-tests`: run all tests
  - `make run-unit-tests`: run unit tests
