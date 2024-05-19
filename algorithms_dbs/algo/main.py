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
    pass

    
    

    

