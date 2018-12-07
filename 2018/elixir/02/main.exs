#!/usr/bin/env elixir

defmodule Exercise do
  def main do
    values = File.open!(Path.dirname(__ENV__.file) <> "/input.txt", [:read])
    |> IO.stream(:line)
    |> Enum.map(&readline/1)

    {_, mult} = calc_line(values, [])
    |> Enum.map_reduce(%{}, fn x, acc -> {x, Map.update(acc, x, 1, &(&1 + 1))} end)

    {_, res} = Enum.map_reduce(Map.values(mult), 1, fn x, acc -> {x, x * acc} end)
    IO.puts("RESULT: " <> Integer.to_string(res))
  end

  defp calc_line([v|tail], res) do
    new_res = String.split(v, "", trim: true)
    |> count(%{})
    |> Map.values
    |> Enum.uniq
    |> Enum.filter(fn x -> x > 1 end)

    calc_line(tail, res ++ new_res)
  end
  defp calc_line([], res), do: res

  defp count([letter|tail], res) do
    {_, res} = Map.get_and_update(res, letter, fn curr ->
      {curr, add_one(curr)}
    end)
    count(tail, res)
  end
  defp count([], res), do: res

  defp add_one(nil), do: 1
  defp add_one(val), do: val+1

  defp readline(line) do
    String.trim(line)
  end

end

Exercise.main()
