# REST Server

## Tooling

- API design to be documented using OpenAPI v3
    - its industry best practice
    - can be used as source of truth via codegen
        - i already have a scaffold for codegen from a different take-home assignment

## Endpoint Design

### Requirements

while `/hotels` is the bareminimum, `api/v1/` prefix is added to add versioning to api (easier to extend & change behavior in future)
`GET /api/v1/hotels` requires support for 2 query parameters: `hotels` and `destination`

#### Considerations

1. how many query parameters should we allow users to add ?
    - strictly enforce 1 query param (either `hotels` OR `destination` filter)
    - any number of query parameters (none, partial, or all query params?)
        - do filters stack ? or are the results additive ?
1. how should the list of `hotels` be passed ?
    - comma seperated list (e.g.`/hotels?hotels=asdf,qwer`)
    - repeating query params (e.g.`hotels?hotels=asdf&hotels=qwer`)
1. should we consider renaming `hotels` -> `id` and `destination` -> `destination_id`


#### Decision

**TL;DR**: Following REST convention doesn't require much more effort.
Follow it unless it directly contradicts the requirements.

1. allow any number of query parameters
    - users should be able to narrow down what they're searching for
    - users can make multiple request to combine results if needed, no need to pollute design
    - scenarios:
        - 0 query params: return all query params
        - 1 query params: filter according to params given
        - 2 query params: data returned must meet BOTH criteria
1. use comma separated list for users to submit list of hotel ids
    - go against [REST convention](https://stackoverflow.com/a/2602127) to duplicate query param (e.g.`hotels?hotels=asdf&hotels=qwer`)
    - requirements don't allow for changes in URL design; "hotels" (plural) imply that it wants multiple values
        - duplicate query param pattern should rename query param to `hotel` (singular)
1. no renaming query params - requirements only allow rename in response schema
    - renaming the fields `id` -> `hotel` and `destination_id` -> `destination` doesn't read well; keep existing
