# Currency Exchange
This is a mini currency exchange webservice to convert BND to USD

## Usage
1) Open a terminal in the root directory
2) Start the localhost server by running `go run .`
3) Open another terminal
4) Execute a curl command on the `/currency_exchange` endpoint:
```
curl http://localhost:8080/currency_exchange \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"from_currency": "BND", "to_currency": "USD", "amount": 1000}'
```
5) To stop the server, enter `CTRL + C` in the first terminal

### Note
The POST request data can have different value at these fields:
- `amount`: non-negative numbers
- `from_currency`: Any [supported currencies][1] by the 3rd-party API
- `to_currency`: BND, SGD, or USD

[1]: https://currencyscoop.com/supported-currencies