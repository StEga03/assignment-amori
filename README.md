## Introduction

This project is a RESTful API service written in Golang designed to manage products and user sessions. It ensures secure access to endpoints through JWT (JSON Web Tokens) and adheres to best practices in code conventions.

## Table of Contents

- [Tech Stack](#tech-stack)
- [Setup](#setup)
- [Manual Setup](#manual-setup)
- [Database Schema](#database-schema)
- [Endpoints](#endpoints)
- [Libraries Used](#libraries-used)
- [Configuration](#configuration)
- [Design Pattern](#design-pattern)
- [Documentation](#documentation)
- [Optimization Considerations](#optimization-considerations)

## Tech Stack

- Golang 1.21
- PostgreSQL
- OpenAI

## Setup

1. Ensure your local PostgreSQL is disabled (or kill it manually), for instance:
   ```bash
   sudo kill -9 PID
   ```
2. Ensure Docker is installed on your machine.
3. Make sure you have filled in the credential or secret value in the .envrc file.
4. Run the following command:
   ```bash
   make run
   ```
5. You're all set!

## Manual Setup

1. Create a database called `assignment-amori` on PostgreSQL.
2. Set up PostgreSQL connection credentials to `postgres` as both the username and password.
3. Make sure you're already set the `.envrc` with your database configuration and the other variable.
4. Execute the command: `make run-http`.
5. You're done!

## Database Schema

There are two tables:

1. **Users**: `(id, first_name, last_name, birth_date, gender, gender_interest, phone_number, relationship_status, relationship_goal)`, with indexing on `phone_number`.
2. **Channels**: `(id, user_id, name)`, with indexing on `user_id`.
3. **Messages**: `(id, channel_id, sender_type, sender_id, content_type, content)`, with indexing on `channel_id`, `sender_id`, `content_type`.
4. **Message Inputs**: `(id, channel_id, source, sender, receiver, receiver_pronoun)`, with indexing on `channel_id`.
5. **Message Sources**: `(id, message_input_id, sender, content_type, content, sent_at)`, with indexing on `message_input_id`, `content_type`, `content`, `sent_at`.

## Endpoints

- **Base URL:** `/api/v1`
- **Base PORT:** `:9000`

1. **Create Channel:** `POST /channels`
    - Request Body:
      ```json
      {
          "name": "Channel Name",
          "messageSource": [
             {
                  "body": "Hey, how was your day?",
                  "sender": "Alice",
                  "sentAt": "2024-06-28T08:00:00"
             },
             {
                  "body": "It was good, thanks for asking! How about yours?",
                  "sender": "Bob",
                  "sentAt": "2024-06-28T08:05:00"
             }
          ]
      }
      ```
    - Response:
      ```json
      {
         "code": 200,
         "message": "Your request processed successfully.",
         "retryable": false,
         "data": 520167476575600641
      }
      ```

2. **Create Message In Channel:** `POST /channels/:id/messages` 
   - URL Parameters:
     - `id` -> Channel ID.
   - Request Body:
     ```json
     {
         "body": "How are you?"
     }
     ```
   - Response: Returns the created message with its ID.
     ```json
     {
        "code": 200,
        "message": "Your request processed successfully.",
        "retryable": false,
        "data": {
           "id": 520167486910365697,
           "channelId": 520167476575600641,
           "body": "How are you?",
           "timestamp": "2024-06-28T18:19:42.810561+07:00"
        }
     }
     ```

3. **List Messages in Channel:** `GET /channels/:id/messages`
    - URL Parameters:
      - `id` -> Channel ID.
    - Response:
      ```json
      {
         "code": 200,
         "message": "Your request processed successfully.",
         "retryable": false,
         "data": [
            {
               "id": 520159085702676481,
               "channelId": 520159024281288705,
               "body": "I'm just a computer program, so I don't have feelings, but thanks for asking! How can I assist you today?",
               "timestamp": "2024-06-28T09:56:15.303344Z"
            },
            {
               "id": 520159082733109249,
               "channelId": 520159024281288705,
               "body": "How are you?",
               "timestamp": "2024-06-28T09:56:13.533594Z"
            }
         ]
      }
      ```

4. **File Parser:** `POST /files/sources/:platformType`
    - Form Data:
      - `file` -> File to be uploaded.
    - Response:
      ```json
      {
         "code": 200,
         "message": "Your request processed successfully.",
         "retryable": false,
         "data": [
            {
               "sender": "Alice",
               "content": "Hey, how was your day?",
               "contentType": "text",
               "sentAt": "2024-06-28T08:00:00Z"
            },
            {
               "sender": "Bob",
               "content": "It was good, thanks for asking! How about yours?",
               "contentType": "text",
               "sentAt": "2024-06-28T08:05:00Z"
            }
         ]
      }
      ```


## Libraries Used

- `gomock` for unit testing.
- `chi` for lightweight routing.
- `uuid` to generate unique strings based on time.
- `sonyflake` to generate unique numbers based on time.
- `pmx` as a driver for PostgreSQL.
- `httprest` for rate limiting.
- `makefile` for base setup.
- `openai` as a OpenAI SDK.

## Configuration

- Configuration is managed using a .envrc file format to define all layers of the application (DB, port, openai, etc.).

## Design Pattern

- The application uses the repository pattern to handle and separate code into various layers, including handler, usecase, and repository levels, which then interact with the database.
- The handler sanitizes the input, the usecase executes the main logic and returns data, and the repository manages data interactions with the datastore.

## Documentation

Find all the documentation in the `/docs` directory.

### Postman Collection

- Access the Postman collection [here](https://api.postman.com/collections/12053329-fddf68f9-bfb7-4a86-9ae5-2ed29c3bca83?access_key=PMAT-01J1FKPQX6BB0NKP90EA4N1CS6) to interact with the API's endpoints.


## Optimization Considerations

- Optimizations include implementing caching strategies to reduce database load, optimizing database queries to improve response times, and considering horizontal scaling to accommodate growth in user traffic and data volume.
