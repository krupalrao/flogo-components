# transformXmlToJson
This activity provides your Flogo application the ability to run transformation on XML input strings  to JSON.

## Installation

```
flogo install github.com/krupalrao/flogo-components/activity/transformXmlToJson
```

Link for flogo web:

```
https://github.com/krupalrao/flogo-components/activity/transformXmlToJson
```


## Schema
Inputs and Outputs:

```json
{
  "inputs": [
    {
      "name": "input",
      "type": "any",
      "required": true
    },
    {
      "name": "spec",
      "type": "any"
      
    }
  ],
  "outputs": [
    {
      "name": "output",
      "type": "any"
    }
  ]
}

```

## Inputs
| Setting     | Required | Description    |
|:------------|:---------|:---------------|
| content     | True | Input XML string to transform |
| spec |  | Kazaam transformation specification   |

## Outputs
| Setting     | Description    |
|:------------|:---------------|
| result | Transformed JSON string

