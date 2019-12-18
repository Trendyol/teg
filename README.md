# Teg

Teg is a library for manage feature toggles. It aims to allow to create and access to feature toggles in Golang easily and quickly!

## Architecture

Teg reads feature toggles from lock-free in-memory storage and sync to in-memory storage from `store.Reader`.

## Storage Backends

Teg comes with builtin storage backends. You can also use any storage backend with `store.Reader` interface.

```go
type Reader interface {
    Get(name string) (*FeatureToggle, error)
    GetAll() (FeatureToggles, error)
}
```

### Builtin Storage Backends

* Environment Variable
* Git Repository
* InMemory

## Update Triggers

Teg comes with builtin update triggers. You can also use any update triggers with `TriggerFunc` easily.

```go
type TriggerFunc func(context.Context, chan<- struct{}) error
```

### Builtin Update Triggers

* Periodic Trigger
* RabbitMQ

## Example

If you want to use custom storage and custom update triggers checkout examples folder.

## Status

Teg is usable but still under active development. We expect it to be production-ready in the near future.

## Contributing

* If you want to contribute to codes, create pull request
* If you find any bugs or error, create an issue

## License

This project is licensed under the MIT Lıcense
