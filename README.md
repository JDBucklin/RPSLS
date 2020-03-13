# RPSLS
This API is used to play games of Rock Paper Scissors Lizard Spock (RPSLS) using the following website: [https://codechallenge.boohma.com/](https://codechallenge.boohma.com/)

More information about RPSLS can be found here: http://www.samkass.com/theories/RPSSL.html

## How to play
Build the project using: `go build`\
Run the project: `./RPSLS`\
By default the API runs on `localhost:8080`\
Navigate to the https://codechallenge.boohma.com/ in your browser\
Once the webpage has loaded, for Step #1 enter http://localhost:8080 (Note that including http:// is necessary for it to work)\
Use the buttons on the page to play.

## Endpoints
The following endpoints are implemented:\
GET `/choices`: Lists all hand choices in the game\
Example Request: `curl http://localhost:8080/choices -X GET`\
Example response:
```json
[
  {
    "id": 1,
    "name": "rock"
  },
  {
    "id": 2,
    "name": "paper"
  },
  {
    "id": 3,
    "name": "scissors"
  },
  {
    "id": 4,
    "name": "lizard"
  },
  {
    "id": 5,
    "name": "spock"
  }
]
```
---
GET `/choice`: Returns a random choice\
Example Request: `curl http://localhost:8080/choice -X GET`\
Example response:
```json
{
  "id": 4,
  "name": "lizard"
}
```
---
POST `/play`: Plays a single game of RPSLS. Expects a payload of the users hand choice.\
Example Payload:
```json
{
  "player": 1
}
```
Example Request: `curl http://localhost:8080/play -X POST -d "{\"player\": 1}"`\
Example Response:
```json
{
  "results": "win",
  "player": 1,
  "computer": 4
}
```
The `results` value can be `win`, `lose`, `tie`