## Herança vs Embedding: C# → Go

Mostrando como a herança (override) em C# se traduz em composição/embedding em Go.

### Pontos de atenção
- C#: `class Car : Vehicle` herda estado/comportamento; `virtual` + `override` controlam polimorfismo.
- Go: não há herança; usamos embedding (`type Car struct { Vehicle }`) para promover campos/métodos e permitir reuse.
- Override: redefinir método em Go é apenas declarar na struct (`func (c Car) Move()`); quando chamado em `car.Move()`, usa a implementação do tipo concreto. Não há override dinâmico em valores `Vehicle` — o método chamado depende do tipo concreto em tempo de compilação.
- Campos promovidos: `bike.Brand` funciona porque `Vehicle` está embutido; ainda assim, é composição, não herança.
- Interfaces: para polimorfismo dinâmico em Go, preferimos programar contra interfaces, não supertipos.

### Como rodar
- C#: `dotnet run` em `examples/02-oop-to-composition/inheritance-vs-embedding/csharp`
- Go: `go run .` em `examples/02-oop-to-composition/inheritance-vs-embedding/go`
