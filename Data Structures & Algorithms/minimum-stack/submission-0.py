
# stack = []          # Create empty stack
# stack.append("A")   # Push item 'A'
# stack.append("B")   # Push item 'B'
# stack.pop()         # Pop item 'B'
# print(stack[-1])    # Peek top item ('A')


class MinStack:

    def __init__(self):
        self.stack1 = []
        self.stack2 = []

    def push(self, val: int) -> None:
        self.stack1.append(val)
        if len(self.stack2) == 0:
            self.stack2.append(val)
        else:
            temp = self.stack2[-1]
            if temp > val:
                self.stack2.append(val)
            else:
                self.stack2.append(temp)

    def pop(self) -> None:
        if len(self.stack1) <= 0:
            return -1
        
        self.stack1.pop()  
        self.stack2.pop()  

    def top(self) -> int:
        if len(self.stack1) <= 0:
            return -1

        return self.stack1[-1]

    def getMin(self) -> int:
        if len(self.stack1) <= 0:
            return -1
            
        return self.stack2[-1]
        
