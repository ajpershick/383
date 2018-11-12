object exer5 
{
    def main(args: Array[String]) 
    {
        // Pattern matching in Scala is similar to Haskell with the case method
        def fib(n: Int): Int = n match 
        {
            case 0 | 1 => n
            case _ => fib(n-1) + fib(n-2)
        }
        // Test cases from exercise
        println("fib(0) = " + fib(0))
        println("fib(1) = " + fib(1))
        println("fib(2) = " + fib(2))
        println("fib(3) = " + fib(3))
        println("fib(10) = " + fib(10))
        println("fib(20) = " + fib(20))
    }
}

// Can't get it to work with Streams :(
