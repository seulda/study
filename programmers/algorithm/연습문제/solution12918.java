class Solution {
    public boolean solution(String s) {
        
        if(s.length() == 4 || s.length() == 6) {
            String[] arr = s.split("");
            for(int i = 0; i < s.length(); i++) {
                if( arr[i].equals("0") || arr[i].equals("1") || arr[i].equals("2") || arr[i].equals("3") 
                    || arr[i].equals("4") || arr[i].equals("5") || arr[i].equals("6") || arr[i].equals("7") 
                     || arr[i].equals("8") || arr[i].equals("9") ) {
                    continue;
                }
                return false;
            }
            return true;
        }
        return false;
    }
}
