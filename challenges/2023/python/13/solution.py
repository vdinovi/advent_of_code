import sys

type input = list[str]


def solve_p1(lines: input) -> int:
    patterns = parse(lines)
    sum = 0
    for pattern in patterns:
        if len(pattern) < 2:
            continue
        for pivot in vertical_reflections(pattern):
            sum += pivot + 1
            break
        for pivot in horizontal_reflections(pattern):
            sum += 100 * (pivot + 1)
            break
    return sum


def solve_p2(lines: input) -> int:
    return 0


def parse(lines: str) -> list[list[str]]:
    patterns = [[]]
    for line in lines:
        if not line:
            patterns.append([])
        else:
            patterns[-1].append(line)
    return patterns


def vertical_reflections(pattern: list[str]) -> list[int]:
    cols = len(pattern[0])
    pivots = []
    for pivot in range(0, cols - 1):
        m = min(pivot, (cols - 1) - (pivot + 1))
        reflection = True
        for i in range(0, m + 1):
            for line in pattern:
                if line[pivot - i] != line[pivot + i + 1]:
                    reflection = False
                    break
            if not reflection:
                break
        if reflection:
            pivots.append(pivot)
    return pivots


def horizontal_reflections(pattern: list[str]) -> list[int]:
    cols = len(pattern[0])
    rows = len(pattern)
    pivots = []
    for pivot in range(0, rows - 1):
        m = min(pivot, (rows - 1) - (pivot + 1))
        reflection = True
        for i in range(0, m + 1):
            for col in range(0, cols):
                if pattern[pivot - i][col] != pattern[pivot + i + 1][col]:
                    reflection = False
                    break
            if not reflection:
                break
        if reflection:
            pivots.append(pivot)
    return pivots
