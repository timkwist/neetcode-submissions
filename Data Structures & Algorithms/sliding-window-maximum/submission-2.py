import heapq
class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        h = []
        ret = []
        for i in range(len(nums)):
            heapq.heappush(h, [-1*nums[i], i])
            while h[0][1] <= i - k:
                t_n, t_i = heapq.heappop(h)
            if i >= k - 1:
                ret.append(-1*h[0][0])
        return ret
