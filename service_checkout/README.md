# service_checkout

service_name: service_checkout service_port: 9003

> 身份可看做是一个bff层

调用链：

1. checkout -> email
2. checkout -> payment
3. checkout -> shipping
4. checkout -> currency
5. checkout -> productCatalog
6. checkout -> cart
7. gateway -> checkout