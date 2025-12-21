## Task vs Goroutines: C# → Go

Requisições HTTP em paralelo: `Task` + `WhenAll` em C# comparado a goroutines + channels em Go.

### Pontos de atenção
- Concorrência: `Task.Run` ou métodos `async` em C# são similares a lançar goroutines com `go func() { ... }`.
- Sincronização: `Task.WhenAll` aguarda todas as tasks; em Go usamos `sync.WaitGroup` para sinalizar fim e um canal para coletar resultados.
- Erros: cada goroutine envia `err` no canal; não há exceção propagada automática — você decide o que fazer com cada resultado.
- Cancelamento/timeout: em C#, `CancellationToken`; em Go, `context.Context` (não mostrado aqui para simplificar, mas é o paralelo).
- HTTP client: `HttpClient` reutilizável em C#; em Go, `http.DefaultClient` é segura para uso concorrente, mas pode ser substituída se precisar de configuração.

### Como rodar
- C#: `dotnet run` em `examples/03-async-patterns/task-vs-goroutines/csharp`
- Go: `go run .` em `examples/03-async-patterns/task-vs-goroutines/go`
