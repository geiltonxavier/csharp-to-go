## Variáveis e Tipos: C# → Go

Exemplo rápido para quem vem de C#: declaração explícita, inferência de tipo, constantes, coleções básicas e valores padrão.

### Pontos de atenção
- `var` (C#) ≈ `:=` (Go) para inferência; use `var x T` em Go quando precisar do tipo explícito ou zero value.
- `const` existe nos dois, mas em Go só para valores compilados (nada de DateTime.Now).
- Arrays fixos vs slices: `int[]` em C# se comporta mais como um slice (`[]int`) em Go; arrays em Go têm tamanho no tipo (`[3]int`).
- `Dictionary` ↔ `map`: ambos são referência; se um `map` não foi inicializado, acessar/escrever causa panic.
- Valores padrão: `default(T)` pode ser `null` em C#; em Go o zero value nunca é `nil` para tipos básicos (`string` vira `""`, `int` vira `0`).
- String interpolation vs `fmt.Printf`: em Go você formata com verbos (`%d`, `%s`); `fmt.Sprint` ajuda a imprimir slices.

### Como rodar
- C#: `dotnet run` em `examples/01-basics/variables-and-types/csharp`
- Go: `go run .` em `examples/01-basics/variables-and-types/go`
