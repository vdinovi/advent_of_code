import typing
import pathlib
import numpy as np

if typing.TYPE_CHECKING:
    from numpy.typing import NDArray


def main():
    print("p1(example) =", p1(read_map(resolve_path("example.txt"))))
    print("P1(input)   =", p1(read_map(resolve_path("input.txt"))))
    print("p2(example) =", p2(read_map(resolve_path("example.txt"))))
    print("p2(input)   =", p2(read_map(resolve_path("input.txt"))))


def p1(map: np.ndarray) -> int:
    kernel = np.array([ord(ch) for ch in "XMAS"])
    maps = [
        map,
        np.rot90(map, k=1),
        np.rot90(map, k=2),
        np.rot90(map, k=3),
    ]
    return scan_line(kernel, maps) + scan_diagonal(kernel, maps)


def p2(map: np.ndarray) -> int:
    kernel = np.array([ord(ch) for ch in "MAS"])
    maps = [
        map,
        np.rot90(map, k=1),
        np.rot90(map, k=2),
        np.rot90(map, k=3),
    ]
    return scan_x(kernel, maps)


def scan_line(
    kernel: "NDArray[np.int_]", maps: typing.Iterable["NDArray[np.int_]"]
) -> int:
    assert kernel.ndim == 1
    K = len(kernel)
    count = 0
    for map in maps:
        rows, cols = map.shape
        for row in range(rows):
            for col in range(cols):
                section = map[row][col : col + K]
                if np.array_equal(section, kernel):
                    count += 1
    return count


def scan_diagonal(
    kernel: "NDArray[np.int_]", maps: typing.Iterable["NDArray[np.int_]"]
) -> int:
    assert kernel.ndim == 1
    K = len(kernel)
    count = 0
    for map in maps:
        _, cols = map.shape
        for diag in [np.diagonal(map, offset=n) for n in range(-cols, cols)]:
            for col in range(len(diag)):
                section = diag[col : col + K]
                if np.array_equal(section, kernel):
                    count += 1
    return count


def scan_x(
    kernel: "NDArray[np.int_]", maps: typing.Iterable["NDArray[np.int_]"]
) -> int:
    assert kernel.ndim == 1
    K = len(kernel)
    count = 0
    for map in maps:
        rows, cols = map.shape
        for row in range(rows - K + 1):
            for col in range(cols - K + 1):
                section = map[row : row + K, col : col + K]
                d1 = section.diagonal()
                d2 = np.fliplr(section).diagonal()
                if np.array_equal(d1, kernel) and np.array_equal(d2, kernel):
                    count += 1
    return count


def read_map(path: pathlib.Path) -> np.ndarray:
    with open(path, "r") as f:
        return np.array([[ord(ch) for ch in line.strip()] for line in f.readlines()])


def resolve_path(
    filename: str, dir: pathlib.Path = pathlib.Path(__file__).parent
) -> pathlib.Path:
    return dir.parent / filename


if __name__ == "__main__":
    main()
