type input = list[str]


def solve_p1(lines: input) -> int:
    patterns = parse(lines)
    sum = 0
    for pattern in patterns:
        if len(pattern) < 2:
            continue
        if (pivot := vertical_reflection(pattern, 0)) is not None:
            sum += pivot + 1
        elif (pivot := horizontal_reflection(pattern, 0)) is not None:
            sum += 100 * (pivot + 1)
    return sum


def solve_p2(lines: input) -> int:
    patterns = parse(lines)
    sum = 0
    for pattern in patterns:
        if len(pattern) < 2:
            continue
        if (pivot := vertical_reflection(pattern, 1)) is not None:
            sum += pivot + 1
        elif (pivot := horizontal_reflection(pattern, 1)) is not None:
            sum += 100 * (pivot + 1)
    return sum


def parse(lines: list[str]) -> list[list[str]]:
    patterns = [[]]
    for line in lines:
        if not line:
            patterns.append([])
        else:
            patterns[-1].append(line)
    return patterns


def vertical_reflection(pattern: list[str], smudges: int) -> int | None:
    cols = len(pattern[0])
    rows = len(pattern)
    for pivot in range(0, cols - 1):
        diff = 0
        width = min(pivot, (cols - 1) - (pivot + 1))
        for i in range(0, width + 1):
            for row in range(0, rows):
                if pattern[row][pivot - i] != pattern[row][pivot + i + 1]:
                    diff += 1
        if diff == smudges:
            return pivot
    return None


def horizontal_reflection(pattern: list[str], smudges: int) -> int | None:
    cols = len(pattern[0])
    rows = len(pattern)
    for pivot in range(0, rows - 1):
        diff = 0
        width = min(pivot, (rows - 1) - (pivot + 1))
        for i in range(0, width + 1):
            for col in range(0, cols):
                if pattern[pivot - i][col] != pattern[pivot + i + 1][col]:
                    diff += 1
        if diff == smudges:
            return pivot
    return None
