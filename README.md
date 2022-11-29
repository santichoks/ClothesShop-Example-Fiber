# ClothesShop-Example-Practice
<h3 align="center">🚨🚨🚨 this document is in progress 🚨🚨🚨</h3>
<details><summary><h3>API Endpoints</h3></summary>
<p>
<ul>
<li>

<p><strong>Get Products</strong></p>
 
Get all products `http://localhost:3000/products`
 
|Endpoint|Method|Optional Params|Example|
|:-:|:-:|-|-|
|`/products`|GET|`gender [Men, Women]`|`http://localhost:8000/products?gender=Women`|
||||`http://localhost:8000/products?gender=Men,Women`|
|||`size [XS, S, M, L, XL]`|`http://localhost:8000/products?size=M`|
||||`http://localhost:8000/products?size=XS,S,M`|
|||`style [Red, Black, Batman, Spiderman]`|`http://localhost:8000/products?style=Red`|
||||`http://localhost:8000/products?style=Red,Spiderman,Batman`|

<strong>Example :</strong> `http://localhost:8000/products?gender=Women&size=XS,S,M`

```
{
    "status": "successfully.",
    "total": 9,
    "results": [
        {
            "product_id": 6,
            "gender": "Women",
            "style": "Plain color / Black",
            "size": "XS",
            "price": 290
        },
        {
            "product_id": 7,
            "gender": "Women",
            "style": "Plain color / Black",
            "size": "S",
            "price": 290
        },
        ...
        ...
        ...
    ]
}
```

</li>

<li>

<strong>Get Orders</strong>

|Endpoint|Method|Optional Params|Example|
|:-:|:-:|-|-|
|`/orders`|GET|`start_date [YYYY-MM-DD]`|`http://localhost:8000/orders?start_date=2022-10-10&end_date=2022-10-14`|
|||`end_date [YYYY-MM-DD]`|`http://localhost:8000/orders?start_date=2022-10-10&end_date=2022-10-14`|
|||`status [placed_order, paid, shipping_out, completed]`|`http://localhost:8000/orders?status=paid`|
||||`http://localhost:8000/orders?status=paid,completed`|

<strong>Example :</strong> `http://localhost:8000/orders?start_date=2022-10-10&end_date=2022-10-14&status=paid,completed`
 
```
{
    "status": "successfully.",
    "results": [
        {
            "order_id": 4,
            "product_id": 4,
            "status": "completed",
            "order_date": "2022-10-05T12:38:13.000Z",
            "paid_date": "2022-10-14T13:08:28.000Z",
            "address": "178/25 Soi Vuthipun Ratchaprarob Road Phayathai Bangkok 10400"
        },
        {
            "order_id": 7,
            "product_id": 7,
            "status": "completed",
            "order_date": "2022-10-11T12:38:13.000Z",
            "paid_date": "2022-10-14T13:08:28.000Z",
            "address": "178/25 Soi Vuthipun Ratchaprarob Road Phayathai Bangkok 10400"
        },
        ...
        ...
        ...
    ]
}
```
</li>

<li>

<strong>Create Order</strong>

|Endpoint|Method|Optional Params|Example|
|:-:|:-:|:-:|-|
|`orders`|POST|-|`http://localhost:8000/orders`|

<strong>JSON Body format</strong>

```
{
    "product_id": string,
    "address": string"
}
```

<strong>Example :</strong> `http://127.0.0.1:8000/orders`
 
```
{
    "status": "order created successfully."
}
```
</li>
</ul>
</p>
</details>
<details><summary><h3>Install the Docker PostgreSQL Container</h3></summary>
<p>
<ul>
 
<li>

<strong>Pull image</strong>
<p><a href="https://hub.docker.com/_/postgres">PostgreSQL Docker Image</a></p>

```
$ docker pull postgres:alpine
```
</li>
 
<li>

<strong>Run the container</strong>

```
$ docker run --name PostgreSQL -e POSTGRES_PASSWORD=123456 -p 1150:5432 -d postgres:alpine
```
</li>
 
<li>

<strong>Create a database</strong>

```
$ docker exec -it PostgreSQL bash
```

```
$ psql -U postgres
```

```
$ CREATE DATABASE clothes_shop;
```
</li>
 
<li>

<strong>Check if the database has been created</strong>
 
```
$ \l
```
</li>

</ul>
</p>
</details>
