## Consistent hash

### Structure

#### Hasher
- `hash_to_used(data)`: Method that returns a hashed value

- `Hash function used`: 
xxhash.Sum64(data) % 1024


#### Consistent
- `hash_function`: Hasher
- `ring`: Map of Hash_value to String
- `sorted_keys`: List of sorted hash_value
- `replication_factor`: Integer for vitual node
- `mapping`: Map of String to String

### Methods

#### Initialize(replication_factor)
- Set `hash_function` to a new instance of Hasher
- Initialize `ring` as an empty map
- Initialize `sorted_keys` as an empty list
- Set `replication_factor` to the given replication_factor
- Initialize `mapping` as an empty map

#### AddServer(server)
- For `i` from 0 to `replication_factor - 1` 
  - Create a key by combining `server` and `i` (Create vitual node for each server)
  - Hash the key using `hash_function`
  - Add the hash and server to the `ring`
  - Add the hash to `sorted_keys`
- Sort `sorted_keys` in ascending order
- If `mapping` is not empty
  - Redistribute keys to the new server

#### AddKey(key)
- Find the appropriate server using `MapKey` method
- Add the key to the `mapping` with the server

#### DelServer(server)
- For `i` from 0 to `replication_factor - 1`
  - Create a key by combining `server` and `i`
  - Hash the key using `hash_function`
  - Remove the hash and server from the `ring`
  - Remove the hash from `sorted_keys`
- If `mapping` is not empty
  - Redistribute keys from the removed server

#### DelKey(key)
- Simply delete the key and value in `mapping`

#### MapKey(key)
- Hash the key using `hash_function`
- Find the next highest hash in `sorted_keys` that is greater than or equal to the hash (The exact rules can be discussed)
- Return the server associated with the hash in the `ring`

#### RedistributeKeys(server, adding)
- For each key in `mapping`
  - If `adding` 
    - For all the key, do function `Mapkey`
    - If the server returned by `MapKey` for the current key equals server, add the key to `mapping`
  - Else if not `adding` 
    - Delete the server and its hash value in `ring` and `sorted_keys`
    - Do function `Mapkey` to all server-related key in `mapping`
    - Due to the returned server, add new pair (key, server) to  `mapping`

## Golang version of the code
- https://github.com/nicottttt/consistent_hashing/tree/main/consistent

## Data and result for testing

- Key "key2222" 's value is 596, and it is mapped to:Server4
- Key "key222222" 's value is 112, and it is mapped to:Server2