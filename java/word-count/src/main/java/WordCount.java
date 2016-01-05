import java.util.HashMap;
import java.util.Map;

public class WordCount {

	public Map<String,Integer> phrase(String phrase) {
		Map<String,Integer> rtn = new HashMap<>();
		if (phrase == null) phrase = "";

		// normalize casing
		phrase = phrase.toLowerCase();

		// ignore punctuation
		phrase = phrase.replaceAll("(:|!|&|@|^|%|\\$|\\^|,)", "");

		for(String word : phrase.split(" ")) {
			if (word.isEmpty()) continue;

			Integer count = rtn.get(word);
			if (count == null) count = 0;
			rtn.put(word, count+1);
		}
		return rtn;
	}
}
