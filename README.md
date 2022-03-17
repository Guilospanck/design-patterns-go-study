# Design Patterns in Go
[![CircleCI](https://circleci.com/gh/Guilospanck/design-patterns-go-study/tree/main.svg?style=svg)](https://circleci.com/gh/Guilospanck/design-patterns-go-study/tree/main)
[![codecov](https://codecov.io/gh/Guilospanck/design-patterns-go-study/branch/main/graph/badge.svg?token=3R0NIG0J0J)](https://codecov.io/gh/Guilospanck/design-patterns-go-study)

Studies following [Refactoring Guru](https://refactoring.guru/design-patterns/go)

## Remember
A pattern isn’t just a recipe for structuring your code in a specific way. It can also communicate to other developers the problem the pattern solves.

### Tests
- Unit tests:
```bash
make test
```
or
```bash
go test ./... -race -coverprofile=coverage.out -covermode=atomic
```

- Coverage
```bash
make test-cov
```
or
```bash
go test ./... -race -coverprofile=coverage.out -covermode=atomic && go tool cover -html=./coverage.out -o coverage.html
```


