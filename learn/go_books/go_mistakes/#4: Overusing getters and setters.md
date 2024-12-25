# Overusing getters and setters（不要过度使用）
* 如果的确有用，也可以用
    ```
    type Person struct {
        balance string  // balance命名规则
    }
    
    func (b Person) Balance() string{  // getter method
        return b.balance
    }
    func (b *Person) SetBalance(balance string) {  // setter method
        b.balance = balance
    }
    ```