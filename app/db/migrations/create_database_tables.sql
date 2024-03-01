CREATE TABLE `users` (
  `email` text,
  `password` text,
  `role`    text NOT NULL,
  `first_name` text NOT NULL,
  `last_name` text NOT NULL,
  `phone_number` text NOT NULL,
  `address` text NOT NULL,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`email`)
);

CREATE TABLE `cards` (
  `card_number` integer,
  `user_id` text NOT NULL,
  `expiration` text NOT NULL,
  `vendor` text NOT NULL,
  `type` text NOT NULL,
  `ccv` integer NOT NULL,
  `created_date` datetime,
  `updated_date` datetime,
  PRIMARY KEY (`card_number`),
  FOREIGN KEY(user_id) REFERENCES users(email)
);

CREATE TABLE `payments` (
  `id` text,
  `merchant_account_id` text NOT NULL,
  `user_id` text NOT NULL,
  `card_number` integer NOT NULL,
  `transaction_id` text,
  `amount` real NOT NULL,
  `status` text,
  `failure_reason` text,
  `created_date` datetime,
  `updated_date` datetime,
  FOREIGN KEY(user_id) REFERENCES users(email),
  FOREIGN KEY(card_number) REFERENCES cards(card_number)
  PRIMARY KEY (`id`)
);

INSERT INTO `users` ( `email`, `first_name`, `last_name`,`role`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('1@mail.com', 'John', 'Doe', 'customer', '1234567890', '123 Main St', datetime(), datetime());
INSERT INTO `users` ( `email`, `first_name`, `last_name`,`role`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('2@mail.com', 'Jane', 'Doe',  'customer', '8504569870', '852 Main St', datetime(), datetime());
INSERT INTO `users` ( `email`, `first_name`, `last_name`,`role`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('3@mail.com', 'John', 'Smith', 'customer', '1234567890', '789 Main St', datetime(), datetime());
INSERT INTO `users` ( `email`, `first_name`, `last_name`, `role`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('4@mail.com', 'Michael', 'Landon', 'merchant', '5201469800', '700 Fifth St', datetime(), datetime());
INSERT INTO `users` ( `email`, `first_name`, `last_name`, `role`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('5@mail.com', 'Barbara', 'Walters', 'merchant', '5201469800', '852 Seventh St', datetime(), datetime());
INSERT INTO `users` ( `email`, `first_name`, `last_name`, `role`, `phone_number`, `address`, `created_date`, `updated_date` ) VALUES ('6@mail.com', 'Michael', 'Jordan', 'merchant', '8520147963', '789 Ninth St', datetime(), datetime());

INSERT INTO `cards` ( `card_number`, `user_id`, `expiration`, `vendor`,`type`, `ccv`, `created_date`, `updated_date` ) VALUES (1654720058763025, '1@mail.com', '12/2028', 'Visa', 'debit', 123, datetime(), datetime());
INSERT INTO `cards` ( `card_number`, `user_id`, `expiration`, `vendor`,`type`, `ccv`, `created_date`, `updated_date` ) VALUES (7048506547895036, '2@mail.com', '05/2028', 'Mastercard', 'credit', 788, datetime(), datetime());
INSERT INTO `cards` ( `card_number`, `user_id`, `expiration`, `vendor`,`type`, `ccv`, `created_date`, `updated_date` ) VALUES (5004896175620369, '3@mail.com', '02/2030', 'Visa', 'debit', 456, datetime(), datetime());

INSERT INTO `payments` ( `id`, `merchant_account_id`, `user_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('1', '4', '1', 1654720058763025, '1@mail.com', 100.00, 'success', '', datetime(), datetime());
INSERT INTO `payments` ( `id`, `merchant_account_id`, `user_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('2', '5', '2', 7048506547895036, '2@mail.com', 200.00, 'success', '', datetime(), datetime());
INSERT INTO `payments` ( `id`, `merchant_account_id`, `user_id`, `card_number`, `transaction_id`, `amount`, `status`, `failure_reason`, `created_date`, `updated_date` ) VALUES ('3', '6', '3', 5004896175620369, '3@mail.com', 300.00, 'success', '', datetime(), datetime());