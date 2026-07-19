# Development Strategy

_consider creating tickets for these_

## Overall Philosophy

1. start with the simplest solution
    - single process to meet data cleaning criteria
        - no external dependencies - no sqlite / redis / etc
        - keep logic as dumb as possible
        - little to no abstraction
    - let directory structure emerge (if any)
    - test ONLY "public" functions & APIs
1. write docs for plan to extend scale project
    - NOTE: remember to keep project "open to extensions, closed to modification"
1. cleanup
    - run through [REQUIREMENTS.md](REQUIREMENTS.md) again to see if all requirements have been met
    - rewrite README.md; focus on usage instructions
1. bonus material

## Build Order

details for "start with simplest solution"

1. [quick win] model api responses and final schema
1. write merge function to get final schema
1. set up data fetchers
    - initial strategy:
        - fetch all 3 suppliers at the same time at regular intervals
        - proceed with partial data if necessar
            - prioritise availability > consistency here
    - ideas for extension:
        - per supplier config (event driven):
            - start background services for each supplier, they can set their own error handling and poll interval
        - distributed systems
            - takes per supplier services to another level; will require external memory bank to consolidate data
1. set up REST API + openapi spec
