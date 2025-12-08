## Tratamento de erros: C# → Go

Comparação entre `try/catch` e o estilo explícito de retorno de erro em Go.

### Pontos de atenção
- Fluxo: em C# você envolve o bloco com `try { ... } catch (Exception ex) { ... }`. Em Go, funções retornam `(valor, error)` e você decide a cada chamada o que fazer.
- Sem exceções em fluxo normal: Go evita throw/catch para controle; erros são valores (`error`) e podem ser inspecionados com `errors.Is/As`.
- Wrapping: `fmt.Errorf("msg: %w", err)` preserva a causa (similar a `throw new InvalidOperationException("msg", ex)`).
- Sentinelas: `ErrDivideByZero` é um erro bem conhecido; use `errors.Is` para comparar em vez de checar mensagens.
- Parse vs throw: `int.Parse` lança `FormatException`; em Go `strconv.Atoi` retorna `error` — checado logo após a chamada.

### Como rodar
- C#: `dotnet run` em `examples/01-basics/error-handling/csharp`
- Go: `go run .` em `examples/01-basics/error-handling/go`
