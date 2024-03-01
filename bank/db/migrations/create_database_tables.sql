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

CREATE TABLE `accounts` (
  `id` text NOT NULL,
  `customer_id` text NOT NULL,
  `balance` real NOT NULL CHECK(balance >= 0),
  `type` text NOT NULL,
  `open_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`id`),
  FOREIGN KEY(customer_id) REFERENCES customers(id)
);

CREATE TABLE `cards` (
  `card_number` integer NOT NULL,
  `customer_id` text NOT NULL,
  `account_id` text,
  `expiration` text NOT NULL,
  `vendor` text NOT NULL,
  `ccv` integer NOT NULL,
  `balance` real CHECK(balance >= 0),
  `type` text NOT NULL,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`card_number`),
  FOREIGN KEY(customer_id) REFERENCES customers(id),
  FOREIGN KEY(account_id) REFERENCES accounts(id)
);

CREATE TABLE `transactions` (
  `id` text NOT NULL,
  `account_id` text NOT NULL,
  `recipient_account_id` text NOT NULL,
  `credit_card_number`integer,
  `amount` real NOT NULL,
  `type` text NOT NULL,
  `status` text,
  `created_date` datetime,
  PRIMARY KEY (`id`)
);

INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('1@mail.com', 'John', 'Doe', 'john@doe.com', '1234567890', '123 Main St', datetime(), datetime());
INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('2@mail.com', 'Jane', 'Doe', 'jane@doe.com', '8504569870', '852 Main St', datetime(), datetime());
INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('3@mail.com', 'John', 'Smith', 'john@smit.com', '1234567890', '789 Main St', datetime(), datetime());

INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('4@mail.com', 'Michael', 'Landon', 'michale@landon.com', '5201469800', '700 Fifth St', datetime(), datetime());
INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('5@mail.com', 'Barbara', 'Walters', 'barbara@walters.com', '5201469800', '852 Seventh St', datetime(), datetime());
INSERT INTO `customers` ( `id`, `first_name`, `last_name`, `email`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('6@mail.com', 'Michael', 'Jordan', 'michael@jordan.com', '8520147963', '789 Ninth St', datetime(), datetime());

INSERT INTO `accounts` ( `id`, `customer_id`, `balance`, `type`, `open_date`, `updated_date` ) VALUES ('1', '1@mail.com', 1000.00, 'checking', datetime(), datetime());
INSERT INTO `accounts` ( `id`, `customer_id`, `balance`, `type`, `open_date`, `updated_date` ) VALUES ('2', '2@mail.com', 2000.00, 'savings', datetime(), datetime());
INSERT INTO `accounts` ( `id`, `customer_id`, `balance`, `type`, `open_date`, `updated_date` ) VALUES ('3', '3@mail.com', 3000.00, 'checking', datetime(), datetime());
INSERT INTO `accounts` ( `id`, `customer_id`, `balance`, `type`, `open_date`, `updated_date` ) VALUES ('4', '4@mail.com', 4000.00, 'savings', datetime(), datetime());
INSERT INTO `accounts` ( `id`, `customer_id`, `balance`, `type`, `open_date`, `updated_date` ) VALUES ('5', '5@mail.com', 5000.00, 'checking', datetime(), datetime());
INSERT INTO `accounts` ( `id`, `customer_id`, `balance`, `type`, `open_date`, `updated_date` ) VALUES ('6', '6@mail.com', 6000.00, 'savings', datetime(), datetime());

INSERT INTO `cards` ( `card_number`, `customer_id`, `account_id`, `expiration`, `vendor`, `ccv`, `balance`, `type`, `created_date`, `updated_date` ) VALUES (1654720058763025, '1@mail.com', '1', '12/2028', 'Visa', 123, 0.0, 'debit', datetime(), datetime());
INSERT INTO `cards` ( `card_number`, `customer_id`, `account_id`, `expiration`, `vendor`, `ccv`, `balance`, `type`, `created_date`, `updated_date` ) VALUES (7048506547895036, '2@mail.com', '', '05/2028', 'Mastercard', 788, 1250.0, 'credit', datetime(), datetime());
INSERT INTO `cards` ( `card_number`, `customer_id`, `account_id`, `expiration`, `vendor`, `ccv`, `balance`, `type`, `created_date`, `updated_date` ) VALUES (5004896175620369, '3@mail.com', '3', '02/2030', 'Visa', 456, 0.0, 'debit', datetime(), datetime());

INSERT INTO `transactions` ( `id`, `account_id`, `recipient_account_id`, `amount`, `type`, `created_date` ) VALUES ('1', '1', '4', 100.00, 'transfer', datetime());