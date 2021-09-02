// 문자열 압축하기
// ex) input01 : a2b5a1c0 >> output : a3b5c0
// ex) input02 : a2a1a4 >> output : a7

class Main {
	public static void main(String[] args) throws Exception {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		String input = br.readLine();
		// 데이터 넣을 것
		String[][] arr = new String[26][2];
		int count = 0;
		int word = -1;
		int check = 0;
		// 데이터 등장 여부
		String[][] array = {{"a", "0"}, {"b", "0"}, {"c", "0"}, {"d", "0"}, {"e", "0"}, {"f", "0"}, 
                        {"g", "0"}, {"h", "0"}, {"i", "0"}, {"j", "0"}, {"k", "0"}, {"l", "0"}, 
                        {"m", "0"}, {"n", "0"}, {"o", "0"}, {"p", "0"}, {"q", "0"}, {"r", "0"}, 
                        {"s", "0"}, {"t", "0"}, {"u", "0"}, {"v", "0"}, {"w", "0"}, {"x", "0"}, {"y", "0"}, {"z", "0"}};
		String result = "";
		
		for(int k = 0; k < arr.length; k++) {
			arr[k][1] = "0";
		}
		
		for(int i = 0; i < input.length(); i++) {
			String tw = String.valueOf(input.charAt(i));
			if(tw.equals("0") == false && tw.equals("1") == false && tw.equals("2") == false && tw.equals("3") == false
				 && tw.equals("4") == false && tw.equals("5") == false && tw.equals("6") == false && tw.equals("7") == false
				 && tw.equals("8") == false && tw.equals("9") == false) {
				check = 0;
				// 숫자가 아니면 = 문자이면. 문자 검출
				if (count == 0) {
					arr[0][0] = tw;
					array[munja(tw)][1] = "1";
					word = count;
					count++;
				} else if (array[munja(tw)][1].equals("1")) {
					for(int w = 0; w < count; w++){
						if(arr[w][0].equals(tw)) {
							word = w; 
						}
					}
				} else {
					arr[count][0] = tw;
					array[munja(tw)][1] = "1";
					word = count;
					count++;
				}
			}
			
			if(tw.equals("0") || tw.equals("1") || tw.equals("2") || tw.equals("3") || tw.equals("4") || tw.equals("5")
				 || tw.equals("6") || tw.equals("7") || tw.equals("8") || tw.equals("9")) {
				//숫자이면. 숫자검출
				if(check>0) {
					String c = "";
					for(int x = check; x >= 0; x--) {
						c = c + String.valueOf(input.charAt(i-x));
					}
					arr[word][1] = c;
				} else {
					arr[word][1] = (Integer.parseInt(arr[word][1]) + Character.getNumericValue(input.charAt(i))) + "";
				}
				check++;
			}
			
		}
		
		for(int j = 0; j < count; j++) {
			result += arr[j][0] + arr[j][1];
		}
		
		System.out.println(result);
	}
	
	public static int munja(String a){
		if(a.equals("a")) { return 0; }
		else if(a.equals("b")) { return 1; }
		else if(a.equals("c")) { return 2; }
		else if(a.equals("d")) { return 3; }
		else if(a.equals("e")) { return 4; }
		else if(a.equals("f")) { return 5; }
		else if(a.equals("g")) { return 6; }
		else if(a.equals("h")) { return 7; }
		else if(a.equals("i")) { return 8; }
		else if(a.equals("j")) { return 9; }
		else if(a.equals("k")) { return 10; }
		else if(a.equals("l")) { return 11; }
		else if(a.equals("m")) { return 12; }
		else if(a.equals("n")) { return 13; }
		else if(a.equals("o")) { return 14; }
		else if(a.equals("p")) { return 15; }
		else if(a.equals("q")) { return 16; }
		else if(a.equals("r")) { return 17; }
		else if(a.equals("s")) { return 18; }
		else if(a.equals("t")) { return 19; }
		else if(a.equals("u")) { return 20; }
		else if(a.equals("v")) { return 21; }
		else if(a.equals("w")) { return 22; }
		else if(a.equals("x")) { return 23; }
		else if(a.equals("y")) { return 24; }
		else if(a.equals("z")) { return 25; }
		return -1;
	}
}
