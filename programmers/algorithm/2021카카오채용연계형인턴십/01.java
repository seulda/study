class Solution {
    public int solution(String s) {
        int answer = 0;
        String text;
        
        text = s.replace("one", "1");
        text = text.replace("two", "2");
        text = text.replace("three", "3");
        text = text.replace("four", "4");
        text = text.replace("five", "5");
        text = text.replace("six", "6");
        text = text.replace("seven", "7");
        text = text.replace("eight", "8");
        text = text.replace("nine", "9");
        text = text.replace("zero", "0");
        
        answer = Integer.parseInt(text);
        return answer;
    }
}
