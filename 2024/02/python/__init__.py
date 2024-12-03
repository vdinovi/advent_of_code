import typing
import pathlib


def main():
    print("p1(example) =", p1(read_lists(resolve_path("example.txt"))))
    print("P1(input)   =", p1(read_lists(resolve_path("input.txt"))))
    print("p2(example) =", p2(read_lists(resolve_path("example.txt"))))
    print("p2(input)   =", p2(read_lists(resolve_path("input.txt"))))


def p1(lists: list[list[int]], lower=1, upper=3) -> int:
    return sum(1 if is_safe(l, lower, upper) else 0 for l in lists)


def p2(lists: list[list[int]], lower=1, upper=3) -> int:
    def sublists(l: list[int]) -> typing.Iterator[list[int]]:
        for i in range(len(l)):
            yield l[:i] + l[i + 1 :]

    return sum(
        1 if any(is_safe(sl, lower, upper) for sl in sublists(l)) else 0 for l in lists
    )


def is_safe(l: list[int], lower: int, upper: int) -> bool:
    if l[1] - l[0] < 0:
        l = l[::-1]
    return all(lower <= cur - prev <= upper for prev, cur in zip(l, l[1:]))


def read_lists(path: pathlib.Path) -> list[list[int]]:
    with open(path, "r") as f:
        return [[int(n.strip()) for n in line.split()] for line in f.readlines()]


def resolve_path(
    filename: str, dir: pathlib.Path = pathlib.Path(__file__).parent
) -> pathlib.Path:
    return dir.parent / filename


if __name__ == "__main__":
    main()
