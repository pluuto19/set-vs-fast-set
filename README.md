# Hash Set vs Bit Set

Benchmarking three set implementations in Go against each other using `int64` keys.

| Implementation | Backing structure |
|---|---|
| `SetBool` | `map[int64]bool` |
| `SetStruct` | `map[int64]struct{}` |
| `BitSet` | `[]int64` (bit array) |

## Run

```bash
go test -bench=. -benchmem
```
