object exer5 {
    def main(args: Array[String]) 
    {
        // Takes advantage of stream processing to 
        // run the operations in parallel
        val fibs: Stream[Int] = 0 #:: fibs.scanLeft(1)(_ + _)
    }
}

// Can't get it to work with Streams :(
