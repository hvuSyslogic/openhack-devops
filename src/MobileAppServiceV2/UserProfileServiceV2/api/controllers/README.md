# HEALTHCHECK ROUTE

## Overview

Healthcheck route will return json indicating service is healthy.

## healthcheck [HTTP GET]

### Example Request

```html
curl http://127.0.0.1:8080/api/healthcheck
```

### Example Response

```json
{
    "message": "healthcheck",
    "status": "healthy"
}
```

# USERPROFILES ROUTE

CRUD operations on user profiles used by my driving application.

## Get All Profiles [HTTP GET]
Retrieve all profiles from MyDriving Db

### Example Request

```html
curl http://127.0.0.1:8080/api/userprofiles/
```

### Example Response

```json
{
    "user_profiles": [
        {
            "Id": "12345",
            "FirstName": "Richard",
            "LastName": "Guthrie",
            "UserId": "67890",
            "ProfilePictureUri": "http://whoknows.com",
            "Rating": 1,
            "Ranking": 1,
            "TotalDistance": 116,
            "TotalTrips": 5,
            "TotalTime": 26,
            "HardStops": 0,
            "HardAccelerations": 0,
            "FuelConsumption": 2.4,
            "MaxSpeed": 70,
            "Version": "AAAAAAAAB9o=",
            "CreatedAt": "2018-04-16T17:33:40.9366667Z",
            "UpdatedAt": "2018-04-16T17:33:40.9541356Z",
            "Deleted": false
        },
        {
            "Id": "123456",
            "FirstName": "Richard",
            "LastName": "Guthrie",
            "UserId": "67890",
            "ProfilePictureUri": "http://whoknows.com",
            "Rating": 1,
            "Ranking": 1,
            "TotalDistance": 116,
            "TotalTrips": 5,
            "TotalTime": 26,
            "HardStops": 0,
            "HardAccelerations": 0,
            "FuelConsumption": 2.4,
            "MaxSpeed": 70,
            "Version": "AAAAAAAAB/A=",
            "CreatedAt": "2018-04-16T17:33:40.9366667Z",
            "UpdatedAt": "2018-04-18T19:23:03.5617471Z",
            "Deleted": false
        }
    ]
}
```

## Get Profile by ID [HTTP GET]

Retrieve profile by ID

### Example Request

```html
curl http://127.0.0.1:8080/api/userprofiles/12345
```

### Example Response

```json
{
    "Id": "12345",
    "FirstName": "Richard",
    "LastName": "Guthrie",
    "UserId": "67890",
    "ProfilePictureUri": "http://whoknows.com",
    "Rating": 1,
    "Ranking": 1,
    "TotalDistance": 116,
    "TotalTrips": 5,
    "TotalTime": 26,
    "HardStops": 0,
    "HardAccelerations": 0,
    "FuelConsumption": 2.4,
    "MaxSpeed": 70,
    "Version": "AAAAAAAAB9o=",
    "CreatedAt": "2018-04-16T17:33:40.9366667Z",
    "UpdatedAt": "2018-04-16T17:33:40.9541356Z",
    "Deleted": false
}
```

## Create User Profile [HTTP POST]

### Example Request

```bash
curl -F POST http://127.0.0.1:8080/api/userprofiles
```