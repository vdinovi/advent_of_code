import enum
import copy

Input = list[str]


def solve_p1(lines: Input) -> int:
    patterns = parse(lines)
    if (n := len(patterns)) != 1:
        raise ValueError(f"Expected only one pattern, got {n}")
    pattern = patterns[0]
    return weight(pattern)


def solve_p2(lines: Input) -> int:
    return 0


Pattern = list[list[str]]


def tilt(pattern: Pattern) -> Pattern:
    tilted = copy.deepcopy(pattern)
    for row in range(len(tilted) - 1, 0, -1):
        for col in range(len(tilted[row])):
            match tilted[row][col]:
                case Rock.CUBE.value | Rock.NONE.value:
                    pass
                case Rock.ROUND.value:
                    # TODO: finish, need to push multiple spheres
                    pass
                case v:
                    raise ValueError(f"Strange rock '{v}'")
    return tilted


class Rock(enum.Enum):
    CUBE = "#"
    ROUND = "O"
    NONE = "."


def weight(pattern: Pattern) -> int:
    w = 0
    for distance, tiles in enumerate(reversed(pattern)):
        for tile in tiles:
            match tile:
                case Rock.CUBE.value | Rock.NONE.value:
                    pass
                case Rock.ROUND.value:
                    w += distance + 1
                case v:
                    raise ValueError(f"Strange rock '{v}'")
    return w


def parse(lines: Input) -> list[Pattern]:
    patterns = [[]]
    for line in lines:
        if not line:
            patterns.append([])
        else:
            patterns[-1].append(line)
    return patterns
