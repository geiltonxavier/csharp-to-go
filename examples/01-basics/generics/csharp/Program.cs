using System;
using System.Collections.Generic;

class Box<T> where T : class
{
    public T Value { get; }

    public Box(T value)
    {
        Value = value ?? throw new ArgumentNullException(nameof(value));
    }

    public void Print()
    {
        Console.WriteLine($"boxed value: {Value}");
    }
}

static class Helpers
{
    public static void PrintAll<T>(IEnumerable<T> items)
    {
        foreach (var item in items)
        {
            Console.WriteLine($"item ({typeof(T).Name}): {item}");
        }
    }
}

class Program
{
    static void Main()
    {
        var textBox = new Box<string>("hello C#");
        textBox.Print();

        var numbers = new List<int> { 1, 2, 3 };
        Helpers.PrintAll(numbers);

        var words = new[] { "go", "csharp" };
        Helpers.PrintAll(words);

        // var valueBox = new Box<int>(5); // não compila por causa do constraint "class"
    }
}
