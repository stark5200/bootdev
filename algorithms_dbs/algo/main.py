def get_avg_brand_followers(all_handles, brand_name):
    list_count = len(all_handles)
    brand_count = 0
    for list in all_handles:
        for follower in list:
            if brand_name in follower:
                brand_count+=1
    return brand_count/list_count
  
def find_last_name(names_dict, first_name):
    if first_name:
        if first_name in names_dict:
            return names_dict[first_name]
    return None
  
def binary_search(target, arr):
    if arr == None or target == None or len(arr) == 0:
        return False
    n = len(arr)
    low, high = 0, n-1
    if target < arr[low] or target > arr[high]:
        return False
    median = low + high // 2
    while low <= high:
        if arr[median] == target or arr[low] == target or arr[high] == target:
            return True
        if target < arr[median]:
            high = median - 1
        if target > arr[median]:
            low = median + 1
        median = (low + high) // 2
        print("target: " + str(target) + ", arr[low]:" + str(arr[low]) + ", arr[high]:" + str(arr[high]) + ", low:" + str(low) + ", median:" + str(median) + ", high:" + str(high) + ".")
    return False
  
def count_names(list_of_lists, target_name):
    count = 0
    for list in list_of_lists:
        for name in list:
            if name == target_name:
                count += 1
    return count
  
def bubble_sort(nums):
    size = len(nums)
    for i in range(size):
        for j in range(i+1, size):
            if nums[i] > nums[j]:
                temp = nums[i]
                nums[i] = nums[j]
                nums[j] = temp
    return nums

def bubble_sort_bootdev(nums):
    swapping = True
    end = len(nums)
    while swapping:
        swapping = False
        for i in range(1, end):
            if nums[i - 1] > nums[i]:
                temp = nums[i - 1]
                nums[i - 1] = nums[i]
                nums[i] = temp
                swapping = True
        end -= 1
    return nums

def merge_sort(nums):
    if len(nums) < 2:
        return nums
    sorted1 = merge_sort( nums[:len(nums)//2] )
    sorted2 = merge_sort( nums[len(nums)//2 :] )
    return merge(sorted1, sorted2)
     


def merge(first, second):
    sorted = []
    i = 0
    j = 0
    while i < len(first) and j < len(second):
        if first[i] <= second[j]:
            sorted.append(first[i])
            i += 1
        else:
            sorted.append(second[j])
            j += 1
    while i < len(first):
        sorted.append(first[i])
        i += 1
    while j < len(second):
        sorted.append(second[j])
        j += 1
    return sorted
  
def insertion_sort(nums):
    sorted = []
    i = 0
    while i < len(nums):
        min = float("inf")
        index = 0
        for n in range(len(nums)):
            if nums[n] < min:
                min = nums[n]
                index = n
        del nums[index]
        sorted.append(min)
    return sorted

def insertion_sort_bootdev(nums):
    for i in range(len(nums)):
        j = i
        while j > 0 and nums[j - 1] > nums[j]:
            temp = nums[j]
            nums[j] = nums[j - 1]
            nums[j - 1] = temp
            j -= 1
    return nums
  
def quick_sort(nums, low, high):
    pass


def partition(nums, low, high):
    pass
  
def fib(n):
    if n <= 1:
        return n
    current = 0
    parent = 1
    grandparent = 0
    for i in range(0, n - 1):
        current = parent + grandparent
        grandparent = parent
        parent = current
    return current






class Influencer:
    def __init__(self, num_selfies, num_bio_links):
        self.num_selfies = num_selfies
        self.num_bio_links = num_bio_links

    def __repr__(self):
        return f"({self.num_selfies}, {self.num_bio_links})"

# dont touch above this line

def vanity(influencer: Influencer):
    return influencer.num_bio_links * 5 + influencer.num_selfies

def vanity_sort(influencers: Influencer):
    influencers_vanity = {}
    for i in influencers:
        influencers_vanity[i] = vanity(i)
    sorted_influencers = sorted(influencers_vanity.items(), key=lambda x:x[1])
    return [s[0] for s in sorted_influencers]

#  Order 2^N - Exponential
    
'''
Order 2^N - Exponential
O(2^n) is the first Big O class that we've dealt with that falls into the scary exponential category of algorithms.

Algorithms that grow at an exponential rate become impossible to compute after so few iterations that they are almost worthless in practicality.

Assignment
At Socialytics we need to be able to compute the power set of a set of influencers. It has something to do with targeting segments of an audience with ads. I don't know, I just work here.

A power set is the set of all possible subsets of a set. For example, the set {1, 2, 3} has the power set:

{
  {},
  {1},
  {2},
  {3},
  {1, 2},
  {1, 3},
  {2, 3},
  {1, 2, 3},
}
Copy icon
We'll work with Python lists instead of sets for simplicity.

Complete the power_set function using the following algorithm:

Check if the input list is empty. If it is, return a list containing an empty list. (The power set of an empty set is a set containing only the empty set)
Otherwise, create an empty list to hold all the final subsets of the input list.
Recursively call power_set. Pass in all of the elements in the input set except the first one.
Iterate over the list of subsets returned from the recursive call. For each subset, append two new subsets to the final list of subsets:
list_with_only_the_first_item_from_input_set + subset
subset
Return the list of subsets
Observe!
Notice how the power_set() output gets exponentially larger with each iteration because its complexity class is O(2^n).

If we could calculate one subset per millisecond, completing the power_set() of just 25 items would take approximately 9 hours, and that's not accounting for the massive amounts of memory we would need. 40 items would take over 34 years!
'''
    
def power_set(input_set):
    if len(input_set) == 0:
        return [[]]

    subsets = []
    first = input_set[0]
    remaining = input_set[1:]
    remaining_subsets = power_set(remaining)
    for subset in remaining_subsets:
        subsets.append([first] + subset)
        subsets.append(subset)
    return subsets

    

