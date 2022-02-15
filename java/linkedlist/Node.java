package linkedlist;

public class Node {
    protected int data;

    protected Node(int data) {
        this.data = data;
    }

    public void print() {
        System.out.print(this.data);
    }

    public void print(String end) {
        System.out.print(this.data + end);
    }

    public void println(String end) {
        System.out.println(this.data + end);
    }

    public void println() {
        System.out.println(this.data);
    }
}
