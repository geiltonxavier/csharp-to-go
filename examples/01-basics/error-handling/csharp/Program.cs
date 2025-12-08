using System;

class Program
{
    static void Main()
    {
        try
        {
            var parsed = ParseOrThrow("42");
            Console.WriteLine($"parsed number: {parsed}");
        }
        catch (FormatException ex)
        {
            Console.WriteLine($"falha ao converter: {ex.Message}");
        }

        try
        {
            var parsed = ParseOrThrow("abc");
            Console.WriteLine($"parsed number: {parsed}");
        }
        catch (FormatException ex)
        {
            Console.WriteLine($"falha ao converter: {ex.Message}");
        }

        try
        {
            var result = Divide(10, 2);
            Console.WriteLine($"10 / 2 = {result}");

            result = Divide(10, 0);
            Console.WriteLine($"10 / 0 = {result}");
        }
        catch (ArgumentException ex)
        {
            Console.WriteLine($"erro ao dividir: {ex.Message}");
        }
    }

    static int ParseOrThrow(string value)
    {
        return int.Parse(value);
    }

    static double Divide(int a, int b)
    {
        if (b == 0)
        {
            throw new ArgumentException("divisão por zero não permitida", nameof(b));
        }

        return (double)a / b;
    }
}
