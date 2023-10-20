# Proof-of-concept race result parser

This project is the result of a short discussion about privacy and the results of sports events. It shows that results of the race time provider "race result" can easily be parsed and analyzed.

For this case I created a little go program that parses the results of an event and converts it to CSV. So its not about the CSV itself. Its more about getting a feel for what is possible.

This project will probably don't work for any other events than the event with the id `246735`. But it wouldn't take a lot of more time to get it work for any other event as well. Even the parsing of multiple events should be possible.

## Requirements

- Go version >= 1.21.0

## Setup / Running

1. Clone this repo
2. Run `go run . 246735` from the main directory of this project. Where the number is the race result specific race id.

## Disclaimer

This project is not intended to be used for any public reasons nor parsing the race result page with huge traffic. There are dedicated APIs from race result for this! It is just a way to find out what is possible with the public data of the page. Please check the race result page for there terms of use or get in contact with them before creating any kind of race result parser!