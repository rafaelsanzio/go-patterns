### <b>Prototype</b> is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

### Applicability

- Use the Prototype pattern when your code shouldn’t depend on the concrete classes of objects that you need to copy.

- Use the pattern when you want to reduce the number of subclasses that only differ in the way they initialize their respective objects.

### How to implement

- Create the prototype interface and declare the clone method in it. Or just add the method to all classes of an existing class hierarchy, if you have one.

- A prototype class must define the alternative constructor that accepts an object of that class as an argument. The constructor must copy the values of all fields defined in the class from the passed object into the newly created instance. If you’re changing a subclass, you must call the parent constructor to let the superclass handle the cloning of its private fields.

- The cloning method usually consists of just one line: running a new operator with the prototypical version of the constructor. Note, that every class must explicitly override the cloning method and use its own class name along with the new operator. Otherwise, the cloning method may produce an object of a parent class.

- Optionally, create a centralized prototype registry to store a catalog of frequently used prototypes.
