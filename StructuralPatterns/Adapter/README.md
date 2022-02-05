# Adapter
Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

## Applicability
Use the Adapter class when you want to use some existing class, but its interface isn't compatible with the rest of your code.

## Conceptual Example
We have a client code that expects some features of an object (`Lightning port`), but we have another object called `adaptee` (Windows laptop) which offers the same functionality but through a different interface (`USB port`).

The adapter accepts a Lightning connector and then translates its signals into a USB format and passes them to the USB port in windows laptop.

## Relation With Other Patterns
`Adapter` changes the interface of an existing object. `Decorator` enhances it.