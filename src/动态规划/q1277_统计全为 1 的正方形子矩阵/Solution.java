/**
 * 动态规划
 * dp状态方程：若构成了一个矩形，则dp[i][j]=dp[i-1][j-1]+1;否则为1
 * 可以改进的操作：在matrix数组最上和最左加上全是0的数，避免判断是否越界 
 */
class Solution {
    public int countSquares(int[][] matrix) {
        int r = matrix.length;
        int c = matrix[0].length;
        int res = 0;
        for (int i = 0; i < r; i++) {
            for (int j = 0; j < c; j++) {
                if (matrix[i][j] == 1) {
                    if (i == 0 || j == 0) {
                        res += matrix[i][j];
                        continue;
                    } else {
                        matrix[i][j] = Math.min(matrix[i - 1][j], Math.min(matrix[i][j - 1], matrix[i - 1][j - 1])) + 1;
                    }
                }
                res += matrix[i][j];
            }
        }
        return res;
    }
}