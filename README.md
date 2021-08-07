# manipulate-matrix

## About

Project created at the request of ECore, to assess the skills and good practices of backend programmer (League Backend Challenge)

## How to generate the build

From the project's root folder run:

```
go build -o manipulate-matrix main.go
```

## How to run web server

1. Download the zipped folder
2. Extract the content to a location you have access to
3. Access the folder with the extracted content (project's root folder)
4. Follow the procedure below, according to the operating system and/or tool you have installed

    ### Linux or Mac

    Double click on the file: `manipulate-matrix`

    Or:
    From the project's root folder run the following command

    ```
    ./manipulate-matrix
    ```

    ### Windows

    Double click on the file: `manipulate-matrix.exe`

    Or:
    From the project's root folder run the following command

    ```
    ./manipulate-matrix.exe
    ```

    ### Docker

    If you have Docker and Docker-compose installed on your computer, do the following

    From the project's root folder:

    ```
    docker-compose up --build
    ```

## How to use

Inside the project's root folder, you should run the following:

### Send request

```
curl -F 'file=@[FILE_PATH]' "localhost:8080/echo"
```

```
curl -F 'file=@file/matrix.csv' "localhost:8080/echo"
```

```
// Expected output
1,2,3
4,5,6
7,8,9
```

To upload other files you must copy these files to the `file` folder inside the project folder.

## All Operations

1. Echo

    command:

    ```
    curl -F 'file=@file/matrix.csv' "localhost:8080/echo"
    ```

    Return the matrix as a string in matrix format.

    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ```

2. Invert

    command

    ```
    curl -F 'file=@file/matrix.csv' "localhost:8080/invert"
    ```

    Return the matrix as a string in matrix format where the columns and rows are inverted

    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ```

3. Flatten

    command

    ```
    curl -F 'file=@file/matrix.csv' "localhost:8080/flatten"
    ```

    Return the matrix as a 1 line string, with values separated by commas.

    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ```

4. Sum

    command

    ```
    curl -F 'file=@file/matrix.csv' "localhost:8080/sum"
    ```

    Return the sum of the integers in the matrix

    ```
    // Expected output
    45
    ```

5. Multiply

    ```
    curl -F 'file=@file/matrix.csv' "localhost:8080/multiply"
    ```

    - Return the product of the integers in the matrix

    ```
    // Expected output
    362880
    ```

    The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.

    To upload other files you must copy these files to the `file` folder inside the project folder.

## Tests

### How to run automated tests

```
go test -v ./...
```

### Automated test results

```

?   	manipulate-matrix	[no test files]
=== RUN   TestCanarySpec
--- PASS: TestCanarySpec (0.00s)
PASS
ok  	manipulate-matrix/app/delivery/http	1.832s
=== RUN   TestCanarySpec
--- PASS: TestCanarySpec (0.00s)
=== RUN   TestHaveErrorReturningFalse
--- PASS: TestHaveErrorReturningFalse (0.00s)
=== RUN   TestHaveErrorReturningTrue
--- PASS: TestHaveErrorReturningTrue (0.00s)
=== RUN   TestCheckIfAllValuesAreConvertibleToNumberReturningTrue
--- PASS: TestCheckIfAllValuesAreConvertibleToNumberReturningTrue (0.00s)
=== RUN   TestCheckIfAllValuesAreConvertibleToNumberReturningFalse
--- PASS: TestCheckIfAllValuesAreConvertibleToNumberReturningFalse (0.00s)
=== RUN   TestCheckIfMatrixIsSquareReturningTrue
--- PASS: TestCheckIfMatrixIsSquareReturningTrue (0.00s)
=== RUN   TestCheckIfMatrixIsEmptyReturningTrue
--- PASS: TestCheckIfMatrixIsEmptyReturningTrue (0.00s)
=== RUN   TestCheckIfMatrix1x1IsEmptyReturningFalse
--- PASS: TestCheckIfMatrix1x1IsEmptyReturningFalse (0.00s)
=== RUN   TestCheckIfMatrix2x2IsEmptyReturningFalse
--- PASS: TestCheckIfMatrix2x2IsEmptyReturningFalse (0.00s)
PASS
ok  	manipulate-matrix/app/delivery/http/helperHttp	0.964s
=== RUN   TestAllValuesAreConvertibleToBigIntReturningTrue
--- PASS: TestAllValuesAreConvertibleToBigIntReturningTrue (0.00s)
=== RUN   TestAllValuesAreConvertibleToBigIntReturningFalse
--- PASS: TestAllValuesAreConvertibleToBigIntReturningFalse (0.00s)
=== RUN   TestTheCheckIfMatrixIsSquare1x1
--- PASS: TestTheCheckIfMatrixIsSquare1x1 (0.00s)
=== RUN   TestTheCheckIfMatrixIsSquare3x3
--- PASS: TestTheCheckIfMatrixIsSquare3x3 (0.00s)
=== RUN   TestTheCheckIfMatrixIsSquare5x5
--- PASS: TestTheCheckIfMatrixIsSquare5x5 (0.00s)
=== RUN   TestTheCheckIfMatrixIsSquare3x3WithError
--- PASS: TestTheCheckIfMatrixIsSquare3x3WithError (0.00s)
=== RUN   TestTheCheckIfMatrixIsSquare1x3WithError
--- PASS: TestTheCheckIfMatrixIsSquare1x3WithError (0.00s)
=== RUN   TestTheCheckIfMatrixIsSquare4x2WithError
--- PASS: TestTheCheckIfMatrixIsSquare4x2WithError (0.00s)
=== RUN   TestTheMatrixIsEmptyReturningTrue
--- PASS: TestTheMatrixIsEmptyReturningTrue (0.00s)
=== RUN   TestTheMatrixIsEmptyReturningFalse
--- PASS: TestTheMatrixIsEmptyReturningFalse (0.00s)
=== RUN   TestConvertMatrixToSlice2x2
--- PASS: TestConvertMatrixToSlice2x2 (0.00s)
=== RUN   TestConvertMatrixToSlice3x3
--- PASS: TestConvertMatrixToSlice3x3 (0.00s)
PASS
ok  	manipulate-matrix/app/usecase/matrix/helperMatrix	0.529s
=== RUN   TestCanarySpec
--- PASS: TestCanarySpec (0.00s)
=== RUN   TestStringify1x1Matrix
--- PASS: TestStringify1x1Matrix (0.00s)
=== RUN   TestStringify2x2Matrix
--- PASS: TestStringify2x2Matrix (0.00s)
=== RUN   TestStringify3x3Matrix
--- PASS: TestStringify3x3Matrix (0.00s)
=== RUN   TestInvert1x1Matrix
--- PASS: TestInvert1x1Matrix (0.00s)
=== RUN   TestInvert2x2Matrix
--- PASS: TestInvert2x2Matrix (0.00s)
=== RUN   TestInvert3x3Matrix
--- PASS: TestInvert3x3Matrix (0.00s)
=== RUN   TestFlatten1x1Matrix
--- PASS: TestFlatten1x1Matrix (0.00s)
=== RUN   TestFlatten4x4Matrix
--- PASS: TestFlatten4x4Matrix (0.00s)
=== RUN   TestSum1x1Matrix
--- PASS: TestSum1x1Matrix (0.00s)
=== RUN   TestSum2x2Matrix
--- PASS: TestSum2x2Matrix (0.00s)
=== RUN   TestSum3x3Matrix
--- PASS: TestSum3x3Matrix (0.00s)
=== RUN   TestSum3x3MatrixWithBigInt
--- PASS: TestSum3x3MatrixWithBigInt (0.00s)
=== RUN   TestMultiply1x1Matrix
--- PASS: TestMultiply1x1Matrix (0.00s)
=== RUN   TestMultiply2x2Matrix
--- PASS: TestMultiply2x2Matrix (0.00s)
=== RUN   TestMultiply3x3Matrix
--- PASS: TestMultiply3x3Matrix (0.00s)
=== RUN   TestMultiply2x2MatrixWithBigInt
--- PASS: TestMultiply2x2MatrixWithBigInt (0.00s)
PASS
ok  	manipulate-matrix/app/usecase/matrix/manipulation	0.748s

```
