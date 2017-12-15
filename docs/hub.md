messaging
=========

- APIs
    - levit.services.users-service
    - levit.services.pail-service
    - levit.services.hub-service
- Clients
    - Websocket
    - gRPC
    - HTTP Long-poll (not urgent)
- Sent/Delivered/Read
    - sent - client -> service
    - delivered - service -> DB
    - read - last_active in chat vs delivered time
- GCM
    - service that receives message sends gcm
    - don't send to users active in channel
- Broker
    - distribute all clients in current hub
    - distribute to broker on channel /hub/{channelID}
- User info
    - locally cache user info
        - first/last name, profile pic
        - expire time of 5 mins
- ScyllaDB
    - https://blog.discordapp.com/how-discord-stores-billions-of-messages-7fa6ec7ee4c7
    ```cql
    CREATE TABLE messages (
      channel_id bigint,
      message_id bigint,
      from bigint,
      content text,
      timestamp uint,
      PRIMARY KEY (channel_id, message_id)
    ) WITH CLUSTERING ORDER BY (message_id DESC);
    ```
- Cache pls