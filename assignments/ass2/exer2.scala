object exer2 {
   def main(args: Array[String]) {

        def divisors(n: Int): List[Int] =
            for (i <- List.range(1, n+1) if n % i == 0) yield i
    
        def primes(n: Int): List[Int] =
            for (i <- List.range(2, n+1) if divisors(i).isEmpty) yield i

        def join(sep:String, list:List[Any]): String = 
            return list.mkString(sep)
        
        // Testing divisors()
        println(divisors(10))

        // Testing primes()
        println(primes(100))

        // Testing join()
        var list = List("apple", "banana", "cherry")
        var sep = ","
        println(join(sep, list))

        list = List("1", "2", "3")
        sep = "+"
        println(join(sep, list))
    
   }
}