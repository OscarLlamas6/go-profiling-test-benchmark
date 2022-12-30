# Golang | Testing, Profiling & Benchmarking 

Just practicing some profiling, testing and benchmark with Golang

```bash

# Install go tools
> go install golang.org/x/tools/cmd/godoc@latest

# Check golang documentation in console
> go doc testing

# Profiling an application
> go tool pprof -http=":9090" -seconds=15 http://localhost:8081/debug/pprof/profile

```
