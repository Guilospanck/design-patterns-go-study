# Abstract Factory

## Intent
It's a creational pattern that lets you produce families of related objects without specifying their concrete classes.

![image](https://user-images.githubusercontent.com/22435398/152313890-f55659ac-ddd4-4986-ac3e-33f93f1ba75a.png)

## Applicability
Use the Abstract Factory when your code needs to work with various families of related products, but don't want it to
depend on the concrete classes of those products - <b>they might be unknown beforehand or you simply want to allow for future
extensibility</b>.

## How to implement
1) Map out a matrix of distinct product types versus variants of these products.

2) Declare abstract product interfaces for all product types. Then make all concrete product classes implement these interfaces.

3) Declare the abstract factory interface with a set of creation methods for all abstract products.

4) Implement a set of concrete factory classes, one for each product variant.

5) Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes, depending on the application configuration or the current environment. Pass this factory object to all classes that construct products.

6) Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate creation method on the factory object

## Conceptual Example
Say, you need to buy a sports kit, a set of two different products: a pair of shoes and a shirt.
You would want to buy a full sports kit of the same brand to match all the items.

If we try to turn this into code, the abstract factory will help us create sets of products so that they would always match each other.

![Screenshot 2022-02-03 061628](https://user-images.githubusercontent.com/22435398/152315235-9e69303d-570d-4f7c-ac37-24c6a6262df5.png)
