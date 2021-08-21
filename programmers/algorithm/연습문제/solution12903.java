class Solution {
    public String solution(String s) {
        String answer = "";
        int cnt = s.length();
        if(cnt%2 == 1) {  // 홀수
            answer = s.substring(cnt/2, cnt/2+1);
        } else { // 짝수
            answer = s.substring(cnt/2-1, cnt/2+1);
        }
        return answer;
    }
}
