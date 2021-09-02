// 2021 카카오 채용연계형 인턴십 > 거리두기 확인하기
// 
// 코로나 바이러스 감염 예방을 위해 응시자들은 거리를 둬서 대기를 해야하는데 개발 직군 면접인 만큼 아래와 같은 규칙으로 대기실에 거리를 두고 앉도록 안내하고 있습니다.
// (1) 대기실은 5개이며, 각 대기실은 5x5 크기입니다.
// (2) 거리두기를 위하여 응시자들 끼리는 맨해튼 거리1가 2 이하로 앉지 말아 주세요.
// (3) 단 응시자가 앉아있는 자리 사이가 파티션으로 막혀 있을 경우에는 허용합니다.
// 5개의 대기실을 본 죠르디는 각 대기실에서 응시자들이 거리두기를 잘 기키고 있는지 알고 싶어졌습니다.
// 자리에 앉아있는 응시자들의 정보와 대기실 구조를 대기실별로 담은 2차원 문자열 배열 places가 매개변수로 주어집니다.
// 각 대기실별로 거리두기를 지키고 있으면 1을, 한 명이라도 지키지 않고 있으면 0을 배열에 담아 return 하도록 solution 함수를 완성해 주세요.
//   제한사항
// places의 행 길이(대기실 개수) = 5
// places의 각 행은 하나의 대기실 구조를 나타냅니다.
// places의 열 길이(대기실 세로 길이) = 5
// places의 원소는 P,O,X로 이루어진 문자열입니다.
// places 원소의 길이(대기실 가로 길이) = 5
// P는 응시자가 앉아있는 자리를 의미합니다.
// O는 빈 테이블을 의미합니다.
// X는 파티션을 의미합니다.
// 입력으로 주어지는 5개 대기실의 크기는 모두 5x5 입니다.
// return 값 형식
// 1차원 정수 배열에 5개의 원소를 담아서 return 합니다.
// places에 담겨 있는 5개 대기실의 순서대로, 거리두기 준수 여부를 차례대로 배열에 담습니다.
// 각 대기실 별로 모든 응시자가 거리두기를 지키고 있으면 1을, 한 명이라도 지키지 않고 있으면 0을 담습니다.


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
