# Directory-Content-Crawler

This Go code is a simple directory content crawler that allows you to get the names and sizes of files in a directory of your choice. This can be useful for quickly identifying the contents of a directory, or for checking the size of individual files.

## Getting Started

To use this code, you'll need to have Go installed on your computer. You can download it from the official Go website.

Once you have Go installed, clone this repository or download the source code and navigate to the project directory in a terminal window. Then, enter the following command:
```Go
go run main.go
```

You'll be prompted to enter the directory path you want to crawl. Once you enter the path, the program will display the names and sizes of all files in that directory.

## How it works

The `getCrawling` function in the code prompts the user to enter a directory path using the bufio package. It then trims any whitespace from the input and changes the working directory to the entered path using the `os` package.

Next, the function uses the `ioutil` package to read the contents of the new working directory. It loops through each file in the directory, printing the name and size of the file using the `fmt` package.

If the directory does not exist, the program will display an error message. If the directory is empty, the program will not display any files.

## Contributing

Feel free to contribute to this project by submitting pull requests or opening issues. Here are some ways you can contribute:

- Improve the user interface or add new features to the program
- Write tests for the program to ensure it works as expected
- Improve the error handling or add more informative error messages
- Translate the program into different languages
- Improve the documentation in this README.md file

## License

This project is licensed under the MIT License, which means you can use, modify, and distribute the code as you wish, as long as you include the original license and copyright information.

## Acknowledgements

This program was inspired by the Directory Tree Crawler program by Andrew Vos.
