### <b>Abstract Factory</b> is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

### Applicability

- Use the Abstract Factory when your code needs to work with various families of related products, but you don’t want it to depend on the concrete classes of those products—they might be unknown beforehand or you simply want to allow for future extensibility.

- The Abstract Factory provides you with an interface for creating objects from each class of the product family. As long as your code creates objects via this interface, you don’t have to worry about creating the wrong variant of a product which doesn’t match the products already created by your app.

- In a well-designed program each class is responsible only for one thing. When a class deals with multiple product types, it may be worth extracting its factory methods into a stand-alone factory class or a full-blown Abstract Factory implementation.

### How to Implement

- Map out a matrix of distinct product types versus variants of these products.

- Declare abstract product interfaces for all product types. Then make all concrete product classes implement these interfaces.

- Declare the abstract factory interface with a set of creation methods for all abstract products.

- Implement a set of concrete factory classes, one for each product variant.

- Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes, depending on the application configuration or the current environment. Pass this factory object to all classes that construct products.

- Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate creation method on the factory object.
