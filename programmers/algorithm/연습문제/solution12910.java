import java.util.Arrays;

class Solution {
    public int[] solution(int[] arr, int divisor) {
        int[] answer = new int[arr.length];
        int[] array = null;
        int cnt = 0;
        
        for(int i = 0; i < arr.length; i++) {
            if ((arr[i] % divisor) == 0) {
                answer[cnt] = arr[i];
                cnt++;
            }
        }
        
        if (cnt == 0) {
            array = new int[1];
            array[0] = -1;
        } else {
            array = Arrays.copyOfRange(answer, 0, cnt);
            Arrays.sort(array);
        }
        
        return array;
    }
}
