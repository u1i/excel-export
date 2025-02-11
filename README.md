# Excel to CSV Converter

A tool for converting Excel workbooks into separate CSV files, with one CSV file per worksheet. Available in both Python and Go implementations. The tool automatically handles line breaks in cells by replacing them with spaces, ensuring clean CSV output.

## Features

- Converts each Excel worksheet to a separate CSV file
- Automatically cleans line breaks from cells
- Preserves worksheet names as CSV filenames
- Uses UTF-8 encoding for proper character handling
- Excludes DataFrame indices from output
- Available in both Python and Go implementations
- Command-line interface for easy use

## Requirements

### Python Version
- Python 3.6 or higher
- pandas library
```bash
pip install -r requirements.txt
```

### Go Version
- Go 1.16 or higher
- excelize package (automatically installed via go.mod)
```bash
go mod tidy
```

## Installation

### Python Version
1. Clone the repository
2. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

### Go Version
1. Clone the repository
2. Build the executable:
   ```bash
   go build xls2csv.go
   ```

## Usage

Both versions accept an Excel file path as a command-line argument:

### Python Version
```bash
python xls2csv.py <excel_file>
```

### Go Version
```bash
./xls2csv <excel_file>
```

Example:
```bash
# Using Python version
python xls2csv.py quarterly_report.xlsx

# Using Go version
./xls2csv quarterly_report.xlsx
```

## Output

Each worksheet will be exported as a separate CSV file with:
- UTF-8 encoding
- No index column
- Line breaks replaced with spaces
- Filename matching the worksheet name

For example, if your Excel file `quarterly_report.xlsx` has worksheets named "Q1" and "Q2", you'll get:
- `Q1.csv`
- `Q2.csv`

## Error Handling

Both versions include error handling for common issues:
- File not found
- Invalid Excel file format
- Permission issues
- Invalid command-line arguments

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the MIT License.
