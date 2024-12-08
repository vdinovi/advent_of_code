import typing
import pathlib
import numpy as np

if typing.TYPE_CHECKING:
    from numpy.typing import NDArray


def main():
    print("p1(example) =", p1(*read_updates(resolve_path("example.txt"))))
    # print("P1(input)   =", p1(*read_file(resolve_path("input.txt"))))
    # print("p2(example) =", p2(*read_file(resolve_path("example.txt"))))
    # print("p2(input)   =", p2(*read_file(resolve_path("input.txt"))))


def p1(rules: dict[int, int], updates: list[list[int]]) -> int:
    def is_correct(pages: list[int]) -> bool:
        seen: set[int] = set([])
        for p in pages:
            if p in rules and rules[p] in seen:
                return False
        return True

    def median(pages: list[int]) -> int:
        assert(len(pages) % 2 != 0)
        return pages[int(len(pages) / 2)]

    return sum(median(pages) if is_correct(pages) else 0 for pages in updates)

def p2(rules: dict[int, int], updates: list[list[int]]) -> int:
    return 0


def read_updates(path: pathlib.Path) -> tuple[dict[int, int], list[list[int]]]:
    with open(path, "r") as f:
        rules: dict[int, int] = {}
        updates: list[list[int]] = []
        for line in f:
            line = line.strip()
            if not line:
                break
            a, b = [int(n.strip()) for n in line.split("|")][0:2]
            rules[a] = b
        for line in f:
            pages = [int(n.strip()) for n in line.split(",")]
            updates.append(pages)

        return rules, updates


def resolve_path(
    filename: str, dir: pathlib.Path = pathlib.Path(__file__).parent
) -> pathlib.Path:
    return dir.parent / filename


if __name__ == "__main__":
    main()
