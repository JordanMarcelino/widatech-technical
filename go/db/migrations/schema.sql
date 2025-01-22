drop table if exists invoices cascade;
drop table if exists products cascade;
drop index if exists idx_invoice_date;
drop index if exists idx_fk_product_invoice;
create table if not exists invoices (
    invoice_no varchar(255) not null primary key,
    invoice_date date not null,
    customer_name varchar(255) not null,
    sales_person_name varchar(255) not null,
    payment_type varchar(255) not null,
    notes text not null
);
create table if not exists products (
    id bigserial primary key,
    invoice_no varchar(255) references invoices(invoice_no) not null,
    name varchar(255) not null,
    quantity int not null,
    total_cost_of_goods decimal(10, 2) not null,
    total_price_sold decimal(10, 2) not null
);
create index if not exists idx_invoice_date on invoices(invoice_date);
create index if not exists idx_fk_product_invoice on products(invoice_no);
