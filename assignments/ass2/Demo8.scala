import Array._

object Demo8 {
   def main(args: Array[String]) {
      var myList1 = range(0, 50, 5)
      var myList2 = range(7, 21)

      // Print all the array elements
      for ( x <- myList1 ) {
         print( " " + x )
      }
      
      println()
      for ( x <- myList2 ) {
         print( " " + x )
      }
   }
}