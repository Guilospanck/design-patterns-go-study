# Strategy
It's a structural design pattern that lets you define a family of algorithms, put each of them into a separate class and make their objects interchangeable.

It turns a set of behaviors into objects and makes them interchangeable inside the original context object.

The original object, called context, holds a reference to a strategy object and delegates it executing the behavior. In order to change the way the context performs its work, other objects may replace the currently linked strategy object with another one.