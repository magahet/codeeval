import java.io.File;
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class LongestWord {
    public static void main (String[] args) {
        File file = new File(args[0]);
        try {
            BufferedReader in = new BufferedReader(new FileReader(file));
            String line;
            while ((line = in.readLine()) != null) {
                String[] lineArray = line.split(" ");
                if (lineArray.length > 0) {
                    String longestWord = "";
                    for (int i = 0; i < lineArray.length; i++) {
                        String word = lineArray[i];
                        if (word.length() > longestWord.length()) {
                            longestWord = word;
                        }
                    }
                    System.out.println(longestWord);
                }
            }
        } catch (IOException e) {
            System.out.println(e);
        }
    }
}
