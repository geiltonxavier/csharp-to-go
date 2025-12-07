using System;
using System.Collections.Generic;

class Program
{
    static void Main()
    {
        int count = 3;
        var name = "CSharp";
        const double Pi = 3.14;
        bool isReady = true;

        int[] numbers = { 1, 2, 3 };
        var primes = new List<int> { 2, 3, 5 };
        var ages = new Dictionary<string, int>
        {
            { "ana", 29 },
            { "joao", 31 },
        };

        Console.WriteLine($"count: {count}, name: {name}, PI: {Pi}, ready: {isReady}");
        Console.WriteLine($"numbers length: {numbers.Length}, first: {numbers[0]}");

        primes.Add(7);
        Console.WriteLine($"primes: {string.Join(", ", primes)}");

        Console.WriteLine($"idade de ana: {ages["ana"]}");

        int zeroInt = default;
        string empty = default;
        Console.WriteLine($"default int: {zeroInt}, default string: {(empty == null ? "null" : empty)}");
    }
}
