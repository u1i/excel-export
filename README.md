# Excel to CSV Converter

A simple Python script that converts Excel workbooks into separate CSV files, one for each worksheet. The script automatically handles line breaks in cells by replacing them with spaces, ensuring clean CSV output.

## Features

- Converts each Excel worksheet to a separate CSV file
- Automatically cleans line breaks from cells
- Preserves worksheet names as CSV filenames
- Uses UTF-8 encoding for proper character handling
- Excludes DataFrame indices from output

## Requirements

```
pandas
```

## Installation

1. Ensure you have Python installed on your system
2. Install the required package:
   ```bash
   pip install -r requirements.txt
   ```

## Usage

1. Place your Excel file in the project directory
2. Update the `file_path` in `xls2csv.py` if your Excel file has a different name than `in.xlsx`
3. Run the script:
   ```bash
   python xls2csv.py
   ```

The script will create one CSV file for each worksheet in your Excel file, named after the respective worksheet.

## Output

Each worksheet will be exported as a separate CSV file with:
- UTF-8 encoding
- No index column
- Line breaks replaced with spaces
- Filename matching the worksheet name

For example, if your Excel file has worksheets named "Europe" and "APAC", you'll get:
- `Europe.csv`
- `APAC.csv`
