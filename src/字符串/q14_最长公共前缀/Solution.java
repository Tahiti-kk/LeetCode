/**
 * 最长公共前缀
 * 将第一个字符串作为基准进行判别startWith()，当不符合条件后长度减一
 */
public class Solution {

    public String longestCommonPrefix(String[] strs) {
        if(strs.length == 0) {
            return "";
        }
        String ans = strs[0];
        for (String str : strs) {
            while (!str.startsWith(ans)) {
                if (ans.length() == 1) {
                    return "";
                }
                ans = ans.substring(0, ans.length() - 1);
            }
        }
        return ans;
    }

}