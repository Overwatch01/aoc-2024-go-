// See https://aka.ms/new-console-template for more information

using AOC;

var inputs = Config.ReadInput();
IEnumerable<int> leftArr = new int[] {};
IEnumerable<int> rightArr = new int[] { };

foreach (var line in inputs)
{
    var vals = line.Split(" ", StringSplitOptions.RemoveEmptyEntries);
    var leftInput = Convert.ToInt32(vals[0]);
    var rightInput = Convert.ToInt32(vals[1]);
    leftArr = leftArr.Append(leftInput);
    rightArr = rightArr.Append(rightInput);
}

leftArr = leftArr.Order();
rightArr = rightArr.Order();
var totalDistanceCovered = 0;
var totalSimilarityScore = 0;
for (int i = 0; i < leftArr.Count(); i++)
{
    totalDistanceCovered += Config.GetDifference(leftArr.ToArray()[i], rightArr.ToArray()[i]);
    var similarityScore = leftArr.ToArray()[i] * rightArr.Where(x => x == leftArr.ToArray()[i]).Count();
    totalSimilarityScore += similarityScore;
}

Console.WriteLine($"Totals Distance covered is {totalDistanceCovered}");
Console.WriteLine($"Totals Similarity score is {totalSimilarityScore}");