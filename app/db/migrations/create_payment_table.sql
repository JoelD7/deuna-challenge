CREATE TABLE `payments` (
  `id` text NOT NULL,
  `merchant_id` text NOT NULL,
  `customer_id` text NOT NULL,
  `card_number` int NOT NULL,
  `transaction_id` text,
  `amount` real NOT NULL,
  `status` text,
  `failure_reason` text,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`id`)
);

INSERT INTO `payments` ( `id`, `merchant_id`, `customer_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('1', '1', '1', 1654720058763025, '1', 100.00, 'success', '', datetime(), datetime());
INSERT INTO `payments` ( `id`, `merchant_id`, `customer_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('2', '2', '2', 7048506547895036, '2', 200.00, 'success', '', datetime(), datetime());
INSERT INTO `payments` ( `id`, `merchant_id`, `customer_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('3', '3', '3', 5004896175620369, '3', 300.00, 'success', '', datetime(), datetime());