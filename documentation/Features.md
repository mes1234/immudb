# Supported features
  * Auth using bearer token
  * Roles 
    * Per stream
      * Reader  
      * Mutator
# Data structure
  * Stream of items
  * Multi indexing
  * Allow not indexed payload - bytes
  * Operations
    * Mutating
      * Add item with ID
      * Remove item with ID - soft delete 
      * Find item with ID
      * Patch item with ID - in reality bump revision but only patched fields changes
      * Update item with ID - in reality bump revision but all fields new
    * Query
      * search for exact match of index/indexes
      * search range of index/es :
        * less than
        * between
        * more than
