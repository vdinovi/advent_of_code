package bingo;

import java.util.ArrayList;

public class Board {
    // Origin is top-left, positive y is down, positive x is right
    private int[][] grid;
    private boolean[][] marked;
    private int height;
    private int width;
    private boolean solved;

    public Board(int height, int width, ArrayList<Integer> numbers) {
        this(height, width, numbers.stream().mapToInt(i -> i).toArray());
    }

    public Board(int height, int width, int[] numbers) {
        this.height = height;
        this.width = width;
        this.solved = false;
        grid = new int[height][width];
        marked = new boolean[height][width];
        for (int y = 0; y < height; y++) {
            for (int x = 0; x < width; x++) {
                grid[y][x] = numbers[y * width + x];
                marked[y][x] = false;
            }
        }
    }

    public void draw() {
        for (int y = 0; y < height; y++) {
            for (int x = 0; x < width; x++) {
                if (marked[y][x]) {
                    System.out.printf("|%2d|", grid[y][x]);
                } else {
                    System.out.printf(" %2d ", grid[y][x]);
                }
            }
            System.out.printf("\n");
        }
    }

    public boolean mark(int value) {
        for (int y = 0; y < height; y++) {
            for (int x = 0; x < width; x++) {
                if (grid[y][x] == value) {
                    return marked[y][x] = true;
                }
            }
        }
        return false;
    }

    public boolean check() {
        if (solved) {
            return true;
        }
        for (int y = 0; y < height; y++) {
            int x = 0;
            for (x = 0; x < width; x++) {
                if (!marked[y][x]) {
                    break;
                }
            }
            if (x == width) {
                return solved = true;
            }
        }
        for (int x = 0; x < width; x++) {
            int y = 0;
            for (y = 0; y < height; y++) {
                if (!marked[y][x]) {
                    break;
                }
            }
            if (y == height) {
                return solved = true;
            }
        }
        return false;
    }

    public boolean isSolved() {
        return solved;
    }

    public int sumOfUnmarkedNumbers() {
        int sum = 0;
        for (int y = 0; y < height; y++) {
            for (int x = 0; x < width; x++) {
                if (!marked[y][x]) {
                    sum += grid[y][x];
                }
            }
        }
        return sum;
    }
}
