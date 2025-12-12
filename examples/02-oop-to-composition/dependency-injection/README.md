## Injeção de Dependência: C# → Go

Comparação de injeção via construtor e o “composition root” sem container em Go.

### Pontos de atenção
- C#: DI costuma usar container (ex.: `AddScoped/AddSingleton`). Aqui mostramos a forma explícita: `new NotificationService(new EmailSender(), new SystemClock())`.
- Go: não há container na stdlib; a prática idiomática é wiring manual via construtores (`NewX(dep1, dep2)`) e interfaces para costurar as dependências.
- Interfaces como seam: `MessageSender` e `Clock` permitem trocar implementações em testes sem reconfigurar container.
- Ciclo de vida: em Go você controla quando instanciar e compartilhar; não há scoping automático como `Scoped`/`Transient`.
- Legibilidade: composition root explícito facilita ver de onde vem cada dependência; frameworks escondem isso mas custam reflexão e mágica.

### Como rodar
- C#: `dotnet run` em `examples/02-oop-to-composition/dependency-injection/csharp`
- Go: `go run .` em `examples/02-oop-to-composition/dependency-injection/go`
