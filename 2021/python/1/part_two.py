import sys
import pdb

class DepthGauge:
    def __init__(self, instream, width = 3):
        self.readings = []
        self.width = width
        self.buffers  = [[] for _ in range(width)]
        self.steps  = 0
        self.input = iter(instream.readline, b'')

    def __next__(self):
        self.steps += 1
        reading = next(self.input)
        if not reading:
            for _ in range(len(self.buffers)):
                self.readings.append(self.buffers.pop(0))
            raise StopIteration
        depth = int(reading.rstrip())
        times = min(self.steps, self.width)
        for index in range(times):
            self.buffers[index].append(depth)
        if len(self.buffers[0]) == self.width:
            self.readings.append(self.buffers.pop(0))
            self.buffers.append([])

    def __iter__(self):
        return self

    def monitor(self):
        [_ for _ in self]

    def depth_increases(self):
        prev = None
        count = 0
        for depths in self.readings:
            depth = sum(depths)
            if prev and depth - prev > 0:
                count += 1
            prev = depth
        return count


if __name__ == '__main__':
    gauge = DepthGauge(sys.stdin)
    gauge.monitor()
    print(gauge.depth_increases())

