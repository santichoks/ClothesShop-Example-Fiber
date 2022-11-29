# ClothesShop-Example-Practice
<h3 align="center">🚨🚨🚨 this document is in progress 🚨🚨🚨</h3>
<details><summary><h3>API Endpoints</h3></summary>
<p>
<ul>
<li>

<p><strong>Get Products</strong></p>

|Endpoint|Method|Optional Params|Example|
|:-:|:-:|-|-|
|`/products`|GET|`gender [Men, Women]`|`http://localhost:3000/products?gender=Men`|
||||`http://localhost:3000/products?gender=Men&gender=Women`|
|||`size [XS, S, M, L, XL]`|`http://localhost:3000/products?size=L`|
||||`http://localhost:3000/products?size=L&size=XL`|
|||`style [Red, Black, Batman, Spiderman]`|`http://localhost:3000/products?style=Batman`|
||||`http://localhost:3000/products?style=Batman&style=Spiderman`|

<strong>Example :</strong> `http://localhost:3000/products?gender=Men&style=Batman&style=Spiderman&size=L&size=XL`

```
[
    {
        "product_id": 14,
        "gender": "Men",
        "style": "Batman",
        "size": "L",
        "price": 430
    },
    {
        "product_id": 15,
        "gender": "Men",
        "style": "Batman",
        "size": "XL",
        "price": 450
    },
    {
        "product_id": 24,
        "gender": "Men",
        "style": "Spiderman",
        "size": "L",
        "price": 430
    },
    {
        "product_id": 25,
        "gender": "Men",
        "style": "Spiderman",
        "size": "XL",
        "price": 450
    }
]
```

</li>

<li>

<strong>Get Orders</strong>

|Endpoint|Method|Optional Params|Example|
|:-:|:-:|-|-|
|||||
|||||
|||||
|||||

<strong>Example :</strong> `#################################################################`
 
```
########################################################################################################
########################################################################################################
########################################################################################################
########################################################################################################
########################################################################################################

```
</li>

<li>

<strong>Create Order</strong>

|Endpoint|Method|Optional Params|Example|
|:-:|:-:|:-:|-|
|`orders`|POST|-|`http://localhost:3000/orders`|

<p><strong>JSON Body format</strong></p>

```
{
    "product_details": {
        "product_id":[6],
        "gender":["Women"],
        "style":["Black"],
        "size":["XS"],
        "price":[290],
        "quantity":[1]
    },
    "address":"12/9 Phaholyothin Sukhumvit Bangkok 10900"
}
```
<strong>Note :</strong> The same index of the array represents the same product detail.
```
{
    "product_details": {
        "product_id":[1, 2, 3],
        "gender":["Men", "Men", "Men"],
        "style":["Red", "Red", "Red"],
        "size":["XS", "S", "M"],
        "price":[400, 400, 420],
        "quantity":[1, 3, 5]
    },
    "address":"12/9 Phaholyothin Sukhumvit Bangkok 10900"
}
```

<strong>Example :</strong> `http://127.0.0.1:3000/orders`
 
```
{
    "status": "order has been created"
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
