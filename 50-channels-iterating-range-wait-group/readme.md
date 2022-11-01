# Channels

- One of the most import ways to handle thread in go
- Channels are a way to communicate between goroutines
- Channels are typed, you can only send and receive values of a specific type
- Channels create a communication between thread/goroutines
- Channels provide security to know the correct moment we can update or read a value
- Channels are used when you have shared resources between two or more goroutines/threads
- Channels between goroutines

## How that works?

- initial state
- Thread 2 is trying to read from Chan, but since there is no value it will wait until has some

```text
+-----------+   +------+       +-----------+
|  Thread 1 |   | Chan | <-x-- |  Thread 2 |
+-----------+   +------+       +-----------+
                Val: nil
```

- Thread 1 insert a value in the channel, so Thread 2 can read it, and then Thread 1 can continue, only when Thread 2 consumes the value from Thread 1

```text
+-----------+    +------+    +-----------+
|  Thread 1 | -> |  10  | <- |  Thread 2 |
+-----------+    +------+    +-----------+
                Val: 10
```