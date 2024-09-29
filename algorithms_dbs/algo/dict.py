def merge(dict1, dict2):
    new_dict = dict1.copy()
    for key in dict2.keys():
        new_dict[str(key)] = str(dict2[key])
        print(dict2[key])
        print(key)
    return new_dict

two_towers = {"Frodo": "One Ring", "Aragorn": "Narsil"}
rotk = {"Aragorn": "Andúril", "Gandalf": "Glamdring"}
merged_dict = merge(two_towers, rotk)
print("dict1:", two_towers)
print("#############################################################")
print("dict2:", rotk)
print("#############################################################")
print("dict3:", merged_dict)
# {'Frodo': 'One Ring', 'Aragorn': 'Andúril', 'Gandalf': 'Glamdring'}

def calculate_total(items_purchased, pinned_list):
    item_prices = {
        "health_potion": 10.00,
        "mana_potion": 12.00,
        "gold_dust": 5.00,
        "dwarven_ale": 8.00,
        "enchanted_scroll": 25.00,
        "ice_cold_milk": 50.00,
        "herbs": 7.00,
        "crystal_shard": 20.00,
        "magic_ring": 100.00,
        "mystic_amulet": 150.00,
    }

    # Don't touch above this line

    unpurchased_items = []
    for i in pinned_list:
        found = False
        for j in items_purchased:
            if i == j:
                found = True
        if not found:
            unpurchased_items.append(i)

    receipt = {}
    total_cost = 0
    for i in items_purchased:
        receipt[str(i)] = item_prices[str(i)]
        total_cost += item_prices[str(i)]

    return unpurchased_items, receipt, total_cost
