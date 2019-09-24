import java.net.URLEncoder;
import java.util.HashMap;
import java.util.PriorityQueue;

/*
    Huffman Algorithm

    - https://en.wikipedia.org/wiki/Huffman_coding
    - https://www.geeksforgeeks.org/huffman-decoding/

    문자열의 빈도 역순으로 정렬하여
    가장 빈도가 높은 문자열은 짧은 값, 낮은 문자열을 상대적으로 긴 값을 주어 압축하는 알고리즘

    tree를 구현하여 encode/decode 하는데 사용한다.


 */


public class Main {

    static HashMap<String, String> encodeMap = new HashMap<>();

    public static void main(String[] args){

        String csets = "+%0123456789ABCDEF";    //URL Encode 결과 나올 수 있는 선택지 모음
        int []freq = new int[18];               //빈도를 저장할 배열

        String origin ="허프만 알고리즘 테스트";
//        String origin ="a";                   // 2개 이상의 단어만 고려하도록 하자.
        String input="";

        try{
            input = URLEncoder.encode(origin,"UTF-8");  // 한글 등 편하게 처리하기 위해 URL Encode를 진행.
        }catch(Exception e){
            e.printStackTrace();
        }

        char[] byteArr = input.toCharArray();   // String char 배열로 변환하여 빈도 계산

        for(int i=0; i<byteArr.length; i++) {
            freq[csets.indexOf(byteArr[i])]++;
        }

        // pq를 이용해 빈도가 낮은것을 골라내고,
        // 빈도가 낮은 두개의 노드를 합쳐서 한개로 만들고,
        // 다시 pq에 담아가며 전체 노드를 구성한다.
        System.out.println("== Freq Map ==");
        PriorityQueue<Node> pq = new PriorityQueue<>();
        for(int i=0; i<17; i++){
            if(freq[i]!=0) {
                System.out.print(csets.charAt(i)+" :"+freq[i]+"\t");
                pq.add(new Node(null, null, Character.toString(csets.charAt(i)), freq[i]));
            }
        }

        while(pq.size()>=2){
            Node n1 = pq.poll();
            Node n2 = pq.poll();
            pq.add(new Node(n1, n2, n1.value+n2.value, n1.freq+n2.freq));
        }

        // 마지막 pq에는 전체 노드의 모양이 남는다
        Node root  = pq.poll();

        // root 노드를 순회하며 encodeMap을 작성한다 ("A" : "00", "B" : "01" ... )
        traverse(root,"");

        System.out.println("\n\nINPUT TEXT : "+origin);
        System.out.println("URL ENCODED TEXT : "+input);
        System.out.println("HUFFMAN ENCODE RESULT : "+encode(input));
        System.out.println("HUFFMAN DECODE RESULT : "+decode(root,encode(input)));
        System.out.println("\nINPUT equals OUTPUT : "+input.equals(decode(root,encode(input))));
    }

    // Huffman Decode Function
    public static String decode(Node root, String input){
        StringBuilder output = new StringBuilder();
        Node tmp = root;
        for(int i=0; i<input.length(); i++){
            if("1".equals(Character.toString(input.charAt(i)))){
                tmp = tmp.left; //왼쪽으로 이동할땐 1을 붙인다.
                if(tmp.left == null && tmp.right == null) {
                    output.append(tmp.value);
                    tmp = root;
                }
            }else{
                tmp = tmp.right; //오른쪽으로 이동할땐 0을 붙인다.
                if(tmp.left == null && tmp.right == null) {
                    output.append(tmp.value);
                    tmp = root;
                }
            }
        }
        return output.toString();
    }

    //Huffman Encode Function
    public static String encode(String input){
        StringBuilder output = new StringBuilder();
        for(int i=0; i<input.length(); i++){
            // 글자를 만나면 encodeMap 에서 값을 찾아 치환한다.
            output.append(encodeMap.get(Character.toString(input.charAt(i))));
        }
        return output.toString();
    }

    // root Node 순회 function
    public static void traverse(Node node, String value){
        if(node.left == null && node.right == null) {
            encodeMap.put(node.value,value);
            return;
        }
        //왼쪽으로 갈땐 1, 오른쪽으로 갈땐 0을 붙여가며 null을 만날때까지 순회한다
        traverse(node.left,value+"1");
        traverse(node.right,value+"0");
    }

    // Node는 왼쪽Node, 오른쪽Node, 값 ("A", "B", "C" 등..), 빈도를 갖는 Class이다.
    public static class Node implements Comparable<Node>{
        Node(Node left, Node right, String value, int freq){
            this.left = left;
            this.right = right;
            this.value = value;
            this.freq = freq;
        }
        Node left;
        Node right;
        String value;
        int freq;

        @Override
        public String toString(){
            return this.value+" , "+this.freq;
        }

        @Override
        public int compareTo(Node o) {
            return -(o.freq-this.freq);
        }
    }
}
