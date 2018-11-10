object exer4 {
    def main(args: Array[String]) 
    {
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

        val intSort = mergeSort((_: Int) < (_: Int)) _
        println(intSort(List(6,2,4,8,4,1)));
    }
}