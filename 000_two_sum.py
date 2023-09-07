class Solution:

    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype List[int]
        """

        for i in nums:
            j = target - i
            start_index = nums.index(i)
            next_index = start_index + 1
            temp_nums = nums[next_index:]
            if j in temp_nums:
                print(nums.index(i), next_index + temp_nums.index(j))
                return (nums.index(i), next_index + temp_nums.index(j))

def add(first: int = 10, second: float = 5.5) -> float: 
    return first + second

if __name__ == "__main__":
    # "asdfasdf".join()

    aaa = ['a', 'b', 1, 'd', 'e']
    add(123,44.4)
    print(str.encode("safasdf"))
