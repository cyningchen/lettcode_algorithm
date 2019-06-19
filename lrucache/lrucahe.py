#! coding:  utf-8

from collections import OrderedDict


class LruCache(object):
    def __init__(self, capacity):
        self.od = OrderedDict()
        self.cap = capacity

    def get(self, key):
        if key in self.od:
            value = self.od[key]
            self.od.move_to_end(key)
            return value
        else:
            return -1

    def put(self, key, value):
        if key in self.od:
            del self.od[key]
            self.od[key] = value
        else:
            self.od[key] = value
            if len(self.od) > self.cap:
                self.od.popitem(last=False)


if __name__ == '__main__':
    lru = LruCache(2)
    lru.put(1, "a")
    lru.put(2, "b")
    lru.get(1)
    lru.put(3, "c")
    print(lru.get(2))
