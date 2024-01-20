* **递送信件:**

```
grpcurl --plaintext 127.0.0.1:2024 list
grpcurl -plaintext -d '{"recipient": "mooon", "letter_body": "hello mooon"}' 127.0.0.1:2024 mooon_mailbox.MooonMailbox/deliver_message
```
