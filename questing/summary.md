## 03/06

- created the algorithm structs. These are the structs that can be used to generate scores and create logic
- model structs. Directly reflect the database and the columns it holds
- loader: will connect the two together so that once the model structs are fed the raw database, they will be able to feed this information into the algorithm structs.

- loader: create maps that will allow us to filter through what each country, destination and activity holds. whether this is the destinations held by country, tags and activities held by a destination or tags held by an activity