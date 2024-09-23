# Gala Project

This project aim to manage events and guests.
Please keep notice that there are maximum of 10,000 guests per event (otherwise the database data-modeling will be incorrect) and maximum of 1,000 events a guest can register to.

## Data Modeling

Each guest will embeded events
Each events will have embeded guests
We will be using the [`Extended Reference`](https://www.mongodb.com/blog/post/building-with-patterns-the-extended-reference-pattern) pattern by mongodb

### Other ways to model that are not relevant to this application

1. Reference Collection - Only have the id (Object Id) of the other collection and if we need certain fields from the other collection we need to make a `$lookup` operation
2. Subset pattern - better when there is a need to query recent events/guests and each array can grow indefenitly
3. Hybrid - One collection using the `Extended Reference` and the other using the `Subset` pattern. It should be the best approach but requires more consideration.
4. Reference Collection - Have another collection that represent a relation between guest and event - like SQL.