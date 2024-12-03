import math
import re
import pathlib


def main():
    print("p1(example1) =", p1(read_ops(resolve_path("example1.txt"))))
    print("P1(input)    =", p1(read_ops(resolve_path("input.txt"))))
    print("p2(example2) =", p2(read_ops(resolve_path("example2.txt"))))
    print("p2(input)    =", p2(read_ops(resolve_path("input.txt"))))


class Operation:
    __slots__ = ("name", "operands")

    def __init__(self, name: str, operands: list[int]):
        self.name = name
        self.operands = operands

    PATTERN = re.compile(r"(mul)\((\d+),(\d+)\)" + r"|(do)\(\)" + r"|(don\'t)\(\)")

    @classmethod
    def parse(cls, data: str) -> list["Operation"]:
        def extract(
            mul: str, left: str, right: str, do: str, dont: str, *_
        ) -> "Operation":
            if mul:
                return Mul(int(left), int(right))
            elif do:
                return Do()
            elif dont:
                return Dont()
            else:
                raise ValueError(mul, left, right, do, dont)

        matches = re.findall(cls.PATTERN, data)
        return [extract(*m) for m in matches]


class Mul(Operation):
    def __init__(self, left: int, right: int):
        super().__init__("mul", [left, right])


class Do(Operation):
    def __init__(self):
        super().__init__("do", [])


class Dont(Operation):
    def __init__(self):
        super().__init__("dont", [])


def p1(ops: list[Operation]) -> int:
    def execute(op: Operation) -> int:
        return math.prod(op.operands) if isinstance(op, Mul) else 0

    return sum(execute(op) for op in ops)


def p2(ops: list[Operation]) -> int:
    enabled = True

    def execute(op: Operation) -> int:
        nonlocal enabled
        if isinstance(op, Mul):
            return math.prod(op.operands) if enabled else 0
        elif isinstance(op, Do):
            enabled = True
        elif isinstance(op, Dont):
            enabled = False
        return 0

    return sum(execute(op) for op in ops)


def read_ops(path: pathlib.Path) -> list[Operation]:
    with open(path, "r") as f:
        return Operation.parse(f.read())


def resolve_path(
    filename: str, dir: pathlib.Path = pathlib.Path(__file__).parent
) -> pathlib.Path:
    return dir.parent / filename


if __name__ == "__main__":
    main()
