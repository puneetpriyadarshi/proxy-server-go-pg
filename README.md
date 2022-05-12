# TODO-GIN-GONIC

# PROXY SERVER DEVELOPMENT

### Description :

- Created a proxy server using gin-gonic,postgres,httputil.
- For getting the credentials for the database connection , check configs/db.go file.

### GIVEN TASKS:-

- #### Task No 1 : Create a new tenant

  - Created a new tenant in the postgres database using go-pg/go for connection and pgadmin4 for handling the database.
  - ##### API ENDPOINTS :

    `[GET] http://localhost:5000/tenant`

    Hitting this URL will fetch all the existing tenants in the database.

    `[GET] http://localhost:5000/tenant/id`

    Hitting this URL will fetch the required single tenant by id in the database.

    `[POST] http://localhost:5000/tenant`

    {
    name : "your-name"
    }

    Hitting this URL will create a new tenant in the database with the above mentioned properties.

- #### Task No 2 : Creating Proxy server

  - A proxy server is created using the httputil package which does the following things :
  - API ENDPOINTS :

    `[GET] http://localhost:5000/metrics/*`

    When _ is a number the proxy server redirects to the page https://httpstat.us/_. In response we get the body in the postman or console and also the bytes transferred in the console log.We also get the error count if we are unable to get the required response.

        ```[POST] http://localhost:5000/metrics/*```

    When \* is a number , then in postman we get the required message along with the bytes transferred in the console.We also get the error count if we are unable to get the required response.

         ```[GET] http://localhost:5000/metrics/*```

    When _ is a string and that string is either "photos" or "albums" or "comments" or "posts" the proxy server redirects to the page https://jsonplaceholder.typicode.com/_. In response we get the body in the postman or console and also the bytes transferred in the console log.We also get the error count if we are unable to get the required response.

        ```[POST] http://localhost:5000/metrics/*```

    When \* is a string and its either "posts" or "photos" or "albums" or "comments" , then in postman we get the required message along with the bytes transferred and error count if any in the console.

- #### Task no 3 : Fetch Usage metrics

  - Successfully created a tenant activity table with the below mentioned properties :

             id                bigserial    not null primary key,
             tenant_id         bigint       not null,
             created_at          timestamp,
             api_end_point     varchar(100) not null,
             bytes_transferred bigint       not null,
             is_error smallint not null, // only on 404
             CONSTRAINT tentant_activity_fk
                    FOREIGN KEY (tenant_id)
                          REFERENCES tenants (id)

  - I was successfully able to transfer the value of the error_count and bytes_transferred to my database , update them and also them.
    `[GET] http://localhost:5000/tenantactivity/id` = = = > Get the tenant activity by their id.
