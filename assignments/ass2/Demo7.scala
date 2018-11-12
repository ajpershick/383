import Array._

// Concatenates two arrays together, then prints new array
object Demo7 {
   def main(args: Array[String]) {
      var myList1 = Array(1.9, 2.9, 3.4, 3.5)
      var myList2 = Array(8.9, 7.9, 0.4, 1.5)
      var myList3 =  concat( myList1, myList2)
      // Print array elements out one-by-one
      for (i <- myList3) {
         println(i)
      }
   }
}