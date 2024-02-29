CREATE TABLE `customers` (
  `id` text NOT NULL,
  `first_name` text NOT NULL,
  `last_name` text NOT NULL,
  `email` text NOT NULL,
  `phone_number` text NOT NULL,
  `address` text NOT NULL,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`id`)
);

CREATE TABLE `cards` (
  `card_number` integer NOT NULL,
  `customer_id` text NOT NULL,
  `expiration` text NOT NULL,
  `vendor` text NOT NULL,
  `ccv` integer NOT NULL,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`card_number`),
  FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE `merchants` (
  `id` text NOT NULL,
  `first_name` text NOT NULL,
  `last_name` text NOT NULL,
  `email` text NOT NULL,
  `phone_number` text NOT NULL,
  `address` text NOT NULL,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`id`)
);

CREATE TABLE `payments` (
  `id` text NOT NULL,
  `merchant_id` text NOT NULL,
  `customer_id` text NOT NULL,
  `card_number` integer NOT NULL,
  `transaction_id` text,
  `amount` real NOT NULL,
  `status` text,
  `failure_reason` text,
  `created_date` datetime,
  `updated_date` datetime,
  FOREIGN KEY(customer_id) REFERENCES customers(id),
  FOREIGN KEY(merchant_id) REFERENCES merchants(id),
  FOREIGN KEY(card_number) REFERENCES cards(card_number)
  PRIMARY KEY (`id`)
);

INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('1', 'John', 'Doe', 'john@doe.com', '1234567890', '123 Main St', datetime(), datetime());
INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('2', 'Jane', 'Doe', 'jane@doe.com', '8504569870', '852 Main St', datetime(), datetime());
INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('3', 'John', 'Smith', 'john@smit.com', '1234567890', '789 Main St', datetime(), datetime());

INSERT INTO `cards` ( `card_number`, `customer_id`, `expiration`, `vendor`, `ccv`, `created_date`, `updated_date` ) VALUES (1654720058763025, '1', '12/2028', 'Visa', 123, datetime(), datetime());
INSERT INTO `cards` ( `card_number`, `customer_id`, `expiration`, `vendor`, `ccv`, `created_date`, `updated_date` ) VALUES (7048506547895036, '2', '05/2028', 'Mastercard', 788, datetime(), datetime());
INSERT INTO `cards` ( `card_number`, `customer_id`, `expiration`, `vendor`, `ccv`, `created_date`, `updated_date` ) VALUES (5004896175620369, '3', '02/2030', 'Visa', 456, datetime(), datetime());

INSERT INTO `merchants` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('4', 'Michael', 'Landon', 'michale@landon.com', '5201469800', '700 Fifth St', datetime(), datetime());
INSERT INTO `merchants` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('5', 'Barbara', 'Walters', 'barbara@walters.com', '5201469800', '852 Seventh St', datetime(), datetime());
INSERT INTO `merchants` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('6', 'Michael', 'Jordan', 'michael@jordan.com', '8520147963', '789 Ninth St', datetime(), datetime());

INSERT INTO `payments` ( `id`, `merchant_id`, `customer_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('1', '1', '1', 1654720058763025, '1', 100.00, 'success', '', datetime(), datetime());
INSERT INTO `payments` ( `id`, `merchant_id`, `customer_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('2', '2', '2', 7048506547895036, '2', 200.00, 'success', '', datetime(), datetime());
INSERT INTO `payments` ( `id`, `merchant_id`, `customer_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('3', '3', '3', 5004896175620369, '3', 300.00, 'success', '', datetime(), datetime());