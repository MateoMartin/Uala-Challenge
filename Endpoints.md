# API Documentation

This document describes the available endpoints, their parameters, request bodies, and possible responses.

---

## 1. Service Status

**Method:** `GET`  
**Path:** `/status`  
**Summary:** Checks if the service is running properly.  
**Description:** Indicates that the service has started up correctly and is ready to accept requests.

### Responses
| Status | Description                |
|--------|----------------------------|
| 200    | Service is running (OK).  |
| 500    | Internal Server Error.    |

---

## 2. Create Tweet

**Method:** `POST`  
**Path:** `/tweet`  
**Summary:** Create a new tweet.  
**Description:** Accepts a JSON body with user information and tweet content.

### Request Body
```json
{
  "user_id": "string (required)",
  "content": "string (required)"
}
```

* user_id: Unique identifier for the user posting the tweet (e.g., UUID).
* content: Tweet text (max 280 characters).

### Responses
| Status | Description                |
|--------|----------------------------|
| 201    | Tweet created successfully.|
| 400    | Bad Request|
| 500    | Internal Server Error.    |

## 3. Follow User

**Method:** `POST`  
**Path:** `/follow`  
**Summary:** Follow another user.

**Description:** Accepts a JSON body specifying which user will follow which other user.

### Request Body
```json
{
  "user_id": "string (required)",
  "user_id_to_follow": "string (required)"
}
```

* user_id: The user initiating the follow.
* user_id_to_follow: The user to be followed.

### Responses
| Status | Description                |
|--------|----------------------------|
| 202    | Follow request accepted. |
| 400    | Bad Request|
| 500    | Internal Server Error.    |

## 4. Get Timeline

**Method:** `GET`  
**Path:** `/timeline/{user_id}`  
**Summary:** Retrieve a user’s timeline.
**Description:** Returns all tweets from the specified user and the users they follow.

### Request Body
```json
{
  "user_id": "string (required)",
  "user_id_to_follow": "string (required)"
}
```

* user_id: The user whose timeline is requested.

### Responses
| Status | Description                |
|--------|----------------------------|
| 200    | Returns a JSON array of tweets from the user and followed accounts. |
| 500    | Internal Server Error.    |

