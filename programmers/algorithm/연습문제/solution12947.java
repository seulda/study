class Solution {
    public boolean solution(int x) {
        boolean answer = false;
        String str = String.valueOf(x);
        int cnt = 0;
        
        for(int i = 0; i < str.length(); i++) {
            cnt += Integer.parseInt(str.substring(i, i+1));
        }
        
        if (x%cnt == 0) {
            answer = true;
        }
        
        return answer;
    }
}
