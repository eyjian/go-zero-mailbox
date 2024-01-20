* **查表：**

```sql
SELECT f_id,f_recipient,f_deliver_time,f_arrival_time,f_state,f_letter_body FROM t_mooon_mailbox WHERE f_recipient='mooon' AND f_state=0 LIMIT 10;
```

* **递送信件：**

```
# grpcurl --plaintext 127.0.0.1:2024 list
for ((i=0; i<100; ++i)) do \
	grpcurl -plaintext -d "{\"recipient\": \"mooon\", \"letter_body\": \"hello mooon-$i\"}" 127.0.0.1:2024 mooon_mailbox.MooonMailbox/deliver_message; \
done
```

* **检索信件：**

```
# grpcurl -plaintext -d '{"recipient": "mooon", "start_letter_id": "", "page_size": 10, "list_action": 0}' 127.0.0.1:2024 mooon_mailbox.MooonMailbox/list_messages
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

* **标记为已读：**

```
grpcurl -plaintext -d '{"recipient": "mooon", "letters_id_list": [1,2]}' 127.0.0.1:2024 mooon_mailbox.MooonMailbox/mark_messages_as_read
```

* **删除信件：**

```
grpcurl -plaintext -d '{"recipient": "mooon", "letters_id_list": [1,2]}' 127.0.0.1:2024 mooon_mailbox.MooonMailbox/delete_messages
```
