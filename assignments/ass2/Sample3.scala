// Find the longest subsequence in the string that has upper and lowercase letters
// Using Scala's collection functions
object Sample3
 {
    def main(args: Array[String]) 
    {
        def longestUpperLower(x: String): String = 
        { 
            val answer = x.split("[0-9]") 
                    .filter(s => s.exists(ch => ch.isUpper)) 
                    .maxBy(s => s.length) 
                    
            return answer
        }

        println(longestUpperLower("2dkYhj3dhjHjsgH4Kjk2Poshk4kj3aaa"))
    }
 }