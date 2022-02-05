# Singleton
This design pattern is not complex. Basically it gives you only one instance for program execution. It's handy when, for example, you got adatabase connection and you just want to open it one time.

But, even so, it's really easy to get into trouble using this design pattern. It can become a mess.


In `Go`, we can also use the `init` function to create a single instance because this function is called only once per file, but this is only achievable when we can declare the instance right away.