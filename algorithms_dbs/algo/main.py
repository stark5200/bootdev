def get_avg_brand_followers(all_handles, brand_name):
    list_count = len(all_handles)
    brand_count = 0
    for list in all_handles:
        for follower in list:
            if brand_name in follower:
                brand_count+=1
    return brand_count/list_count