/**
 * dp:当前元素为尾的递增子序列的长度
 */
class Solution {
    public int lengthOfLIS(int[] nums) {
        int len = nums.length;
        if(len == 0) {
            return 0;
        }
        int[] dp = new int[len];

        dp[0] = 1;
        // 记录最大值
        int max = 1;
        for(int i = 1; i < len; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                // 如果大于num[j]并且不大于dp[j]，则说明可以往后增加一位
                if(nums[i] > nums[j] && dp[i] <= dp[j]) {
                    dp[i] = dp[j] + 1;
                }
            }
            max = Math.max(max, dp[i]);
        }
        return max;

    }
}