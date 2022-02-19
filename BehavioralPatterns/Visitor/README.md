# Visitor
It's a behavioral pattern that lets you separate algorithms from the objects on which they operate.

## Example
Let's say that we have some shape classes and we want to export them as XML.
We already have those classes running in production, so we don't want to change much of them and get ourselves some bug.
And also we want to concentrate this function of "exporting a XML" in just a single file.

We can define an interface for this various shapes (if they don't already have one) and add a method called `accept`.
This method `accept` will accept a `visitor type (interface)` and he will its method from this visitor.
For example, if a `Dot` receives an `accept` call, it will call the `IVisitor.visitDot(d: Dot)` and so on.

The `IVisitor` interface will have all methods necessary for the "visitation" of each shape class.
Therefore, it will have:
```go
visitDot(d: Dot)
visitCircle(c: Circle)
... and so on
```

By having this `IVisitor` interface, we can then have other different functions, like "exporting to PDF". For this we would only create a different concrete visitor.

If we were to have new different shapes, we'd just add their methods into the `IVisitor` interface and in its subsequent concrete implementations.