# Data Fetching

## Hotel Data Properties

1. largely unstructured
1. unlikely to change
    - hotels are unlikely to change address, rename or add new facilities out of the blue
1. don't require high precision
    - e.g. lat / lng are likely to be estimates anyway
    - customers can likely accept some degree of staleness

## CAP Theorem

given the above properties, we can safely prioritise having information available over being extremely correct.

## Initial Strategy - fetch all data at the same time, at regular intervals

this strategy was chosen because of its simplicity:
    - no need to piece together data that arrives at different times
    - less configuration required for intervals
    - most errors can be ignored - if a supplier fails to return data, we can probably afford to fetch it at the next interval
        - (see section on CAP Theorem)
