# Testing Go Fuzzing

As of Go 1.19.3, Go fuzzing does not seem to search int arguments very effectively. I believe this is due to this Go issue: https://github.com/golang/go/issues/48291

This code attempts to test this. The file `gofuzztesting_test.go` includes three fuzz tests that test a function which crashes if an int argument is > 5000. My assumption is this should be relatively easy to find failing test cases. The three tests are:

* `FuzzInt`: Takes one `int` argument with no corpus (no calls to `f.Add`). Does not find failing cases.
* `FuzzCorpus`: Use `f.Add()` before calling the function. If the argument is "too far", it doesn't help. However, if the argument is close, it will find "nearby" failures quickly. This one finds failures quickly.
* `FuzzBytes`: Takes a []byte argument and casts it into an int. This seems to trigger more reliably, even if I change the function to things that need specific values like `arg == 12345`. This does not find specific "large" values like `arg == 123456789`.

Test it:

```
go test . -fuzz=FuzzInt -fuzztime=1m
```

This was tested with Go 1.19.3 on 2022-11-15.

## Workarounds

* Seed the corpus with `f.Add`
* Using `[]byte` might be better than `int`, at least for some workloads