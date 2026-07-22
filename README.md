# Hotels Data Merge

To see what this project is about, see [REQUIREMENTS.md](REQUIREMENTS.md)

## Usage

_**TODO**: list instructions and pre-requisites to running this project_

## Deviations from REQUIREMENTS.md

- `GET /api/v1/hotels` endpoint
    - returns 200 OK if no query params are applied
        - follows REST conventions to return all
    - any number of query parameters can be applied (none, 1, both)
- `GET /livez`, `GET /readyz` endpoints
    - quick win to be k8s ready
    - `readyz` is used in Dockerfile healthcheck
