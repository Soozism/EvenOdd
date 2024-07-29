# EvenOdd

The **EvenOdd** project provides a simple example of using slices and `for` loops in Go to classify numbers as even or odd. This program iterates through a slice of integers and prints whether each number is even or odd.

## Overview

The program performs the following steps:
1. Initializes a slice with integers from 0 to 10.
2. Iterates over the slice using a `for` loop.
3. Checks if each number is even or odd.
4. Prints the result to the console.

## Usage

To run the program, ensure you have Go installed on your system. Then, follow these steps:

1. Save the code to a file, for example, `main.go`.
2. Open a terminal and navigate to the directory containing `main.go`.
3. Run the following command to execute the program:

```bash
go run main.go
```

## Code Explanation

The main components of the code are:

1. **Creating a Slice:**
   ```go
   ns := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
   ```

   A slice named `ns` is created and initialized with integers from 0 to 10.

2. **Iterating Over the Slice:**
   ```go
   for _, n := range ns {
       if n%2 == 0 {
           fmt.Printf("%v is even \n", n)
       } else {
           fmt.Printf("%v is odd \n", n)
       }
   }
   ```

   - The `for` loop iterates over each element in the slice.
   - The `if` statement checks if the number is even (i.e., divisible by 2) and prints "even" if true; otherwise, it prints "odd".

## Output

Running the program will produce the following output:

```
0 is even 
1 is odd 
2 is even 
3 is odd 
4 is even 
5 is odd 
6 is even 
7 is odd 
8 is even 
9 is odd 
10 is even 
```

## Notes

- Slices in Go provide a flexible way to work with collections of data.
- The `for` loop with `range` allows you to iterate over the elements of a slice.
- The `%` operator is used to determine if a number is divisible by 2, helping to classify numbers as even or odd.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you have any suggestions or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
