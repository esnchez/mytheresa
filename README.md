
## MyTheresa catalog application

Minimal Go application that provides a REST API endpoint to fetch products from a PostgreSQL database. 
Filters [category, priceLessThan] have been implemented. 

Other remarkable considerations taken:

- Graceful shutdown of the server
- Pagination with limit and offset
- Timeouts on database query
- Retry on database connection
- Discounts are loaded in memory into a map for faster lookups
- Populating database with Postgres init scripts
- Minimize external dependencies, leverage standard library

From the root folder, 

Run all tests with:

```
make test
```

Build up and run the system:

```
docker compose up --build
```

Curl the following urls to get different results:

```
/products 
/products?category=boots
/products?priceLessThan=80000
/products?limit=4&offset=1
/products?limit=3
/products?offset=3
```
