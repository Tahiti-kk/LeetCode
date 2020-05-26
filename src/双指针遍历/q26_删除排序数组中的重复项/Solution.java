/**
 * lp用于改变数组，rp用于遍历数组
 * lp、rp在所指向的值上始终只差一个数字
 * 当相同时，rp向右移动
 * 当不相同时，lp右边一个数改为rp所指的值（nums[lp+1]与nums[rp]相同时，相当于重复操作；不同时相当于改变数组）
 */
public class Solution {
    public int removeDuplicates(int[] nums) {
        int length = nums.length;
        if(length <= 1) {
            return length;
        }
        int lp = 0, rp = 1;
        while(rp < length) {
            if(nums[lp]!=nums[rp]) {
                nums[++lp] = nums[rp++];
            } else {
                rp++;
            }
        }
        return lp+1;
    }
}