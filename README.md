# CSV Tool

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Alvaroalonsobabbel/csvtool) ![Test](https://github.com/Alvaroalonsobabbel/csvtool/actions/workflows/go-test.yml/badge.svg)

Updates CSV files with values from another CSV file and generates an output file `out.csv`

## Installation

### MacOS

Download the latest release [here](https://github.com/Alvaroalonsobabbel/csvtool/releases/latest/download/csvtool) and install it.

### non MacOS

You have to install Go and compile the version for your OS.

## Usage

`csvupdate -compareby=<field_name> -update=<field_name> outdated_file.csv updated_file.csv`

Flags:

- `-compareby` indicates the field use to compare both CSV files.
- `-update` indicates the fields to update. You can indicate more than one field by separating them with commas: `-update=FIELD1,FIELD2`
- `-help` will display the help information.
  
Arguments:

The app will expect you to provide both the CSV missing information (outdated) and the CSV that contains the new information (updated). The outdated CSV have to be listed first.

Considerations:

- Listed fields passed in `compareby` and `update` flags must exist in both CSVs. Otherwise the application will rise an error.
- Output file will be `out.csv`

## Example

Comparing CSVs using the **UUID** field and updating the **RATING** and **COMMENTS** fields.

```cmd
csvupdate -compareby=UUID -update=RATING,COMMENTS outdated.csv updated.csv
```
