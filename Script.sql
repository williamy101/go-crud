create database if not exists inventory_management

use inventory_management

create table products (
	product_id int auto_increment primary key,
	name varchar(50) not null,
	description text,
	price decimal(10,2) not null,
	category varchar(50),
	created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
)

create table inventories (
	inventory_id int auto_increment primary key,
	product_id int not null,
	stock int not null default 0,
	location varchar(100) not null,
	created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
	 foreign key (product_id) references product(product_id)
)

create table orders (
	order_id int auto_increment primary key,
	product_id int not null,
	quantity int not null,
	order_date timestamp default current_timestamp, 
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    foreign key (product_id) references product(product_id)
)

insert into products (name, description, price, category, created_at, updated_at) 
values
('laptop', 'high-performance laptop with 16gb ram and 1tb ssd.', 1200.00, 'electronics', now(), now()),
('office chair', 'ergonomic chair with adjustable height and lumbar support.', 200.00, 'furniture', now(), now()),
('smartphone', '5g-enabled smartphone with 128gb storage and 12mp camera.', 800.00, 'electronics', now(), now());

insert into inventories (product_id, stock, location, created_at, updated_at) 
values
(1, 50, 'warehouse a', now(), now()),
(2, 30, 'store b', now(), now()); 

insert into inventories (product_id, stock, location, created_at, updated_at) 
values
(3, 25, 'store b', now(), now()); 

insert into orders (product_id, quantity, order_date, created_at, updated_at) 
values
(1, 2, '2024-11-20', now(), now()),
(2, 1, '2024-11-21', now(), now()),
(3, 3, '2024-11-22', now(), now()); 

select location, sum(stock) as total_stock
from inventories
group by location
order by total_stock desc

select p.name as product_name, sum(o.quantity * p.price) as total_revenue
from products p
join orders o on p.product_id = o.product_id
group by p.name
order by total_revenue desc;

select avg(o.quantity * p.price) as avg_revenue
from orders o
join products p on o.product_id = p.product_id
order by avg_revenue desc;

alter table products
add image_path varchar(255) default null