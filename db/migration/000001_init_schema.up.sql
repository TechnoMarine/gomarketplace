CREATE TYPE payment_status AS ENUM ('SUCCESS', 'ERROR', 'WAITING', 'CANCELED');

CREATE TABLE "users" (
                         "user_id" serial PRIMARY KEY,
                         "first_name" varchar not null,
                         "last_name" varchar not null,
                         "sur_name" varchar not null,
                         "email" varchar not null,
                         "password" varchar not null,
                         "address" varchar not null,
                         "created_at" timestamp not null,
                         "updated_at" timestamp
);

CREATE TABLE "categories" (
                              "category_id" serial PRIMARY KEY,
                              "category_name" varchar not null,
                              "parent_id" int,
                              "created_at" timestamp not null,
                              "updated_at" timestamp
);

CREATE TABLE "products" (
                            "product_id" serial PRIMARY KEY,
                            "name" varchar not null,
                            "description" text not null,
                            "price" decimal not null,
                            "stock_quantity" int not null,
                            "seller_id" int not null,
                            "category_id" int not null,
                            "is_active" boolean not null,
                            "created_at" timestamp not null,
                            "updated_at" timestamp
);

CREATE TABLE "orders" (
                          "order_id" serial PRIMARY KEY,
                          "order_date" date not null,
                          "buyer_id" int not null,
                          "created_at" timestamp not null,
                          "updated_at" timestamp
);

CREATE TABLE "order_details" (
                                 "detail_id" serial PRIMARY KEY,
                                 "order_id" integer not null,
                                 "product_id" int not null,
                                 "quantity" int not null,
                                 "total_amount" decimal not null,
                                 "created_at" timestamp not null,
                                 "updated_at" timestamp
);

CREATE TABLE "reviews" (
                           "review_id" serial PRIMARY KEY,
                           "product_id" int not null,
                           "user_id" int not null,
                           "rating" int not null,
                           "comment" text,
                           "created_at" timestamp not null,
                           "updated_at" timestamp
);

CREATE TABLE "shopping_cart" (
                                 "cart_id" serial PRIMARY KEY,
                                 "user_id" int not null,
                                 "product_id" int not null,
                                 "quantity" int not null,
                                 "created_at" timestamp not null,
                                 "updated_at" timestamp
);

CREATE TABLE "shipping_addresses" (
                                      "address_id" serial PRIMARY KEY,
                                      "user_id" int not null,
                                      "address" varchar not null,
                                      "created_at" timestamp not null,
                                      "updated_at" timestamp
);

CREATE TABLE "payments" (
                            "payment_id" serial PRIMARY KEY,
                            "order_id" int not null,
                            "amount" decimal not null,
                            "status" payment_status not null,
                            "payment_date" date not null,
                            "created_at" timestamp not null,
                            "updated_at" timestamp
);

CREATE TABLE "discounts" (
                             "discount_id" serial PRIMARY KEY,
                             "product_id" int not null,
                             "discount_percentage" decimal not null,
                             "created_at" timestamp not null,
                             "updated_at" timestamp
);

CREATE UNIQUE INDEX email_unique ON users (email);

ALTER TABLE "categories" ADD FOREIGN KEY ("parent_id") REFERENCES "categories" ("category_id");

ALTER TABLE "products" ADD FOREIGN KEY ("seller_id") REFERENCES "users" ("user_id");

ALTER TABLE "products" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("buyer_id") REFERENCES "users" ("user_id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "shopping_cart" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "shopping_cart" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");

ALTER TABLE "shipping_addresses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "payments" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "discounts" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");
