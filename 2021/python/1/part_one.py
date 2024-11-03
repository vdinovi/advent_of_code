import sys

def depth_increase_count(readings):
    count = 0
    prev = None
    for index, value in enumerate(readings):
        if prev is not None:
            if value - prev > 0:
                count += 1
        prev = value
    return count


if __name__ == '__main__':
    readings = [int(line.rstrip()) for line in sys.stdin]
    count = depth_increase_count(readings)
    print(count)

