object exer4 
{
    def main(args: Array[String]) 
    {
        def hailstone(n: Int): Int =
        {
            if (n % 2 == 0)
            {
		        return (n / 2)
            } 
            else 
            {
		        return 3*n + 1
            }
        }
            
        def hailSeq(n: Int): List[Any] =
        {
            var list = List(n)
            var i = n
            while (i > 1)
            {
                i = hailstone(i)
                list = list :+ i
            }
            return list
        }

        def mergeSort[T](pred: (T, T) => Boolean)(xs: List[T]): List[T] = 
        {
            def merge(ls: List[T], rs: List[T]): List[T] = (ls, rs) match 
            {
                case (List(), _) => rs
                case (_, List()) => ls
                case (l :: ls1, r :: rs1) =>
                    if (pred(l, r)) l :: merge(ls1, rs)
                    else r :: merge(ls, rs1)
            }   

            val m = xs.length / 2;

            if (m == 0) 
                xs;

            else 
            {
                val (l, r) = xs splitAt m;
                merge(mergeSort(pred)(l), mergeSort(pred)(r));
            }
        }

        // Testing hailSeq()
        println("hailSeq(1) = " + hailSeq(1))
        println("hailSeq(11) = " + hailSeq(11))
        println("hailSeq(6) = " + hailSeq(6))
        
        // Testing mergeSort on integers
        val intSort = mergeSort((_: Int) < (_: Int)) _ // intSort is just a version of mergeSort for sorting ints
        println("mergeSort(6,2,4,8,4,1) = " + intSort(List(6,2,4,8,4,1)));
    }
}