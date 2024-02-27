CREATE TYPE payment_status AS ENUM ('SUCCESS', 'ERROR', 'WAITING', 'CANCELED');

CREATE TABLE "users" (
                         "user_id" bigserial PRIMARY KEY,
                         "first_name" varchar,
                         "last_name" varchar,
                         "sur_name" varchar,
                         "email" varchar,
                         "password" varchar,
                         "address" varchar,
                         "created_at" timestamp,
                         "updated_at" timestamp
);

CREATE TABLE "categories" (
                              "category_id" bigserial PRIMARY KEY,
                              "category_name" varchar,
                              "parent_id" int,
                              "created_at" timestamp,
                              "updated_at" timestamp
);

CREATE TABLE "products" (
                            "product_id" bigserial PRIMARY KEY,
                            "name" varchar,
                            "description" text,
                            "price" decimal,
                            "stock_quantity" int,
                            "seller_id" int,
                            "category_id" int,
                            "is_active" boolean,
                            "created_at" timestamp,
                            "updated_at" timestamp
);

CREATE TABLE "orders" (
                          "order_id" bigserial PRIMARY KEY,
                          "order_date" date,
                          "buyer_id" int,
                          "created_at" timestamp,
                          "updated_at" timestamp
);

CREATE TABLE "order_details" (
                                 "detail_id" bigserial PRIMARY KEY,
                                 "order_id" int,
                                 "product_id" int,
                                 "quantity" int,
                                 "total_amount" decimal,
                                 "created_at" timestamp,
                                 "updated_at" timestamp
);

CREATE TABLE "reviews" (
                           "review_id" bigserial PRIMARY KEY,
                           "product_id" int,
                           "user_id" int,
                           "rating" int,
                           "comment" text,
                           "created_at" timestamp,
                           "updated_at" timestamp
);

CREATE TABLE "shopping_cart" (
                                 "cart_id" bigserial PRIMARY KEY,
                                 "user_id" int,
                                 "product_id" int,
                                 "quantity" int,
                                 "created_at" timestamp,
                                 "updated_at" timestamp
);

CREATE TABLE "shipping_addresses" (
                                      "address_id" bigserial PRIMARY KEY,
                                      "user_id" int,
                                      "address" varchar,
                                      "created_at" timestamp,
                                      "updated_at" timestamp
);

CREATE TABLE "payments" (
                            "payment_id" bigserial PRIMARY KEY,
                            "order_id" int,
                            "amount" decimal,
                            "status" payment_status,
                            "payment_date" date,
                            "created_at" timestamp,
                            "updated_at" timestamp
);

CREATE TABLE "discounts" (
                             "discount_id" bigserial PRIMARY KEY,
                             "product_id" int,
                             "discount_percentage" decimal,
                             "created_at" timestamp,
                             "updated_at" timestamp
);

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
