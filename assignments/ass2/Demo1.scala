// Counts the number of characters in a List/Sequence using foldl
object Demo1 
{
   def main(args: Array[String]) 
   {
        val list = List("how many", "character are", "in", "this", "list?") 
        val count = list.foldLeft(0) ((len, word) => len + word.length)
        println(count)
   }
}