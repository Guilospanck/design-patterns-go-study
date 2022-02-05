# Adapter
Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

## Conceptual Example
We have a client code that expects some features of an object (`Lightning port`), but we have another object called `adaptee` (Windows laptop) which offers the same functionality but through a different interface (`USB port`).

The adapter accepts a Lightning connector and then translates its signals into a USB format and passes them to the USB port in windows laptop.
