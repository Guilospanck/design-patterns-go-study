# Abstract Factory

## Intent
It's a creational pattern that lets you produce families of related objects without specifying their concrete classes.

![image](https://user-images.githubusercontent.com/22435398/152313890-f55659ac-ddd4-4986-ac3e-33f93f1ba75a.png)

## Applicability
Use the Abstract Factory when your code needs to work with various families of related products, but don't want it to
depend on the concrete classes of those products - <b>they might be unknown beforehand or you simply want to allow for future
extensibility</b>.

## Conceptual Example
Say, you need to buy a sports kit, a set of two different products: a pair of shoes and a shirt.
You would want to buy a full sports kit of the same brand to match all the items.

If we try to turn this into code, the abstract factory will help us create sets of products so that they would always match each other.
