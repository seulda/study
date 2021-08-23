class Solution {
    public String[] solution(String[] strings, int n) {
        String[] arr = strings;
        for(int i = 0; i < strings.length; i++) {
            for(int j = 0; j < strings.length; j++) {
                int test = String.valueOf(arr[i].charAt(n)).compareTo(String.valueOf(arr[j].charAt(n)));
                if (test < 0) {
                    String temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
                else if (test == 0 && arr[i].compareTo(arr[j]) < 0) {
                    String temp = arr[i];
                    arr[i] = arr[j];
                    arr[j] = temp;
                }
            }
        }
        
        return arr;
    }
}
