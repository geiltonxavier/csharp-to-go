## Interfaces: C# → Go

Comparação direta de contrato + implementação e consumo via injeção simples.

### Pontos de atenção
- Declaração: `interface IMessageSender { ... }` em C# vs `type MessageSender interface { ... }` em Go.
- Implementação: em C# você declara `: IMessageSender`; em Go basta ter os métodos — implementação é implícita, sem palavra-chave.
- Injeção/composição: o padrão de passar a dependência no construtor se traduz para uma função `NewX` em Go retornando a struct com a interface.
- Valor vs referência: interfaces em Go são valores; ao armazenar a interface você guarda um par (tipo concreto + valor). Structs sem ponteiro são copiadas, mas aqui só chamamos método.
- Métodos com receiver: `func (EmailSender) Send(...)` é equivalente a um método de instância; use ponteiro (`*T`) se precisar mutar estado.

### Como rodar
- C#: `dotnet run` em `examples/01-basics/interfaces/csharp`
- Go: `go run .` em `examples/01-basics/interfaces/go`
