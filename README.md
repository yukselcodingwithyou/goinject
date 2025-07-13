# goinject
Dependency injection in golang

## Examples

See [examples/main.go](examples/main.go) for a simple usage example demonstrating how to register dependencies and inject them into a struct.
The `autowire` tag can optionally include a qualifier name to select a specific
dependency provided with `ProvideNamed`.

Run the example:

```bash
go run ./examples
```
