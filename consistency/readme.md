# consistency levels
- strong consistency
    - def: after a write, any read will see the latest written value
    - example: financial transactions, inventory count
- eventual consistency
    - def: after a write, read may see stale data for some time, but will eventually see the latest written value
    - example: social media posts, listing updates in ecommerce