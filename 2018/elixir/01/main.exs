#!/usr/bin/env elixir

defmodule Exercise do
  def main do
    values = File.open!(Path.dirname(__ENV__.file) <> "/input.txt", [:read])
    |> IO.stream(:line)
    |> Enum.map(&readline/1)

    IO.puts("SUM: " <> Integer.to_string(Enum.sum(values)))

    calculate_twice(values, 0, %{0 => 0}, values)
  end

  defp readline(line) do
    String.trim(line)
    |> String.to_integer
  end

  defp calculate_twice([v | tail], sum, acc, init) do
    sum = sum + v
    if Map.has_key?(acc, sum) do
      IO.puts("Twice: " <> Integer.to_string(sum))
    else
      calculate_twice(tail, sum, Map.merge(acc, %{sum => sum}), init)
    end
  end

  defp calculate_twice([], sum, acc, init) do
    calculate_twice(init, sum, acc, init)
  end
end

Exercise.main()
