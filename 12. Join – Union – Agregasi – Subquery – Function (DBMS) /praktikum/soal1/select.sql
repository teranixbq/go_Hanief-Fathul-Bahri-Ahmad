-- 1. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name FROM users WHERE gender = 'L';

-- 2. Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id = 3;

-- 3. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * FROM users WHERE created_at >= DATE_SUB(CURRENT_DATE(), INTERVAL 7 DAY) AND name LIKE '%a%';

-- 4. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) AS total_perempuan FROM users WHERE gender = 'P';

-- 5. Tampilkan data pelanggan dengan urutan sesuai nama abjad.
SELECT * FROM users ORDER BY name ASC;

-- 6. Tampilkan 5 data pada data product.
SELECT * FROM products LIMIT 5;
