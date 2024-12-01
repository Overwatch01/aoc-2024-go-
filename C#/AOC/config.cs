namespace AOC;

public static class Config
{
    private const string FilePath = "../../input.txt";

    public static IEnumerable<string> ReadInput()
    {
        IEnumerable<string> inputs = new string[] { };
        try
        {
            using var reader = new StreamReader(FilePath);
            while (reader.ReadLine() is { } line) inputs = inputs.Append(line);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"An error occurred: {ex.Message}");
        }

        return inputs;
    }

    public static int GetDifference(int a, int b)
    {
        if (a > b)
            return a - b;
        return b - a;
    }
}