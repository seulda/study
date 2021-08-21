class Solution {
    public String solution(int a, int b) {
        String answer = "";
        int[][] md = { {1, 31}, {2, 29}, {3, 31}, {4, 30}, {5, 31}, {6, 30}, {7, 31}, {8, 31}, {9, 30}, {10, 31}, {11, 30}, {12, 31} };
        String[] w = { "SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT" };
        int cnt = 4; // 2016 01 01 금요일 set
        
        for(int i = 0; i < a-1; i++) {
            cnt += md[i][1];
        }
        cnt += b;
        
        answer = w[cnt%7];
        
        return answer;
    }
}
