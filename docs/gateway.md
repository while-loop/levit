gateway
=======

- APIs
    - levit.services.users-service
    - levit.services.auth-service
 ```cql
    CREATE TABLE sessions (
              id,
              user_id
              token
    );
    ```
- Auth user
    - send auth info in JWT with expire time of 1 min
    - send request UUID
    - user_id, scopes, etc..