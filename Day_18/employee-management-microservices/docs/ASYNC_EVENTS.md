# Asynchronous Communication

Kafka and separate message brokers are not used.

1. Employee Service saves employee data.
2. Employee Service saves the event in its outbox table
   within the same database transaction.
3. A Go background worker reads pending events.
4. The worker sends events to an HTTP receiver.
5. Successfully delivered events are marked as published.

Consumers should persist event IDs and ignore duplicate events.