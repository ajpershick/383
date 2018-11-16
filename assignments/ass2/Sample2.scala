 // Converting lowercase to lowercaseUPPERCASE alternating using flatMap
 // This could become handy if you're building a Lexographical ordering system
 object Sample2 
 {
    def main(args: Array[String]) 
    {
        val lowercase = Seq('a', 'b', 'c', 'd')  // 'val' creates immutable variables
        val alternating = lowercase.flatMap(x => List(x.toUpper, x))
        println(alternating)
    }
 }