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

class SmsSender : IMessageSender
{
    public void Send(string to, string message)
    {
        Console.WriteLine($"[sms] para {to}: {message}");
    }
}

class Notifier
{
    private readonly IMessageSender _sender;

    public Notifier(IMessageSender sender)
    {
        _sender = sender;
    }

    public void Notify(string user, string message)
    {
        _sender.Send(user, message);
    }
}

class Program
{
    static void Main()
    {
        IMessageSender sender = new EmailSender();
        var notifier = new Notifier(sender);
        notifier.Notify("ana", "olá do C#");

        sender = new SmsSender();
        notifier = new Notifier(sender);
        notifier.Notify("joao", "olá por SMS");
    }
}
