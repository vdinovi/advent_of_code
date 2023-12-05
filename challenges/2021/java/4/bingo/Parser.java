package bingo;

import java.util.ArrayList;
import java.util.Scanner;
import java.io.InputStream;
import java.util.regex.Pattern;

public class Parser {
    private Record record;
    private ArrayList<Board> boards;

    public Parser(InputStream instream) {
        Scanner scanner = new Scanner(instream);

        String line = scanner.nextLine();
        record = new Record(parseIntegers(line, Pattern.compile(",")));
        scanner.nextLine();

        boards = new ArrayList<Board>();
        while (scanner.hasNextLine()) {
            StringBuilder builder = new StringBuilder();
            while (scanner.hasNextLine() && !(line = scanner.nextLine()).isEmpty()) {
                line.replace("\n", " ").replace("\r", " ");
                builder.append(line);
                builder.append(" ");
            }
            ArrayList<Integer> numbers = parseIntegers(builder.toString(), Pattern.compile("\\s+"));
            Board board = new Board(5, 5, numbers);
            boards.add(board);
        }
    }

    public Record getRecord() {
        return record;
    }

    public ArrayList<Board> getBoards() {
        return boards;
    }

    private ArrayList<Integer> parseIntegers(String input, Pattern delim) {
        ArrayList<Integer> numbers = new ArrayList<Integer>();
        Scanner scanner = new Scanner(input);
        scanner.useDelimiter(delim);
        while (scanner.hasNextInt()) {
            numbers.add(scanner.nextInt());
        }
        scanner.close();
        return numbers;
    }
}
