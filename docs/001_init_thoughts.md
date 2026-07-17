# Initial Requirements Breakdown

## Deliverables

- documentation
  - how to run project
  - further optimisations
  - how to scale this further
- WEB API with single REST endpoint
  - endpoint must have parameters for filtering by destination and hotel
  - data about hotels are sourced from the provided endpoints in [REQUIREMENTS.md](REQUIREMENTS.md)
- one of these:
  - deployment
  - test pipeline

## Evaluation Criteria

- data quality: how data is selected and cleaned
- solution design
- tests

## Other Considerations

- **KISS / Keep It Simple, Stupid**: previous interview mentioned that company values keeping solutions simple
  - project has short delivery period
- NO MACHINE LEARNING

## Initial Thought / Question Dump

1. how do i identify common properties that are useful to the user?
    1. who is the user anyway? public or analyst? they'll prioritize different things
        - A: looks like public
1. initial ideas for _fetching_ data
    - simple script + cron (kube, serverless, etc)
    - microservice
    - modular monolith
1. initial ideas for _storing_ data
    - in-memory (ephemeral)
        - separate process (e.g. redis)
        - hash map (raw code) / miniredis (library)
    - disk (persistent)
        - same machine (sqlite ?)
        - separate service (postgres, mongodb, etc)
            - how structured is the data we're storing?
    - no storage; fetch on demand
        - can cache responses, but is generally slow and expensive IRL ?
        - we're not going to be making requests for this project, might be worth considering
1. data availability:
    - when does data become stale?
    - do we remove stale data or keep it available?
    - how do i get the cache / persistent storage available before the endpoints are called?
