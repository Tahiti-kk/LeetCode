class Solution {
    public int maxProfit(int[] prices) {
        int result = 0;
        if(prices.length <= 1){
            return result;
        }
        int minPrice = prices[0];
        for(int i = 1; i < prices.length; i++) {
            // 左指针记录第i天前的股票最低值
            minPrice = Math.min(minPrice, prices[i]);
            // 右指针遍历数组，result记录最大的值
            result = Math.max(result, prices[i] - minPrice);
        }
        return result;
    }
}