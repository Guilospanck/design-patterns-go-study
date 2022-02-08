# Flyweight
It's a structural pattern where an object stores only its intrinsic state, not its extrinsic one. This is used in order to reduce the amount of RAM memory needed to run some application.

Intrinsic it's something that doesn't change. Something deep inside the object.

Extrinsic, in the other hand, it's something that changes.

## Applicability
Use the `Flyweight` pattern only when your program must support a huge number of objects which barely fit into available RAM.
It's most useful when:
- an application needs to spawn a huge number of similar objects
- this drains all available RAM on a target device
- the object contains duplicate states which can be extracted and shared between multiple objects