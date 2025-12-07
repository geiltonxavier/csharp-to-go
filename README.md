## csharp-to-go: Rosetta Stone para devs C#

Um repositório interativo que coloca C# e Go lado a lado. Cada conceito traz um exemplo C# conhecido e sua contraparte em Go, com explicações curtas e código executável.

### Como usar
- Entre em cada exemplo em `examples/` para ver C# (`csharp/`) e Go (`go/`) na prática.
- Execute ambos os lados para sentir as diferenças: `dotnet run` no diretório `csharp/` e `go run .` em `go/`.
- Leia o `README.md` local de cada exemplo para entender a tradução de conceitos.

### Roadmap de exemplos
- `01-basics/` variáveis/tipos, interfaces, generics, tratamento de erros.
- `02-oop-to-composition/` herança vs embedding, polimorfismo, injeção de dependência.
- `03-async-patterns/` Task vs goroutines, async/await vs channels, cancellation tokens vs context.
- `04-linq-to-go/` filtros/map, agregações, iteradores customizados.
- `05-common-patterns/` repository, unit of work, middleware pipeline, builder.
- `06-real-world/` Web API, EF Core vs sqlx, DI prática.
- `challenges/` exercícios para migrar serviços passo a passo.

### Convenções
- Estrutura lado a lado: `examples/<bloco>/<tema>/csharp/` e `go/`.
- Tudo com README curto explicando escolhas idiomáticas de Go para quem vem de C#.
- Commit por implementação: cada exemplo vive como uma pequena unidade independente.

### Requisitos locais
- .NET 7+ (`dotnet run`).
- Go 1.21+ (`go run .`).
