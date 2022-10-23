DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
	account_id VARCHAR PRIMARY KEY NOT NULL,
	balance FLOAT NOT NULL
);

CREATE TABLE users (
	user_id VARCHAR PRIMARY KEY NOT NULL,
	account_id VARCHAR NOT NULL,
  	CONSTRAINT fk_user_account
    	FOREIGN KEY(account_id) 
	  	REFERENCES accounts(account_id)
);

CREATE TABLE orders (
	order_id VARCHAR PRIMARY KEY NOT NULL,
	account_id VARCHAR NOT NULL,
	service_id VARCHAR NOT NULL,
	cost FLOAT NOT NULL,
	status VARCHAR NOT NULL,
  	CONSTRAINT fk_order_account
    	FOREIGN KEY(account_id) 
	  	REFERENCES accounts(account_id)
);