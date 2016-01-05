import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Etl {
   public Map<String, Integer> transform(Map<Integer, List<String>> old) {
      Map<String, Integer> transformed = new HashMap<>();

      for(Integer pointVal : old.keySet()) {
         List<String> letters = old.get(pointVal);

         for (String letter : letters) {
            transformed.put(letter.toLowerCase(), pointVal);
         }
      }
      return transformed;
   }
}
