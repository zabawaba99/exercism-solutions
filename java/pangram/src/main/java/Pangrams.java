import java.util.ArrayList;
import java.util.Arrays;

public class Pangrams {
	private static final ArrayList<String> ALL_THE_LETTERS = new ArrayList<>(Arrays.asList("a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
			"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" ));

	public static boolean isPangram(String s) {
		if (s == null) s = "";

		// normalize case
		final String candidate = s.toLowerCase();
        return ALL_THE_LETTERS.stream().allMatch((s1) -> candidate.contains(s1));
	}
}