using System;

class Vehicle
{
    public string Brand { get; }

    public Vehicle(string brand)
    {
        Brand = brand;
    }

    public virtual void Move()
    {
        Console.WriteLine($"{Brand} moving generically");
    }
}

class Car : Vehicle
{
    public int Doors { get; }

    public Car(string brand, int doors) : base(brand)
    {
        Doors = doors;
    }

    public override void Move()
    {
        Console.WriteLine($"{Brand} car driving with {Doors} doors");
    }
}

class Bicycle : Vehicle
{
    public Bicycle(string brand) : base(brand) { }
}

class Program
{
    static void Main()
    {
        Vehicle vehicle = new Vehicle("Generic");
        vehicle.Move();

        Vehicle car = new Car("Tesla", 4);
        car.Move();

        Vehicle bike = new Bicycle("Caloi");
        bike.Move(); // usa implementação base
    }
}
