package linkedlist;

class SNode extends Node {
    /**
     * Next pointer, points to next node in a linked list
     */
    SNode next;

    /**
     * Create SNode with data passed as parameter
     * 
     * @param data integer data
     */
    public SNode(int data) {
        super(data);
        this.next = null;
    }

}

public class SingleLinkedList {
    SNode head;

    /**
     * Create a singly linked list with no data
     */
    public SingleLinkedList() {
        this.head = null;
    }

    /**
     * Create a singly linked list with data passed in parameter
     * 
     * @param arr integer array
     */
    public SingleLinkedList(int[] arr) {
        SNode temp = this.head;
        for (int i = 0; i < arr.length; i++) {
            SNode tmp = new SNode(arr[i]);
            if (this.head == null) {
                this.head = tmp;
                temp = this.head;
            } else {
                temp.next = tmp;
                temp = temp.next;
            }
        }
    }

    /**
     * Display linked list data using loops
     */
    public void displayUsingLoop() {
        SNode temp = this.head;
        while (temp != null) {
            if (temp.next == null) {
                temp.println();
            } else {
                temp.print(", ");
            }
            temp = temp.next;
        }
    }

    /**
     * Display linked list data using loops but in reverse order
     */
    public void displayUsingLoopInReverseOrder() {
        SNode last = null;
        SNode tmp = this.head;
        while (tmp != last) {
            while (tmp.next != last) {
                tmp = tmp.next;
            }
            if (tmp == this.head) {
                tmp.println();
            } else {
                tmp.print(", ");
            }
            last = tmp;
            tmp = this.head;
        }
    }

    /**
     * Display linked list data using recursion
     */
    public void displayUsingRecursion() {
        this.displayUsingRecursionUtil(this.head);
    }

    /**
     * Display linked list data using recursion but in reverse order
     */
    public void displayUsingRecursionInReverseOrder() {
        this.displayUsingRecursionInReverseOrderUtil(this.head);
    }

    /**
     * Count No of nodes in linked list, using loop and return
     * 
     * @return no of nodes
     */
    public int countNoOfNodesUsingLoop() {
        int count = 0;
        SNode tmp = this.head;
        while (tmp != null) {
            count++;
            tmp = tmp.next;
        }
        return count;
    }

    /**
     * Count No of nodes in linked list, using recursion and return
     * 
     * @return no of nodes
     */
    public int countNoOfNodesUsingRecursion() {
        return this.countNoOfNodesUsingRecursionUtil(this.head);
    }

    /**
     * Calculate sum of data in linked list using loop and return the sum
     * 
     * @return total sum
     */
    public int sumUsingLoop() {
        int sum = 0;
        SNode tmp = this.head;
        while (tmp != null) {
            sum += tmp.data;
            tmp = tmp.next;
        }
        return sum;
    }

    /**
     * Calculate sum of data in linked list using recursion and return the sum
     * 
     * @return total sum
     */
    public int sumUsingRecursion() {
        return this.sumUsingRecursionUtil(this.head);
    }

    /**
     * Util function for displayUsingRecursion function
     * 
     * @param node Linked list node
     */
    private void displayUsingRecursionUtil(SNode node) {
        if (node == null) {
            return;
        } else if (node.next == null) {
            node.println();
        } else {
            node.print(", ");
            this.displayUsingRecursionUtil(node.next);
        }
    }

    /**
     * Util function for displayUsingRecursionInReverseOrder function
     * 
     * @param node Linked List node
     */
    private void displayUsingRecursionInReverseOrderUtil(SNode node) {
        if (node == null) {
            return;
        } else {
            this.displayUsingRecursionInReverseOrderUtil(node.next);
            if (this.head == node) {
                node.println();
            } else {
                node.print(", ");
            }
        }
    }

    /**
     * Util function for countNoOfNodesUsingRecursion function
     * 
     * @param node Node of linked list
     * @return no of nodes
     */
    private int countNoOfNodesUsingRecursionUtil(SNode node) {
        if (node == null) {
            return 0;
        } else {
            return 1 + this.countNoOfNodesUsingRecursionUtil(node.next);
        }
    }

    /**
     * Util function for sumUsingRecursion function
     * 
     * @param node Node of linked list
     * @return total sum
     */
    private int sumUsingRecursionUtil(SNode node) {
        if (node == null) {
            return 0;
        } else {
            return this.sumUsingRecursionUtil(node.next) + node.data;
        }
    }
}
