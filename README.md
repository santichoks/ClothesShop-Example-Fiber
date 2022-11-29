# ClothesShop-Example-Fiber
<h3 align="center">🚨🚨🚨 this document is in progress 🚨🚨🚨</h3>
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
