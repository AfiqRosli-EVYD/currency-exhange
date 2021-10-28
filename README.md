# Currency Exchange

curl http://localhost:8080/currency_exchange \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"from_currency": "BND", "to_currency": "USD", "amount": 1000}'
