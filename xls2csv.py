import pandas as pd
import sys
import os

def excel_to_csv(file_path):
    """
    Convert Excel file to CSV files, one per worksheet.
    Replaces line breaks with spaces in cell contents.
    
    Args:
        file_path (str): Path to the Excel file
    """
    # Check if file exists
    if not os.path.exists(file_path):
        print(f"Error: File '{file_path}' not found")
        sys.exit(1)
        
    # Read the Excel file
    excel_data = pd.ExcelFile(file_path)
    
    # Iterate through each sheet and clean data
    for sheet_name in excel_data.sheet_names:
        df = pd.read_excel(excel_data, sheet_name=sheet_name)
        
        # Replace line breaks with space
        df = df.replace({r'[\r\n]+': ' '}, regex=True)
        
        # Export the cleaned data to CSV
        csv_file_path = f"{sheet_name}.csv"
        df.to_csv(csv_file_path, index=False, encoding='utf-8')
        print(f"Worksheet '{sheet_name}' has been exported to '{csv_file_path}'.")

def main():
    if len(sys.argv) != 2:
        print("Usage: python xls2csv.py <excel_file>")
        sys.exit(1)
    
    file_path = sys.argv[1]
    excel_to_csv(file_path)

if __name__ == "__main__":
    main()

