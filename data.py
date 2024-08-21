import pands as pd

#Load the Excel file
Excel_file = 'fyb_data.xlsx'
sheet_name = 'sheet1'

#Read the excel file
df =pd.read_excel(Excel_file,sheet_name=sheet_name)

#save the data as woord
output_file='output.txt'
df.to_csv(output_file,sep='\t',index=false)

print(f"Data successfully extracted to{output_file}")