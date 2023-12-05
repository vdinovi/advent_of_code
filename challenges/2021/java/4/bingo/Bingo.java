package bingo;

import java.util.ArrayList;

public class Bingo {
    public static void main(String args[]) {
        Parser parser = new Parser(System.in);
        Solver solver = new Solver(parser.getRecord(), parser.getBoards());
        int solution = solver.solvePartTwo();
        System.out.println(solution);
    }
}
