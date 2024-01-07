import enum
import copy
from typing import Any

Input = list[str]


def solve_p1(lines: Input) -> int:
    dish = Dish.parse(lines)
    return dish.weight()


def solve_p2(lines: Input) -> int:
    return 0


class Rock(enum.Enum):
    CUBE = "#"
    ROUND = "O"
    NONE = "."


class Dish:
    def __init__(self, rocks: list[list[Rock]]) -> None:
        self._rocks = rocks

    def tilt(self) -> "Dish":
        tilted = copy.deepcopy(self)
        for row in range(len(tilted._rocks) - 1, 0, -1):
            for col in range(len(tilted._rocks[row])):
                match tilted._rocks[row][col]:
                    case Rock.CUBE.value | Rock.NONE.value:
                        pass
                    case Rock.ROUND.value:
                        # TODO: finish, need to push multiple spheres
                        pass
                    case v:
                        raise ValueError(f"Strange rock '{v}'")
        return tilted

    def weight(self) -> int:
        w = 0
        for row in range(len(self._rocks) - 1, 0, -1):
            for col in range(len(self._rocks[row])):
                if self._rocks[row][col] == Rock.ROUND:
                    w += row + 1
        return w

    @classmethod
    def parse(cls, lines: Input) -> "Dish":
        rocks = [[Rock(c) for c in line] for line in lines]
        return cls(rocks)
