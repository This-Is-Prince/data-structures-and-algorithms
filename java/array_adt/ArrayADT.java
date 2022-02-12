package array_adt;

public class ArrayADT {
    private int length;
    private int size;
    private int[] A;

    public ArrayADT() {
        this.length = 0;
        this.size = 0;
        this.A = new int[this.size];
    }

    public ArrayADT(int size) {
        this.length = 0;
        this.size = size;
        this.A = new int[this.size];
    }

    public ArrayADT(int[] arr) {
        this.length = arr.length;
        this.size = arr.length;
        this.A = new int[this.size];
        for (int i = 0; i < this.length; i++) {
            this.A[i] = arr[i];
        }
    }

    public int getSize() {
        return this.size;
    }

    public int getLength() {
        return this.length;
    }

    private void increaseArraySizeToDouble() {
        this.size = this.size * 2;
        int temp[] = new int[this.size];
        for (int i = 0; i < this.length; i++) {
            temp[i] = this.A[i];
        }
        this.A = temp;
    }

    private void decreaseArraySizeToHalf() {
        this.size = (int) (this.size / 2);
        int temp[] = new int[this.size];
        for (int i = 0; i < this.length; i++) {
            temp[i] = this.A[i];
        }
        this.A = temp;
    }

    public void append(int value) {
        if (this.size == this.length) {
            this.increaseArraySizeToDouble();
        }
        this.A[this.length] = value;
        this.length++;
    }

    public void display() {
        for (int i = 0; i < this.length; i++) {
            if (i == this.length - 1) {
                System.out.print(this.A[i]);
            } else {
                System.out.print(this.A[i] + ", ");
            }
        }
        System.out.println();
    }
}
