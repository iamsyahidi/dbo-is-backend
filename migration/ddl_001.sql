CREATE TABLE IF NOT EXISTS public.customer (
	id varchar(36) NOT NULL,
	email varchar(100) NOT NULL,
	"name" varchar(250) NOT NULL,
	"password" varchar(150) NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT customer_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_customer_email ON public.customer USING btree (email);
CREATE INDEX IF NOT EXISTS idx_customer_id ON public.customer USING btree (id);
CREATE INDEX IF NOT EXISTS idx_customer_name ON public.customer USING btree ("name");
CREATE INDEX IF NOT EXISTS idx_customer_password ON public.customer USING btree ("password");
CREATE INDEX IF NOT EXISTS idx_customer_status ON public.customer USING btree ("status");

INSERT INTO public.customer (id, email, "name", "password", "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'ilhamsyahidi66@gmail.com', 'ilham', '$2a$10$2RLJwYzS2wCGTkgr145miOvXe52nI0JRE.8uOgqEmM8ulmmRufWle', 'active', now(), 'ilham', NULL, NULL);


CREATE TABLE IF NOT EXISTS public.product (
	id varchar(36) NOT NULL,
	"name" varchar(250) NOT NULL,
	price numeric NULL,
	stock numeric NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT product_pkey PRIMARY KEY (id)
);
CREATE INDEX IF NOT EXISTS idx_product_id ON public.product USING btree (id);
CREATE INDEX IF NOT EXISTS idx_product_name ON public.product USING btree ("name");
CREATE INDEX IF NOT EXISTS idx_product_price ON public.product USING btree (price);
CREATE INDEX IF NOT EXISTS idx_product_stock ON public.product USING btree (stock);
CREATE INDEX IF NOT EXISTS idx_product_status ON public.product USING btree ("status");


INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Chitato', 10000, 108, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Pocari Sweat 500ml', 5000, 191, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Chiki ball rasa', 8000, 1000, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Pulpen Standar', 1500, 100, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Susu SGM', 50000, 2, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Nextar', 2000, 50, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Milo', 20000, 10, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Sirop', 20000, 5, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Pempes', 50000, 5, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Maxim Neostone', 215000, 2, 'active', now(), 'admin', NULL, NULL);
INSERT INTO public.product (id, "name", price, stock, "status", created_at, created_by, updated_at, updated_by) VALUES(gen_random_uuid(), 'Lifebuoy Body Foam', 25000, 5, 'active', now(), 'admin', NULL, NULL);


CREATE TABLE IF NOT EXISTS public."order" (
	invoice varchar(100) NOT NULL,
	customer_id varchar(100) NOT NULL,
	amount numeric NULL,
	payment bool DEFAULT false NOT NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL,
	CONSTRAINT order_pkey PRIMARY KEY (invoice)
);
CREATE INDEX IF NOT EXISTS idx_order_amount ON public."order" USING btree (amount);
CREATE INDEX IF NOT EXISTS idx_order_invoice ON public."order" USING btree (invoice);
CREATE INDEX IF NOT EXISTS idx_order_payment ON public."order" USING btree (payment);
CREATE INDEX IF NOT EXISTS idx_order_status ON public."order" USING btree ("status");
CREATE INDEX IF NOT EXISTS idx_order_customer_id ON public."order" USING btree (customer_id);


CREATE TABLE IF NOT EXISTS public.order_detail (
	invoice varchar(100) NOT NULL,
	product_id varchar(100) NOT NULL,
	qty numeric NULL,
	price numeric NULL,
	amount numeric NULL,
	"status" varchar(10) NOT NULL,
	created_at timestamptz DEFAULT now() NOT NULL,
	created_by varchar(150) NOT NULL,
	updated_at timestamptz NULL,
	updated_by varchar(150) DEFAULT NULL::character varying NULL
);
CREATE INDEX IF NOT EXISTS idx_order_detail_amount ON public.order_detail USING btree (amount);
CREATE INDEX IF NOT EXISTS idx_order_detail_invoice ON public.order_detail USING btree (invoice);
CREATE INDEX IF NOT EXISTS idx_order_detail_price ON public.order_detail USING btree (price);
CREATE INDEX IF NOT EXISTS idx_order_detail_product_id ON public.order_detail USING btree (product_id);
CREATE INDEX IF NOT EXISTS idx_order_detail_qty ON public.order_detail USING btree (qty);
CREATE INDEX IF NOT EXISTS idx_order_detail_status ON public.order_detail USING btree ("status");