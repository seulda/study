// 2021 카카오 채용연계형 인턴십 > 숫자 문자열과 영단어
//
// 다음은 숫자의 일부 자릿수를 영단어로 바꾸는 예시입니다.
// 1478 → "one4seveneight"
// 234567 → "23four5six7"
// 10203 → "1zerotwozero3"
// 이렇게 숫자의 일부 자릿수가 영단어로 바뀌어졌거나, 혹은 바뀌지 않고 그대로인 문자열 s가 매개변수로 주어집니다. s가 의미하는 원래 숫자를 return 하도록 solution 함수를 완성해주세요.
// 제한사항
// 1 ≤ s의 길이 ≤ 50
// s가 "zero" 또는 "0"으로 시작하는 경우는 주어지지 않습니다.
// return 값이 1 이상 2,000,000,000 이하의 정수가 되는 올바른 입력만 s로 주어집니다.
  
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
