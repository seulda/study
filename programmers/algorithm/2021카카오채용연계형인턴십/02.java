class Solution {
    public int[] solution(String[][] places) {
        int[] answer = new int [5];
        String[][] temp = new String[5][5];

        for (int i=0; i<5; ++i) {
            int pCount = 0;
            int[] xp = new int[25];
            int[] yp = new int[25];

            for (int j=0; j<5; ++j) {
                for (int k=0; k<5; ++k) {
                    temp[j][k] = String.valueOf(places[i][j].charAt(k));
                    if(temp[j][k].equals("P")){ 
                        xp[pCount] = j;
                        yp[pCount] = k;
                        pCount += 1;
                    }
                }
            }

            if(pCount == 1 || pCount == 0) { 
                answer[i] = 1;
                continue;
            } else if(pCount > 13) {  
                answer[i] = 0;
                continue;
            }

            for (int s=0; s<pCount; ++s) {
                for (int t=s; t<pCount; ++t) { 
                    if (t==s) 
                        continue;
                    if( (Math.abs(xp[s]-xp[t]) + Math.abs(yp[s]-yp[t])) > 2 ) {
                        answer[i] = 1;
                        continue;
                    } else if( (Math.abs(xp[s]-xp[t]) + Math.abs(yp[s]-yp[t])) <= 2 ) { 
                        if( temp[( Math.abs(xp[s]-1) )][( yp[s] )].equals("P") || temp[( 4-Math.abs(3-xp[s]) )][( yp[s] )].equals("P")
                                || temp[( xp[s] )][( Math.abs(yp[s]-1) )].equals("P") || temp[( xp[s] )][( 4-Math.abs(3-yp[s]) )].equals("P") )  {
                            answer[i] = 0;
                            break;
                        } 
                        else if( temp[( Math.abs(xp[t]-1) )][( yp[t] )].equals("P") || temp[( 4-Math.abs(3-xp[t]) )][( yp[t] )].equals("P")
                                || temp[( xp[t] )][( Math.abs(yp[t]-1) )].equals("P") || temp[( xp[t] )][( 4-Math.abs(3-yp[t]) )].equals("P") ) {
                            answer[i] = 0;
                            break;
                        }
                        else if( ( (((xp[s]-xp[t])==0) && (Math.abs(yp[s]-yp[t])==2))  &&  (temp[xp[s]][(yp[s]+yp[t])/2].equals("O")) )     ) {
                            answer[i] = 0;
                            break;
                        } 
                        else if( (((yp[s]-yp[t])==0) && (Math.abs(xp[s]-xp[t])==2))  &&  (temp[(xp[s]+xp[t])/2][yp[t]].equals("O")) ) {
                            answer[i] = 0;
                            break;
                        }
                        else if( (Math.abs(xp[s]-xp[t])==1) && (Math.abs(yp[s]-yp[t])==1) ) {
                            if( (xp[s]-xp[t]) < 0 ) {
                                if( (temp[xp[t]][yp[s]].equals("O")) || (temp[xp[s]][yp[t]].equals("O")) ) {
                                    answer[i] = 0;
                                    break;
                                }
                            }
                        }
                        else {
                            answer[i] = 1;
                        }
                    }
                }
                if (answer[i] == 0)
                    break;
            }
        }
        return answer;
    }
}
