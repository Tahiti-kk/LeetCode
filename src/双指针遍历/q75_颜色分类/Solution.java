class Solution {
    /**
     * 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
     * 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
     * 示例：
     * 输入: [2,0,2,1,1,0]
     * 输出: [0,0,1,1,2,2]
     */
    public void sortColors(int[] nums) {
        int lp = 0, rp = nums.length-1, index = 0;
        while(index <= rp) {
            switch (nums[index]) {
                case 0:
                    nums[index] = nums[lp];
                    nums[lp++] = 0;
                    // 此时lp为0或1，不会是2
                    index++;
                    break;
                case 1:
                    index++;
                    break;
                // 将2置换到尾部
                case 2:
                    nums[index] = nums[rp];
                    nums[rp--] = 2;
                    break;
                default:
                    return;
            }
        }
    }
}