# Copymultiple
This utility is a simple Go application designed to copy a single file to multiple destinations specified in a list file. This can be useful for distributing a file across different directories or for backup purposes.

## Features
- Copies a single source file to multiple destination directories.

## Requirements
Go (Golang) installed on your system.

## Installation
No installation is required. You just need to have Go installed on your system to build the utility.
```
go build
```

## Usage
To use this utility, you need to pass two command-line arguments:
- The path to the source file you want to copy.
- The path to the text file containing a list of destination directories.

`go run . [source_file_path] [destination_list_file_path]`

For example:
```
./copymultiple /path/to/source/file.txt /path/to/destination_list.txt
```
The destination_list_file_path should be a text file with each destination directory on a new line.

## Error Handling
The utility includes basic error handling, which will output errors to the console if any of the file operations fail.

## Limitations
- The utility currently does not support recursive directory copying.
- It does not check if the file already exists at the destination.

## Contributing
Feel free to fork the repository, make improvements, and submit pull requests.

## License
This utility is open-sourced software licensed under the MIT license.