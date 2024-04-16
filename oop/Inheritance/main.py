class Human:
    def __init__(self, name):
        self.__name = name

    def get_name(self):
        return self.__name


## don't touch above this line


class Archer(Human):
    def __init__(self, name, num_arrows):
        super().__init__(name)
        self.__num_arrows = num_arrows

    def get_num_arrows(self):
        return self.__num_arrows

    def use_arrows(self, num):
        if self.__num_arrows < num:
            raise Exception("not enough arrows")
        self.__num_arrows -= num


class Crossbowman(Archer):
    def __init__(self, name, num_arrows):
        super().__init__(name, num_arrows)

    def triple_shot(self, target):
        self.use_arrows(3)
        return f"{target.get_name()} was shot by 3 crossbow bolts"

# Sword types with polymorhism

class Sword:
    def __init__(self, sword_type):
        self.sword_type = sword_type

    def __add__(self, other):
        if self.sword_type != other.sword_type:
            raise Exception("can not craft")
        if self.sword_type == "bronze":
            return Sword("iron")
        elif self.sword_type == "iron":
            return Sword("steel")
        else:
            raise Exception("can not craft")
