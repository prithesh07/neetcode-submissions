class MinStack:
    def __init__(self):
        self.stack1 = []
        self.stack2 = []

    def push(self, val: int) -> None:
        self.stack1.append(val)
        if len(self.stack2) == 0 or self.stack2[-1] > val:
            self.stack2.append(val)
        else:
            self.stack2.append(self.stack2[-1])

    def pop(self) -> None:
        self.stack1.pop()  
        self.stack2.pop()  

    def top(self) -> int:
        return self.stack1[-1]

    def getMin(self) -> int:
        return self.stack2[-1]
        
