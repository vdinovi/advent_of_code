package bingo;

import java.util.ArrayList;
import java.lang.StringBuilder;

public class Record {
    private int[] numbers;
    private int position;

    public Record(ArrayList<Integer> numbers) {
        this(numbers.stream().mapToInt(i -> i).toArray());
    }

    public Record(int[] numbers) {
        this.numbers = new int[numbers.length];
        for (int i = 0; i < numbers.length; i++) {
            this.numbers[i] = numbers[i];
        }
        position = 0;
    }

    public int next() {
        return numbers[position++];
    }

    public boolean isDone() {
        return position == numbers.length;
    }

    public String toString() {
        StringBuilder builder = new StringBuilder();
        for (int i = 0; i < numbers.length; ++i) {
            builder.append(numbers[i]);
            if (i < numbers.length - 1) {
                builder.append(" ");
            }
        }
        return builder.toString();
    }
}
