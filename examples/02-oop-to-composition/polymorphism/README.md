## Polimorfismo via Interface: C# → Go

Uso de interface comum e implementação implícita para despachar chamadas de método em coleção heterogênea.

### Pontos de atenção
- Interface: `IShape` em C# ↔ `type Shape interface` em Go. Em Go, basta que o tipo tenha o método para satisfazer a interface.
- Coleção heterogênea: `List<IShape>` ↔ `[]Shape`. A fat interface em Go guarda (tipo concreto + valor) permitindo chamada dinâmica.
- Métodos sem estado mutável: receivers por valor (`func (r Rectangle) Area()`) são suficientes aqui; use ponteiro se o método alterar o estado.
- Nomeação: funções utilitárias podem receber `[]Shape` para operar sobre qualquer implementação, similar a `IEnumerable<IShape>`.
- Performance: interfaces têm custo de boxing semelhante a armazenar referências em C#. Para alta performance, slices de structs concretas evitam interface.

### Como rodar
- C#: `dotnet run` em `examples/02-oop-to-composition/polymorphism/csharp`
- Go: `go run .` em `examples/02-oop-to-composition/polymorphism/go`
