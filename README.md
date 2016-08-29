[![Build Status](https://travis-ci.org/edwinlab/inventory.svg?branch=master)](https://travis-ci.org/edwinlab/inventory)

# Try Golang
User and product management cases. The end goal is to build API endpoint for that case.

## Model & Database
I keep the database schema fairly simple with `goose` to manage migration and `glide` to can easily manage with another developers.

## Install
```
mv .env.tmp .env
```

## API Endpoints
Each description of endpoint will formatted like the following:

```
[HTTP METHOD] [ENDPOINT] - [DESCRIPTION]
```

#### Create a user
This endpoint is used to create a user. 

```
POST /users - Create a user

# curl(1) test, copy & paste this on your terminal
curl -X POST \             
  -H 'Content-Type: application/json' \
  -d '{"name":"John Doe"}' \
  localhost:1323/users
```

#### Get a user
This endpoint is used to get a user. 

```
GET /users/:id - Create a user

# curl(1) test, copy & paste this on your terminal
# change id with one of user ID
curl localhost:1323/users/id
```

#### Update a user
This endpoint is used to update a user. 

```
PUT /users/:id - Update a user

# curl(1) test, copy & paste this on your terminal
# change id with one of user ID
curl -X PUT \             
  -H 'Content-Type: application/json' \
  -d '{"name":"John Doe"}' \
  localhost:1323/users/id
```

#### Delete a user
This endpoint is used to delete a user. 

```
DELETE /users/:id - Delete a user

# curl(1) test, copy & paste this on your terminal
# change id with one of user ID
curl -X DELETE localhost:1323/users/id
```

## To Do
1. Deploy heroku
2. Unit Test
3. ???

## License
BSD 3-clause
