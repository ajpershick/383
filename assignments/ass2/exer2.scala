object exer2 {
   def main(args: Array[String]) {

        def divisors(n: Int): List[Int] =
            for (i <- List.range(1, n+1) if n % i == 0) yield i
    
        def primes(n: Int): List[Int] =
            for (i <- List.range(2, n+1) if divisors(i).isEmpty) yield i

        // def join(sep:String, list:List[String]): List[String] = 
            // list.mkString(sep);
        
        println(divisors(10));
        println(primes(100));
        // val list = Array("apple", "banana", "cherry");
        // val sep = ","
        // println(join(sep, list));

    
   }
}