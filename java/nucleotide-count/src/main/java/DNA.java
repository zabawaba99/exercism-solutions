import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class DNA {
	@SuppressWarnings("serial")
	private static final ArrayList<Character> VALID_NUCLEOTIDES = new ArrayList<Character>() {
		{
			add('A');
			add('C');
			add('G');
			add('T');
		}
	};

	private String strain;

	public DNA(String strain) {
		this.strain = strain == null ? "" : strain;
	}

	/**
	 *
	 * @return A map containing all the nucleotides as keys and the number of
	 *         times they appear in the strain as the value.
	 */
	public Map<Character, Integer> nucleotideCounts() {
		@SuppressWarnings("serial")
		Map<Character, Integer> rtn = new HashMap<Character, Integer>() {
			{
				put('A', 0);
				put('C', 0);
				put('G', 0);
				put('T', 0);
			}
		};
		for (char nucleotide : strain.toCharArray()) {
			Integer count = rtn.get(nucleotide);
			rtn.put(nucleotide, count + 1);
		}
		return rtn;
	}

	/**
	 * Identifies how many times the given nucleotide occurs in the given strain.
	 *
	 * @param nucleotide - The nucleotide to search for.
	 * @throws IllegalArgumentException
	 *             if the given nucleotide is not valid symbol.
	 * @return The amount of times the nucleotide appeared in the strain.
	 */
	public int count(char nucleotide) {
		if (!VALID_NUCLEOTIDES.contains(nucleotide)) {
			throw new IllegalArgumentException("Nucleotide no bueno");
		}
		return nucleotideCounts().get(nucleotide);
	}
}
