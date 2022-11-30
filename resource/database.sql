CREATE TYPE "gender_list" AS ENUM (
  'Women',
  'Men'
);

CREATE TYPE "size_list" AS ENUM (
  'XS',
  'S',
  'M',
  'L',
  'XL'
);

CREATE TYPE "status_list" AS ENUM (
  'placed_order',
  'paid',
  'shipping_out',
  'completed'
);

CREATE TABLE "products" (
  "product_id" serial2 PRIMARY KEY,
  "gender" gender_list NOT NULL,
  "style" varchar(255) NOT NULL,
  "size" size_list NOT NULL,
  "price" int4 NOT NULL
);

CREATE TABLE "orders" (
  "order_id" serial2 PRIMARY KEY,
  "status" status_list NOT NULL DEFAULT ('placed_order'),
  "order_date" timestamp NOT NULL DEFAULT (now()),
  "paid_date" timestamp DEFAULT null,
  "address" varchar(255) NOT NULL
);

CREATE TABLE "order_item" (
  "id" serial2 PRIMARY KEY,
  "order_id" int,
  "product_id" int,
  "gender" gender_list NOT NULL,
  "style" varchar(255) NOT NULL,
  "size" size_list NOT NULL,
  "price" int4 NOT NULL,
  "quantity" int NOT NULL
);

ALTER TABLE "order_item" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "order_item" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

INSERT INTO public.products (gender,"style","size",price) VALUES
	 ('Men','Red','XS',400),
	 ('Men','Red','S',400),
	 ('Men','Red','M',420),
	 ('Men','Red','L',430),
	 ('Men','Red','XL',450),
	 ('Women','Black','XS',290),
	 ('Women','Black','S',290),
	 ('Women','Black','M',290),
	 ('Women','Black','L',320),
	 ('Women','Black','XL',320),
	 ('Men','Batman','XS',400),
	 ('Men','Batman','S',400),
	 ('Men','Batman','M',420),
	 ('Men','Batman','L',430),
	 ('Men','Batman','XL',450),
	 ('Women','Batman','XS',290),
	 ('Women','Batman','S',290),
	 ('Women','Batman','M',290),
	 ('Women','Batman','L',320),
	 ('Women','Batman','XL',320),
	 ('Men','Spiderman','XS',400),
	 ('Men','Spiderman','S',400),
	 ('Men','Spiderman','M',420),
	 ('Men','Spiderman','L',430),
	 ('Men','Spiderman','XL',450),
	 ('Women','Spiderman','XS',290),
	 ('Women','Spiderman','S',290),
	 ('Women','Spiderman','M',290),
	 ('Women','Spiderman','L',320),
	 ('Women','Spiderman','XL',320);
