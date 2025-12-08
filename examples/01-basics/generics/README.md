## Generics: C# → Go

Exemplo com tipo genérico e função genérica para mostrar a sintaxe e constraints.

### Pontos de atenção
- Sintaxe: `Box<T>` em C# ↔ `type Box[T any] struct` em Go. Use `[T any]` para declarar o parâmetro de tipo.
- Constraints: `where T : class` restringe a referência em C#. Em Go, o `any` é sem restrições; para restringir use type sets (`interface { ~int | ~float64 }`, `comparable` etc.).
- Implementação implícita continua valendo: um tipo concreto satisfaz a interface se tem os métodos, mesmo com genéricos envolvidos.
- Métodos genéricos: `PrintAll<T>(IEnumerable<T>)` ↔ `func PrintAll[T any]([]T)`.
- Inferência: ambos inferem o tipo a partir do argumento (`PrintAll(numbers)`), mas Go exige escrever os colchetes na declaração.

### Como rodar
- C#: `dotnet run` em `examples/01-basics/generics/csharp`
- Go: `go run .` em `examples/01-basics/generics/go`
