public class Pangrams {
	private static final String[] ALL_THE_LETTERS = new String[] { "a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
			"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z" };

	public static boolean isPangram(String s) {
		if (s == null)
			s = "";

		// normalize case
		s = s.toLowerCase();
		for (int i = 0; i < ALL_THE_LETTERS.length; i++) {
			if (!s.contains(ALL_THE_LETTERS[i]))
				return false;
		}
		return true;
	}
}