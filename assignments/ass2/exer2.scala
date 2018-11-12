object exer2 
{
   def main(args: Array[String]) 
   {

        def divisors(n: Int): List[Int] =
            for (i <- List.range(2, (n/2) + 1) 
                if n % i == 0) 
                    yield i
    
        def primes(n: Int): List[Int] =
            for (i <- List.range(2, n+1) 
                if divisors(i).isEmpty) 
                    yield i

        def join(sep: String, list: List[Any]): String = 
            return list.mkString(sep)

        def pythagorean(n: Int): List[(Int, Int, Int)] =
            for (c <- List.range(1, n+1); b <- List.range(1, c+1); a <- List.range(1, c+1)
                if (a*a + b*b == c*c) && a < b)
                    yield (a, b, c)

        
        // Testing divisors()
        println("divisors(30) = " + divisors(30))
        println("divisors(64) = " + divisors(64))
        println("divisors(127) = " + divisors(127))

        // Testing primes()
        println("primes(7) = " + primes(7))
        println("primes(10) = " + primes(100))
        // Testing join()
        var list = List("one", "two", "three")
        var sep = ", "
        println("joining list with ', ' = " + join(sep, list))

        list = List("1", "2", "3")
        sep = "+"
        println("joining list with '+' = " + join(sep, list))

        println("pythagorean(10) = " + pythagorean(10))
        println("pythagorean(30) = " + pythagorean(30))

   }
}