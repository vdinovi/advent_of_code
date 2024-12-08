import pathlib
import typing
import functools
from collections import defaultdict, OrderedDict


def main():
    print("p1(example) =", p1(*read_updates(resolve_path("example.txt"))))
    print("P1(input)   =", p1(*read_updates(resolve_path("input.txt"))))
    print("p2(example) =", p2(*read_updates(resolve_path("example.txt"))))
    print("p2(input)   =", p2(*read_updates(resolve_path("input.txt"))))


def p1(rules: dict[int, set[int]], updates: list[list[int]]) -> int:
    return sum(
        median(pages) if is_correct(rules, iter(pages)) else 0 for pages in updates
    )


def p2(rules: dict[int, set[int]], updates: list[list[int]]) -> int:
    return sum(
        median(reorder(rules, pages))
        for pages in (pages for pages in updates if not is_correct(rules, iter(pages)))
    )


def reorder(rules: dict[int, set[int]], pages: list[int]) -> list[int]:
    P = set(pages)
    result: dict[int, bool] = OrderedDict()

    @functools.cache
    def insert(n: int):
        rs = rules[n] & P
        for r in rs.difference(result):
            insert(r)
        if rs.issubset(result):
            result[n] = True

    for p in pages:
        insert(p)

    return list(reversed(result.keys()))


def is_correct(rules: dict[int, set[int]], pages: typing.Iterator[int]) -> bool:
    seen: set[int] = set([])
    for p in pages:
        if seen.intersection(rules[p]):
            return False
        seen.add(p)
    return True


def median(pages: list[int]) -> int:
    assert len(pages) % 2 != 0
    return pages[int(len(pages) / 2)]


def read_updates(path: pathlib.Path) -> tuple[dict[int, set[int]], list[list[int]]]:
    with open(path, "r") as f:
        rules: dict[int, set[int]] = defaultdict(set)
        updates: list[list[int]] = []
        for line in f:
            line = line.strip()
            if not line:
                break
            a, b = [int(n.strip()) for n in line.split("|")][0:2]
            rules[a].add(b)
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
