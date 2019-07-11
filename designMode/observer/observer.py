class Publisher(object):
    def __init__(self, name):
        self.name = name
        self.observers = []
        self._data = 0

    def add(self, observer):
        if observer not in self.observers:
            self.observers.append(observer)
        else:
            print("duplicate observer")

    def remove(self, observer):
        try:
            self.observers.remove(observer)
        except ValueError:
            print("failed to remove {}".format(observer))

    @property
    def data(self):
        return self._data

    @data.setter
    def data(self, value):
        self._data = int(value)
        self.notify()

    def notify(self):  # 调用观察者的回调
        [o.notify_by(self) for o in self.observers]


class Consumer(object):
    def notify_by(self, publisher):
        print("{}: '{}' has now data = {}".format(
            type(self).__name__,
            publisher.name,
            publisher.data
        ))


pb = Publisher("chenxl")
cm = Consumer()
pb.add(cm)
pb.data = 2
pb.data = 3

