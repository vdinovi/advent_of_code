import sys
import numpy as np

def parse():
    return np.asarray([[int(ch) for ch in line.rstrip()] for line in sys.stdin])

def adjacent(x, y, w, h):
    adj = []
    for y1 in range(y - 1, y + 2):
        if y1 < 0 or y1 >= h:
            continue
        for x1 in range(x - 1, x + 2):
            if x1 < 0 or x1 >= w or (y == y1 and x == x1):
                continue
            adj.append([x1, y1])
    return adj

def process(grid, x, y, width, height, flashed):
    if (x, y) not in flashed:
        level = grid[y][x]
        if level == 9:
            flashed[(x, y)] = True
            grid[y][x] = 0
            for point in adjacent(x, y, width, height):
                process(grid, point[0], point[1], width, height, flashed)
        else:
            grid[y][x] += 1
    else:
        level = grid[y][x]


def step(grid, n):
    width = len(grid)
    height = len(grid[0])
    flashed = {}
    for y in range(0, height):
        for x in range(0, width):
            process(grid, x, y, width, height, flashed)
    return len(flashed)



def simulate(grid, steps):
    return sum([step(grid, n) for n in range(1, steps + 1)])

def simulate_until(grid, fn):
    n = 0
    while True:
        n += 1
        if fn(step(grid, n)):
            return n

if __name__ == "__main__":
    grid = parse()
    size = sum([len(r) for r in grid])
    step = simulate_until(grid, lambda x: x == size)
    print(step)

