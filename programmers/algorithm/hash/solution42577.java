import java.util.*;

class Solution {
    public boolean solution(String[] phone_book) {
        Arrays.sort(phone_book);
        
        for (int i = 0; i < phone_book.length-1; i++) {
            if (phone_book[i].charAt(0) != phone_book[i+1].charAt(0))
                continue;
            else if (phone_book[i].length() >= phone_book[i+1].length())
                continue;
        	for (int j = i+1; j < phone_book.length; j++) {
                if (phone_book[i].charAt(0) != phone_book[j].charAt(0))
                    break;
                else if (phone_book[i].length() >= phone_book[j].length())
                    break;
                else if (phone_book[j].startsWith(phone_book[i])) {
                    return false;
                }
        	}
        }
        
        return true;
    }
}
