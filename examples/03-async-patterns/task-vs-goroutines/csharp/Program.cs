using System;
using System.Diagnostics;
using System.Net.Http;
using System.Threading.Tasks;

class Program
{
    static async Task Main()
    {
        using var http = new HttpClient();

        var sw = Stopwatch.StartNew();

        var googleTask = FetchLengthAsync(http, "https://www.google.com");
        var bingTask = FetchLengthAsync(http, "https://www.bing.com");

        var lengths = await Task.WhenAll(googleTask, bingTask);
        sw.Stop();

        Console.WriteLine($"google length: {lengths[0]}");
        Console.WriteLine($"bing length:   {lengths[1]}");
        Console.WriteLine($"elapsed: {sw.ElapsedMilliseconds}ms");
    }

    static async Task<int> FetchLengthAsync(HttpClient http, string url)
    {
        var response = await http.GetAsync(url);
        response.EnsureSuccessStatusCode();
        var content = await response.Content.ReadAsStringAsync();
        return content.Length;
    }
}
