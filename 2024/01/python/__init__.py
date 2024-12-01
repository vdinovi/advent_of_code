from pathlib import Path
from functools import reduce
from collections import Counter


def main():
    print("p1(example) =", p1(*read_lists(resolve_path("example.txt"))))
    print("P1(input)   =", p1(*read_lists(resolve_path("input.txt"))))
    print("p2(example) =", p2(*read_lists(resolve_path("example.txt"))))
    print("p2(input)   =", p2(*read_lists(resolve_path("input.txt"))))


def p1(left: list[int], right: list[int]) -> int:
    return sum(abs(pair[0] - pair[1]) for pair in zip(sorted(left), sorted(right)))


def p2(left: list[int], right: list[int]) -> int:
    counts = Counter(right)
    return sum(n * counts[n] for n in left)


def read_lists(path: Path) -> tuple[list[int], list[int]]:
    with open(path, "r") as f:
        left, right = zip(
            *(
                (int(l.strip()), int(r.strip()))
                for l, r in (line.split()[:3] for line in f.readlines())
            )
        )
        return list(left), list(right)


def resolve_path(filename: str, dir: Path = Path(__file__).parent) -> Path:
    return dir / filename


if __name__ == "__main__":
    main()
