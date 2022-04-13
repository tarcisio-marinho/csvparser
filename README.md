# CSV Parser
Rain CSV Parser challenge

CSV Parser is an CLI app that parses CSV files and outputs two files, one with the correct data and another with bad data.

the project specification can be found in: [Rain_CSV_Parser_v4.docx](Rain_CSV_Parser_v4.docx)
## Run

```console
go run src/main.go samples/roster1.csv config/full_config.json
```
CSV parser loads the csv file as the first parameter and configurations from the second parameter

## Output
logs are outputed in the terminal:

![](screenshots/logs-success.png)

and the output files are inserted into OUTPUT dir:

![](screenshots/outputdir.png)

original CSV file:

![](screenshots/example3.png)

correct data csv:

![](screenshots/correct-csv.png)

bad data csv:

![](screenshots/bad-csv.png)


## Configuration file
for configuring the required fields (values that will be filtered) and unique fields (not duplicated fields)
they must be declared following the below example:
```json
{
  "RequiredFields": {
    "email": [
      {
        "Name": [
          "email"
        ],
        "MultipleCol": false
      },
      {
        "Name": [
          "e-mail"
        ],
        "MultipleCol": false
      }
    ],
    "id": [
      {
        "Name": [
          "id"
        ],
        "MultipleCol": false
      }],
    "name": [
      {
        "Name": [
          "f.name",
          "l.name"
        ],
        "MultipleCol": true
      },
      {
        "Name": [
          "name"
        ],
        "MultipleCol": false
      }]
  },
  "UniqueFields": [
    "email",
    "id"
  ]
}
```
    
each required field is an key in the "RequiredFields" JSON,
and for each required field, it has an list of possible combinations.
For example, the required field "name", can have "f.name" with "l.name" and "name" 
as possible values in the csv headers, for example

csv:

| l.name |  email   | id  |   f.name  |
| :---: |:---:|:---:|-----|
| marinho | tarcisio_marinho09@hotmail.com | 3 |  tarcisio   |

the parser would get the required field name as: "tarcisio marinho"

Or if the csv only have one field of name:

csv:

| name | id  | email |
| :---: |:---:| :---: |
| tarcisio marinho |  3  | tarcisio_marinho09@hotmail.com |

would also get the name `tarcisio marinho`

As for the unique fields, they are fields that cannot be duplicated.
For example, for the unique fields: ```["email", "id"]```:
    
csv:


| name | id  | email |
| :---: |:---:| :---: |
| tarcisio |  3  | tarcisio_marinho09@hotmail.com |
| joao | 32  | joao@hotmail.com |
| maria |  3  | tarcisio_marinho09@hotmail.com |


the csv has a duplicated email ``tarcisio_marinho09@hotmail.com`` 


## Parser

The parser library can be used regardless the input source, 
it can be used in APIs, web apps, CLIs, and so on.

## Testing
Run unit tests:

```console
go test ./... --cover
```

![](screenshots/tests1.png)

![](screenshots/tests2.png)

    ~$ go test ./... --cover
    ?   	csvparser/src	[no test files]
    ok  	csvparser/src/io	(cached)	coverage: 14.5% of statements
    ok  	csvparser/src/models	(cached)	coverage: 95.0% of statements
    ok  	csvparser/src/parser	(cached)	coverage: 100.0% of statements
    ok  	csvparser/src/utils	(cached)	coverage: 100.0% of statements
    ok  	csvparser/src/validations	(cached)	coverage: 100.0% of statements

## The architecture
I've chosen an this architecture because I wanted this code to be reusable for
APIs, CLIs, web apps and so on. The parser library is agnostic.
Also I wanted it to be configurable, and extensible. Making easy to modify the required and unique fields (simply changing the configuration file)
making easy to change, implement new features searchable fields (since we don't control the input csv's), and it's also easy to deploy new configuration.

## Next steps
I would evolve the code implementing tests for the `io` package
which is 