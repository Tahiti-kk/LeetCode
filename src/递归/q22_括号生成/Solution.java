public class Solution {
    /**
     * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
     * 
     * 示例：
     * 输入：n = 3
     * 输出：[
     *        "((()))",
     *        "(()())",
     *        "(())()",
     *        "()(())",
     *        "()()()"
     *      ]
     */
    public List<String> generateParenthesis(int n) {
        List<String> list = new ArrayList<>();
        generate(list, "", 0, 0, n);
        return list;
    }

    /**
     * 通过递归生成
     * left为左括号个数，right为右括号个数
     */
    public void generate(List<String> list, String str, int left, int right, int n){
        // 进行剪枝
        if(left<right || left>n || right>n) {
            return;
        }
        // 如果最终字符串符合要求，则将str加入list中
        if(left==n && right==n) {
            list.add(str);
            return;
        }
        generate(list, str+"(", left+1, right, n);
        generate(list, str+")", left, right+1, n);
    }
}