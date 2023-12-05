defmodule Solution do
    def parse() do
        IO.read(:stdio, :line)
            |> String.trim
            |> String.split(",")
            |> Enum.map(fn timer -> String.to_integer(timer) end)
    end

    def tick(counters) do
        [birthdays | next_generation] = counters
        next_generation = next_generation
            |> Enum.with_index(0)
            |> Enum.map(fn {element, index} ->
                case index do
                    6 -> element + birthdays
                    _ -> element
                end
            end)
        next_generation ++ [birthdays]
    end

    def count(counters) do
        Enum.reduce(counters, 0, fn (count, acc) -> count + acc end)
    end


    def simulate(fishes, times) do
        counters = Enum.map(0..8, fn _ -> 0 end)
        counters = fishes
            |> Enum.reduce(counters, fn (count, acc) ->
                List.update_at(acc, count, fn v -> v + 1 end)
            end)
        Enum.reduce(1..times, counters, fn(_, generation) -> tick(generation) end)
    end

    def main() do
        counters = simulate(parse(), 256)
        IO.puts(Enum.reduce(counters, 0, fn (count, acc) -> count + acc end))
    end
end


Solution.main()




