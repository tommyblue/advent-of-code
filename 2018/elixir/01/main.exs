#!/usr/bin/env elixir

defmodule Exercise do
  def main do
    File.open!(Path.dirname(__ENV__.file) <> "/input.txt", [:read])
    |> IO.stream(:line)
    |> Enum.map(&readline/1)
    |> Enum.sum
    |> IO.puts
  end

  defp readline(line) do
    String.trim(line)
    |> String.to_integer
  end
end

Exercise.main()
