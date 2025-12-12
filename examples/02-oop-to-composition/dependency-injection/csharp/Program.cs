using System;

interface IMessageSender
{
    void Send(string to, string message);
}

class EmailSender : IMessageSender
{
    public void Send(string to, string message)
    {
        Console.WriteLine($"[email] para {to}: {message}");
    }
}

interface IClock
{
    DateTime Now();
}

class SystemClock : IClock
{
    public DateTime Now() => DateTime.UtcNow;
}

class NotificationService
{
    private readonly IMessageSender _sender;
    private readonly IClock _clock;

    public NotificationService(IMessageSender sender, IClock clock)
    {
        _sender = sender;
        _clock = clock;
    }

    public void Notify(string user, string message)
    {
        var timestamp = _clock.Now();
        _sender.Send(user, $"[{timestamp:O}] {message}");
    }
}

class Program
{
    static void Main()
    {
        // Composition root: manual wiring (poderia ser um container de DI em C#)
        IMessageSender sender = new EmailSender();
        IClock clock = new SystemClock();
        var notifier = new NotificationService(sender, clock);

        notifier.Notify("ana", "olá usando DI manual em C#");
    }
}
