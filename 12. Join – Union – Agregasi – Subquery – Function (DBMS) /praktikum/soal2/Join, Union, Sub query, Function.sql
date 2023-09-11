-- 1. Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT * FROM transactions WHERE user_id IN (1, 2);

-- 2. Tampilkan jumlah harga transaksi user id 1.
SELECT user_id, SUM(total_price) AS total_harga_transaksi 
FROM transactions WHERE user_id = 1 GROUP BY user_id;

-- 3. Tampilkan total transaksi dengan product type 2.
SELECT COUNT(*) AS total_transaksi_product_type_2 FROM transactions trx
JOIN transaction_details trxd ON trx.id = trxd.transaction_id
JOIN products prod ON trxd.product_id = prod.id
WHERE prod.product_type = 2;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT p.*, pt.name AS product_type_name
FROM products p
JOIN product_types pt ON p.product_type = pt.id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT trx.*, p.name AS product_name, u.name AS user_name FROM transactions trx
JOIN transaction_details td ON trx.id = td.transaction_id
JOIN products p ON td.product_id = p.id
JOIN users u ON trx.user_id = u.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER //
CREATE TRIGGER after_delete_transaction
BEFORE DELETE ON transactions
FOR EACH ROW
BEGIN
  DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END;
//
DELIMITER ;

-- 7. Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER //
CREATE TRIGGER after_delete_transaction_detail
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
  UPDATE transactions SET total_qty = total_qty - OLD.qty WHERE id = OLD.transaction_id;
END;
//
DELIMITER ;


-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT p.* FROM products p
WHERE p.id NOT IN (SELECT DISTINCT product_id FROM transaction_details);
