# Data Cleaning

## Data Flow

(consider drawing a flow chart if there's time)

1. fetch raw bytes from REST APIs
1. parse bytes / unmarshal bytes to supplier-specific Go object (e.g. ACME)
    - avoid business and cleaning logic at this level
    - supplier-specific data should be left as raw as possible
1. normalize all supplier-specific objects to project-specific `Hotel` object
1. group information about `Hotel`s with each other
    - GROUP BY hotel.id
    - i.e. split the big array of hotels into smaller ones; 1 id should have an array of 1 or more hotels
1. for each hotel id, merge all the information we have about it (i.e. Hotel[] -> Hotel)
    - this is where most of the business logic will live
        - e.g. take longest description
        - e.g. trim whitespace on both ends
        - e.g. convert to lower case
        - etc
1. the final data structure (hashmap<key=hotel.id, val=hotel object>) will be used in some writer / setter function on a Repostory class
    - consider denormalization where we pre-compute a 2nd copy that's grouped by _destination_id_
        - this is a quick win to allow fetching records in O(1) time rather than O(n) - hashmap vs array lookup
    - Repo layer hides how data is stored, will make it more extensible should we change to something like kv store / sql / nosql / etc.
1. the Web API / REST server will contain a reference to the Repository class via Dependency Injection

## TL;DR

- supplier-specific models (AcmeHotel, etc)
  - contains "raw" data
- project-specific model(s) (Hotel)
  - contains "cleaned" data
- data cleaning happens during normalization (e.g. AcmeHotel -> Hotel)
- data merging happens during collation (e.g. []Hotel -> Hotel)
