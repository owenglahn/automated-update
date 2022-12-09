# automated-update

## Running tests

- To run all unit tests:

  `$ go test -v`

- To run a specific test:

  `$ go test -v -run {Name of test}`

## Documentation

### Configuration

- Set the configuration for the API in the config.yml file.
  - `API_BASE_URL` is the host for the API
  - `API_TOKEN` is the auth token for accessing the API, which should be a system env variable that is updated whenever the token changes.

### Data Structures

`` type Application struct {
    ApplicationId string `json:"applicationId"`
    Version       string `json:"version"`
  } ``

- Defines an application and it's version on a particular player.

`` type Profile struct {
    Applications []Application `json:"applications"`
  } ``

- Defines versions for a collection of application

`type Client struct {
    baseUrl string
  }`

- Holds data for sending API requests

### Functions

- `UpdateFromCSV(pathToCSV string, profile Profile)`
  - Makes a PUT request with `profile` as the request body for all players with MAC Addresses found in the csv file `pathToCSV`
- `UpdateVersion(macAddress string, profile Profile)`
  - Makes a PUT request with body `profile` for the player associated with `macAddress`
