import java.util.Comparator;
import java.util.HashMap;
import java.util.PriorityQueue;

/**
 * 关键是使用优先队列和优先队列比较器的判定
 * 优先队列中间接使用了最小堆的思想
 */
class Solution {

    public int[] topKFrequent(int[] nums, int k) {
        // 将元素的出现个数放入HashMap中保存
        // key:元素  value:出现次数
        HashMap<Integer, Integer> count = new HashMap();
        for (int n: nums) {
            count.put(n, count.getOrDefault(n, 0) + 1);
        }

        // 创建优先队列并传入比较器
        PriorityQueue<Integer> heap = new PriorityQueue<>(Comparator.comparingInt(count::get));
        // 向优先队列里加入元素，若size大于K，则将频率小的弹出
        for (int n : count.keySet()) {
            heap.add(n);
            if (heap.size() > k) {
                heap.poll();
            }
        }

        int length = heap.size();
        int[] result = new int[heap.size()];
        for(int i = 0; i < length; i++) {
            result[i] = heap.poll();
        }

        return result;
    }
}