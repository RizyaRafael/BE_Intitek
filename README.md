# BE_Intitek

## Models :

_Products_
- sku: string, required
- quantity: integer, required
- location: string, required
- status: string, required

_Users_
- email: string, required
- password: string, required

## Endpoints :
List of available endpoints:
- `POST /user/register`
- `POST /user/login`

Routes below need authentication:
- `GET /product`
- `POST /product`
- `PUT /product`
- `DELETE /product`

&nbsp;
## 1. POST /user/register
Description:
- User Register
Request:
- body:
```json
{
  "email": "string",
  "password": "string"
}
```
_Response (201 - Created)_
```json
{
  "message": "string",
}
```
_Response (400 - Bad Request)_
```json
{
  "message": "Email already registerd",
}
OR
{
  "message": "All data is required",
}
```

&nbsp;

## 2. POST /user/login
Description:
- Login user

Request:
- body:
```json
{
  "email": "string",
  "password": "string"
}
```

_Response (200 - OK)_
```json
{
  "access_token": "bearer <token>"
}
```

_Response (400 - Bad Request)_
```json
{
  "message": "Email already registerd",
}
OR
{
  "message": "Email and password required",
}
OR
{
  "message": "All data is required",
}
```

_Response (401 - Unauthorized)_
```json
{
  "message": "Invalid email or password"
}
```

&nbsp;

## 3. GET /product
Description:
- Get all product

Request:

- headers:
```json
{
  "access_token": "bearer <token>"
}
```

- params:
```json
{
  "sort": "string"
}
OR
{
  "status": "string"
}
```

_Response (200 - OK)_
```json
[
    {
    "sku": "string",
    "quantity": "integer",
    "location": "string",
    "status": "string"
    }
]
```
&nbsp;

## 4. POST /product
Description:
- Create a product

Request:

- headers:
```json
{
    "access_token": "bearer <token>"
}
```

- body:
```json
{
    "sku": "string",
    "quantity": "integer",
    "location": "string",
    "status": "string"
}
```

_Response (201 - OK)_
```json
{
    "message" : "You've succesfully created the product"
}
```

_Response (400 - BAD REQUEST)_
```json
{
    "message" : "Database error"
}
```
&nbsp;

## 4. PUT /product
Description:
- Update a product

Request:
- headers:
```json
{
    "access_token": "bearer <token>"
}
```

- body:
```json
{
    "id": "integer",
    "sku": "string",
    "quantity": "integer",
    "location": "string",
    "status": "string"
}
```

_Response (201 - Created)_
```json
{
    "message" : "Product succesfully updated"
}
```

_Response (404 - NOT FOUND)_
```json
{
    "message" : "Data not found"
}
```
&nbsp;

## 5. DELETE /product
Description:
- Delete a product

Request:
- headers:
```json
{
    "access_token": "bearer <token>"
}
```
- params
```json
{
  "id": "integer"
}
```

_Response (200 - OK)_
```json
{
    "message" : "data succesfully deleted"
}
```

## Global Errror

_Response (401 - Unauthorized)_

```json
{
  "message": "Please login"
}
```

_Response (500 - Internal Server Error)_

```json
{
  "message": "an error occured"
}
``` 