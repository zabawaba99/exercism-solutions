public class HelloWorld {

    /**
     * Creates a unique greeting.
     *
     * @param name - The name of the person to greet
     * @return the generated greeting
     */
    public static String hello(String name) {
        if (name == null || name.isEmpty() ) {
            name = "World";
        }

        return String.format("Hello, %s!", name);
    }
}

