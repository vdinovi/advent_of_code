package bingo;

import java.util.ArrayList;

public class Solver {
    private Record record;
    private ArrayList<Board> boards;

    public Solver(Record record, ArrayList<Board> boards) {
        this.record = record;
        this.boards = boards;
    }

    public int solvePartOne() {
        int value = -1;
        Board winner = null;

        while (!record.isDone() && winner == null) {
            value = record.next();
            for (Board board : boards) {
                board.mark(value);
                if (board.isSolved()) {
                    winner = board;
                    break;
                }
            }
        }
        assert(winner != null);
        return value * winner.sumOfUnmarkedNumbers();
    }

    public int solvePartTwo() {
        int value = -1;
        ArrayList<Board> activeBoards = new ArrayList<Board>(boards);
        Board loser = null;

        while (!record.isDone() && activeBoards.size() > 0) {
            value = record.next();
            for (Board board : boards) {
                if (!board.isSolved()) {
                    board.mark(value);
                    if (board.check()) {
                        activeBoards.remove(board);
                        if (activeBoards.isEmpty()) {
                            loser = board;
                            break;
                        }
                    }
                }
            }
        }
        assert(loser != null);
        return value * (loser.sumOfUnmarkedNumbers());
    }

}

