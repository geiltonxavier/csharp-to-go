using System;
using System.Collections.Generic;

interface IShape
{
    double Area();
}

class Rectangle : IShape
{
    public double Width { get; }
    public double Height { get; }

    public Rectangle(double width, double height)
    {
        Width = width;
        Height = height;
    }

    public double Area() => Width * Height;
}

class Circle : IShape
{
    public double Radius { get; private set; }

    public Circle(double radius)
    {
        Radius = radius;
    }

    public double Area() => Math.PI * Radius * Radius;
}

static class ShapePrinter
{
    public static void PrintAreas(IEnumerable<IShape> shapes)
    {
        foreach (var shape in shapes)
        {
            Console.WriteLine($"area: {shape.Area():0.00}");
        }
    }
}

class Program
{
    static void Main()
    {
        var shapes = new List<IShape>
        {
            new Rectangle(3, 4),
            new Circle(2)
        };

        ShapePrinter.PrintAreas(shapes);
    }
}
