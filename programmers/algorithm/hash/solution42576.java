import java.util.*;

class Solution {
    public String solution(String[] participant, String[] completion) {
        String answer = "";
        
        HashMap<String, Integer> hm = new HashMap<>();

		for (String com : completion){
			Integer value = hm.get(com);
			if (value == null){
				hm.put(com, 1);
				continue;
			}
			hm.put(com, ++value);
		}
				
		for (String part : participant){
			Integer value = hm.get(part);
			if (value == null) {
				answer = part;
				break;
			} else if (value == 1) {
				hm.remove(part);
			} else {
				hm.put(part, value - 1);
			}
		}
        
        return answer;
    }
}
