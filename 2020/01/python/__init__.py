from pathlib import Path


def main():
    print("p1(example) =", p1(read_numbers(resolve_path("example.txt")), 2020))
    print("P1(input)   =", p1(read_numbers(resolve_path("input.txt")), 2020))
    print("p2(example) =", p2(read_numbers(resolve_path("example.txt")), 2020))
    print("P2(input)   =", p2(read_numbers(resolve_path("input.txt")), 2020))


def p1(numbers: list[int], target: int) -> int | None:
    try:
        return next((n * (target - n) for n in set(numbers) if (target - n) in numbers))
    except StopIteration:
        return None


def p2(numbers: list[int], target: int) -> int | None:
    for i, n in enumerate(numbers):
        rest = numbers[:i] + numbers[i + 1 :]
        m = p1(rest, target - n)
        if m is not None:
            return n * m
    return None


def read_numbers(path: Path) -> list[int]:
    with open(path, "r") as f:
        return [int(line) for line in f]


def resolve_path(filename: str, dir: Path = Path(__file__).parent) -> Path:
    return dir / filename


if __name__ == "__main__":
    main()
