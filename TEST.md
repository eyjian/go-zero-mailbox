* **递送信件:**

```
grpcurl --plaintext 127.0.0.1:2024 list
grpcurl -plaintext -d '{"recipient": "mooon", "letter_body": "hello mooon"}' 127.0.0.1:2024 mooon_mailbox.MooonMailbox/deliver_message
```

* **检索信件：**

```
# grpcurl -plaintext -d '{"recipient": "mooon", "page_start": 0, "page_size": 10, "list_action": 0}' 127.0.0.1:2024 mooon_mailbox.MooonMailbox/list_messages
{
  "recipient": "mooon",
  "letters": [
    {
      "letterId": "3",
      "deliverTime": "1705734339",
      "arrivalTime": "1705734339",
      "letterBody": "hello mooon world"
    },
    {
      "letterId": "2",
      "deliverTime": "1705730808",
      "arrivalTime": "1705730808",
      "letterBody": "hello mooon world"
    },
    {
      "letterId": "1",
      "deliverTime": "1705730776",
      "arrivalTime": "1705730776",
      "letterBody": "hello mooon"
    }
  ],
  "nextPageStart": "3"
}
```
