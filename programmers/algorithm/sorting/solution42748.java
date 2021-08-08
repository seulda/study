import java.util.*;

class Ascending implements Comparator<Integer>{
	public int compare(Integer a, Integer b)
	{
		return a.compareTo(b);
	}
}

class Solution {
    public int[] solution(int[] array, int[][] commands) {
        int[] answer = new int[commands.length];
        
        ArrayList<Integer> temp = new ArrayList<>();
        ArrayList<Integer> list = new ArrayList<>();
        
        for (int i = 0; i < commands.length; ++i) {
            for (int j = commands[i][0]-1; j < commands[i][1]; ++j) {
                temp.add(array[j]);
            }
            Collections.sort(temp, new Ascending());
            list.add(temp.get(commands[i][2]-1));
            temp.clear();
        }
        
        for (int i = 0; i<list.size(); ++i) {
            answer[i] = list.get(i);
        }
        
        return answer;
    }
}
