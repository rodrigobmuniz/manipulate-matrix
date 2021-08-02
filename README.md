# manipulate-matrix

## About

## How to run

```
curl -F 'file=@file/matrix.csv' "localhost:8080/echo"
```

## Operations

Given an uploaded csv file

```
1,2,3
4,5,6
7,8,9
```

1. Echo (already given in Go solution)

    - Return the matrix as a string in matrix format.

    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ```

2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ```
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ```
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ```
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ```

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.

## Tests

###

```
go test -v ./...
```

```
go test -cover ./...
```
