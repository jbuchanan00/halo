# Go API for Postgres Location Management

This Go API allows you to input location data into a PostgreSQL database and provides two key API endpoints for location-related queries.

## Features

1. **Input JSON Data into PostgreSQL:**
   The API includes a JSON file, which is parsed and input into a PostgreSQL database.
   Run 'go run main.go' from 'halo/server/cmd/adhoc/processGeoData'

3. **Two API Routes:**
   - **GET /withinradius:** This route retrieves locations within a specified radius of a given latitude and longitude. Params are lng: float, lat: float, and radius: int
   - **POST /autofill:** This route is used for location-based autofill. You send a location string and the base location's information, and the API will process and return relevant data.
   -   Body: {
        "location": string, //location from input that is being searched
        "baseLoc": {
          "id": int,
          "name": string,
          "state": string,
          "country": string,
          "coords": {
            "lat": float,
            "lng": float
          }
        }
      }

