import java.util.*;

public class Solution {
    public int[] solution(int []arr) {
        int[] answer = new int[arr.length];
        int set = -1;
        int cnt = 0;
        
        for(int i = 0; i < arr.length; i++) {
            if(set == arr[i]) 
                continue;
            set = arr[i];
            answer[cnt] = set;
            cnt++;
        }
        
        int[] array = Arrays.copyOfRange(answer, 0, cnt);

        return array;
    }
}
