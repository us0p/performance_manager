# Last stop
- Add proper client interaction with the expected input source, Presenters and Views and how all of this ties back with HTMX.
    - How to represent request failures? It's an UI choice that you must make.
    - Since the server must respond with Hypermedia, how do I standardize error responses according to Go error types?
        - Each layer should have a set of errors or error behaviors it should be responsible for. E.g.: validations (business rules), 
        transport like HTTP errors or interactions (framework and drivers).
        - The HTTP layer is what turn them into HTTP codes before sending back to clients.
        - The conversion of errors to HTML must be made in a single place in the Framework and drivers layer.
    - Where should I store page fragments? Should it be with the pages? Or should it be close the routes in which they're used.
        - For errors the user can fix, you should return the same component highlighting the errors so that the user can fix it and try again.
        - For errors the user can't fix, you should log the error and return a generic message to the user, unrelated to the trigger component.
        - In HTMX, by default, only 2xx status codes can produce swaps in the screen. For others you must literally allow swapping for other status codes:
        ```html
            <meta name="htmx-config" content='{"responseHandling":[
                  {"code":"204","swap":false},
                  {"code":"[23]..","swap":true},
                  {"code":"422|409","swap":true},
                  {"code":"500","swap":true},
                  {"code":"[45]..","swap":false,"error":true}
            ]}'>
        ```
        - Fragments must be stored with the templates since it belongs to the view layer. You must name the define block carefully to avoid namespace shadowing. 
        Prefix each partial with feature to avoid conflicts.

## Questions
- What is 'application/x-www-form-urlencoded'?
