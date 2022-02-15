package linkedlist;

public class Debug {
    public void debug() {
        SingleLinkedList arr = new SingleLinkedList(new int[] { 3, 5, 7, 10, 15 });
        arr.displayUsingLoop();
        arr.displayUsingRecursion();
        arr.displayUsingLoopInReverseOrder();
        arr.displayUsingRecursionInReverseOrder();
    }
}
