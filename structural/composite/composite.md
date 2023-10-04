### <b>Composite</b> is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.

### Applicability

- Use the Composite pattern when you have to implement a tree-like object structure.
- Use the pattern when you want the client code to treat both simple and complex elements uniformly.

### How to implement

- Make sure that the core model of your app can be represented as a tree structure. Try to break it down into simple elements and containers. Remember that containers must be able to contain both simple elements and other containers.
- Declare the component interface with a list of methods that make sense for both simple and complex components.
- Create a leaf class to represent simple elements. A program may have multiple different leaf classes.
- Create a container class to represent complex elements. In this class, provide an array field for storing references to sub-elements. The array must be able to store both leaves and containers, so make sure it’s declared with the component interface type.
- Finally, define the methods for adding and removal of child elements in the container.
