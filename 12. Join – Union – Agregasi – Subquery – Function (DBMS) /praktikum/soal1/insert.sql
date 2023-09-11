-- 1. Insert 5 operators pada table operators.
INSERT INTO operators (name) VALUES
('Operator 1'),
('Operator 2'),
('Operator 3'),
('Operator 4'),
('Operator 5');

-- 2. Insert 3 product types.
INSERT INTO product_types (name) VALUES
('Type 1'),
('Type 2'),
('Type 3');

-- 3. Insert 2 products dengan product type id = 1, dan operators id = 3.
INSERT INTO products (product_type, operator_id, code, name, status) VALUES
(1, 3, 'P001T1', 'Product 1', 'Completed'),
(1, 3, 'P002T1', 'Product 2', 'Completed');

-- 4. Insert 3 products dengan product type id = 2, dan operators id = 1.
INSERT INTO products (product_type, operator_id, code, name, status) VALUES
(2, 1, 'P001T2', 'Product 3', 'Completed'),
(2, 1, 'P002T2', 'Product 4', 'Completed'),
(2, 1, 'P003T2', 'Product 5', 'Completed');

-- 5. Insert 3 products dengan product type id = 3, dan operators id = 4.
INSERT INTO products (product_type, operator_id, code, name, status) VALUES
(3, 4, 'P001T3', 'Product 6', 'Completed'),
(3, 4, 'P002T3', 'Product 7', 'Completed'),
(3, 4, 'P003T3', 'Product 8', 'Completed');

-- 6. Insert product descriptions pada setiap product.
INSERT INTO product_description (product_id, description) VALUES
(1, 'Deskripsi produk untuk ID 1'),
(2, 'Deskripsi produk untuk ID 2'),
(3, 'Deskripsi produk untuk ID 3'),
(4, 'Deskripsi produk untuk ID 4'),
(5, 'Deskripsi produk untuk ID 5'),
(6, 'Deskripsi produk untuk ID 6'),
(7, 'Deskripsi produk untuk ID 7'),
(8, 'Deskripsi produk untuk ID 8');

-- 7. Insert 3 payment methods.
INSERT INTO payment_methods (name, status) VALUES
('Qris', 1),
('Gopay', 1),
('Tunai', 1);

-- 8. Insert 5 users pada tabel user.
INSERT INTO users (id, name, status, dob, gender) VALUES
(1, 'Noru', 1, '2000-01-01', 'L'),
(2, 'Nino', 1, '1945-07-17', 'L'),
(3, 'Nana', 1, '1908-03-10', 'P'),
(4, 'Nunu', 1, '1990-12-20', 'P'),
(5, 'Naez', 1, '2024-08-05', 'L');

-- 9. Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
(1, 1, 'Completed', 2, 100.00),
(1, 2, 'Completed', 3, 150.00),
(1, 3, 'Pending', 1, 500.00);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
(2, 1, 'Completed', 1, 40.00),
(2, 2, 'Completed', 2, 100.00),
(2, 3, 'Completed', 4, 750.00);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
(3, 1, 'Pending', 2, 100.00),
(3, 2, 'Completed', 3, 150.00),
(3, 3, 'Completed', 1, 50.00);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
(4, 1, 'Completed', 3, 150.00),
(4, 2, 'Completed', 2, 100.00),
(4, 3, 'Completed', 1, 50.00);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
(5, 1, 'Pending', 1, 12.00),
(5, 2, 'Completed', 3, 10.00),
(5, 3, 'Completed', 2, 90.00);

-- 10. Insert 3 products di masing-masing transaksi.
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
(1, 1, 'Completed', 2, 20.00),
(3, 2, 'Completed', 1, 15.00),
(2, 3, 'Completed', 3, 30.00);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
(4, 1, 'Completed', 1, 10.00),
(5, 2, 'Completed', 2, 20.00),
(6, 3, 'Completed', 2, 20.00);

INSERT INTO transaction_details (transaction_id, product_id, status, qty, price)
VALUES
(9, 1, 'Completed', 3, 30.00),
(8, 2, 'Completed', 1, 10.00),
(7, 3, 'Completed', 4, 40.00);
