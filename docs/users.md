users
=====

- API
    - levit.services.users-service
    - levit.services.pail-service
 ```cql
    CREATE TABLE users (
      id bigint,
      first_name
      last_name
      created_at
      timestamp uint,
      deleted bool
      google_id
      facebook_id
      PRIMARY KEY (id)
    );
    ```