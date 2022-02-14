## Conceptual Example
Suppose you are building an In-Memory-Cache. Since it’s in memory, it has a limited size. Whenever it reaches its maximum size, some entries have to be evicted to free-up space. This can happen via several algorithms. Some of the popular algorithms are:

- Least Recently Used (LRU): remove an entry that has been used least recently.
- First In, First Out (FIFO): remove an entry that was created first.
- Least Frequently Used (LFU): remove an entry that was least frequently used.

The problem is how to decouple our cache class from these algorithms so that we can change the algorithm at run time. Also, the cache class should not change when a new algorithm is being added.

This is were Strategy pattern comes into the picture.