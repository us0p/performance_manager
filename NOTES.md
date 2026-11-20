# Last stop
- Consolidate learnings about Clean Arquitecture with notes
- Add proper client interaction with the expected input source, Presenters and Views and how all of this ties back with HTMX.

## Reflection
- `main.go` is the composition root(?). Build use cases, controllers and repositories here and they wire then together in the right order so that they're accessible within the app.
- repository interfaces should be defined in the use cases layer since that's where they're going to be used.
- making each rote a method under `App` is fine for small applications. Must separate into one handler struct per feature, each holding only the use cases it needs as the app increases.
- The controler should take clean inputs and have no reference to frameworks and it must return a presentation of the data ready to be displayed.
- **Compostion root** is the single place where all the concrete objects are created and connected to each other. In a go application that's usually `main.go`.
Since inner circles depend only on abstractions, there must be a place that know the concrete type and actually wire them together. The composition root is where that decision goes.
It's the only component allowed to know every concrete type, so all the other components don't have to. The goal is for everything outside main to depend on abstractions and never 
construct its own dependencies. That gives you:
    - Testability: a test can compose the same use cases with fakes, effectively acting as a second composition root.
    - Replaceability: swapping Postgres or net/http means changing the wiring in one file.
    - Visibility: the whole dependency graph can be read in one place, from top to bottom.
    - Pluggability: in Bob's words, main is a plugin to the application. You can have one main for production, another for dev, and another for a CLI, all reusing the same inner layers.
Rules of thumb and anti-patterns:
- Keep only one, as close to the entry point as possible. Building objects in many places scatters your wiring decisions.
- Build the graph once at startup. Don't create repositories per request inside handlers.
- Don't let use cases construct their own dependencies. For example, NewCreateUser() should not call pgrepo.New... internally, because that bypasses the root.
- Avoid a service locator. Don't pass the whole App or a container around so that code can pull whatever it needs. Pass each component only the dependencies it actually uses.
- You don't need a DI container in Go. Manual wiring in main is idiomatic and easy to follow. Libraries like wire or fx only start to help when the graph becomes very large.
