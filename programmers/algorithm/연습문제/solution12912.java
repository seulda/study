class Solution {
    public long solution(int a, int b) {
        long answer = 0;
        
        if (a==b) return a;
        
        int x = a>b ? b : a;
        int y = a>b ? a : b;
        for(int i = x; i <= y; i++) {
            answer += i;
        }
        
        return answer;
    }
}
