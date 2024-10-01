# UNDERSTANDING CONCURRENCY IN GOLANG
Concurrency means dealing with a lot of things at the same time.
Concurrency means the simultaneous and independent execution of 
multiple processes.
Simultaneously making progress on more than one task is known as **concurrency**.

**SIMPLEST DEFINITION**: Concurrency refers to **managing** multiple tasks that can
progress **independently**

**NOTE**:
When two or more threads share or access the same resource, 
they don’t lead to race conditions, deadlock, or other similar problems

Independent calculations can be performed in any sequence to produce
the same result. This is called **Concurrency**.

### WHAT IS CONTEXT SWITCHING?
- **Context Switching**: It refers to the process by which a CPU switches 
from executing one task (or thread) to another. 
- It involves **saving** the state of the current task and **loading** 
the state of the next task.
- ##### Understanding State in Context Switching
-**State** refers to all the information that describe the current 
condition of a task or process.
- Components of the state include:
    - **CPU Registers**: Small storage locations in the CPU that 
        hold data or addresses temporarily while CPU is working on tasks.
    - **Program Counter (PC)**: A register that holds the address of the 
        next instruction to execute.
    - **Stack Pointer (SP)**: A register that points to the top of the stack, 
        which is used to keep track of function calls and local variables.

### Concurrency vs Parallelism
- In **single-core hardware**, we can achieve concurrency using context switching
- In **multicore hardware**, we can achieve concurrency using parallelism.
- **Parallelism** is the paradigm that is used to utilize the hardware's power completely.

#### Working with Kernel Threads
kernel threads—including the creation of threads, destruction of threads, 
context switching, and changing the value of registers—is a very costly operation
#### Golang has the following built-in concurrency constructs:
1. Goroutines: Goroutines are functions executing concurrently with other
 goroutines in the same thread or set of threads. It runs independently of 
 the function or goroutine from which it starts
2. Channels: A channel is a pipeline for sending and receiving data 
to communicate between goroutines.
3. Select: The select clause chooses one out of several communications.

### What is a Goroutine?
A goroutine is a lightweight thread managed by the Go runtime. 
It is a function that runs concurrently with other functions. 
Goroutines are multiplexed onto multiple OS threads so if one should block, 
such as waiting for I/O, others continue to run.

### What is a Process?
**A process** is composed of one or more operating system threads that are 
simultaneously executing **entities** that share the same address space

**A process** is an independently executing entity that runs in its own 
address space in memory or resource.

### What is Parallelism?
**Parallelism** is the ability to make things run quickly by using multiple
 processors simultaneously.
Parallelism requires multiple **cores or processors** to run at the same time.
**SIMPLE DEFINITION**: Parallelism refers to the **simultaneous** execution of tasks.