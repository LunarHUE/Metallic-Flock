```bash
go install github.com/bufbuild/buf/cmd/buf@latest
```

## Getting the vendor hash

```bash
go mod vendor && nix hash path vendor
```
