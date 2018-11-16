// Good example of a use case for Scala streams. 
// When working with massive lists (big Data)
// Trying to print from a normal list will run out of memory/heap space
// But we can use stream values easily
object Sample4
 {
    def main(args: Array[String]) 
    {
        val list = (1 to 100000000).toStream
        println(list(0))  // returns 1
        println(list(1))  // returns 2

        val list2 = List.range(1, 100000000)
        println(list(0)) // Runs out of heap space, will give error
        println(list(1))
    }
 }