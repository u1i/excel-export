import pandas as pd

# Function to read Excel file and write CSV files for each sheet
def excel_to_csv(file_path):
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

# Path to your Excel file
file_path = 'in.xlsx'

# Call the function to convert Excel to CSV
excel_to_csv(file_path)

