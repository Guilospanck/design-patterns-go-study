# Builder

## Intent
It's a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce
diffrent types and representations of an object using the same construction code.

## Applicability
- Use the Builder pattern to get rid of a "telescopic constructor".
- Use the Builder pattern when you want your code to be able to create different representations of some product.
- Use the Builder pattern to construct Composite trees or other complex objects.

## How to implement
1) Make sure that you can clearly define the common construction steps for building all available product representations. Otherwise, you won’t be able to proceed with implementing the pattern.

2) Declare these steps in the base builder interface.

3) Create a concrete builder class for each of the product representations and implement their construction steps.

Don’t forget about implementing a method for fetching the result of the construction. The reason why this method can’t be declared inside the builder interface is that various builders may construct products that don’t have a common interface. Therefore, you don’t know what would be the return type for such a method. However, if you’re dealing with products from a single hierarchy, the fetching method can be safely added to the base interface.

5) Think about creating a director class. It may encapsulate various ways to construct a product using the same builder object.

5) The client code creates both the builder and the director objects. Before construction starts, the client must pass a builder object to the director. Usually, the client does this only once, via parameters of the director’s constructor. The director uses the builder object in all further construction. There’s an alternative approach, where the builder is passed directly to the construction method of the director.

6) The construction result can be obtained directly from the director only if all products follow the same interface. Otherwise, the client should fetch the result from the builder.

## Conceptual Example
The Builder pattern is used when the desired product is complex and requires multiple steps to complete. The Builder pattern keeps the product private until it's fully built.

We're going to create two different houses: <code>igloo</code> and <code>normalHouse</code>