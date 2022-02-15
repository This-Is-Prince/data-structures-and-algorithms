package linkedlist;

public class Debug {
    public void debug() {
        SingleLinkedList arr = new SingleLinkedList(new int[] { 3, 5, 7, 10, 15 });
        arr.displayUsingLoop();
        arr.displayUsingRecursion();
        arr.displayUsingLoopInReverseOrder();
        arr.displayUsingRecursionInReverseOrder();
        System.out.println(arr.countNoOfNodesUsingLoop());
        System.out.println(arr.countNoOfNodesUsingRecursion());
        System.out.println(arr.sumUsingLoop());
        System.out.println(arr.sumUsingRecursion());
    }
}
