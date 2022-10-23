CREATE TABLE IF NOT EXISTS accounts (
	account_id VARCHAR PRIMARY KEY NOT NULL,
	balance FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	user_id VARCHAR PRIMARY KEY NOT NULL,
	account_id VARCHAR NOT NULL,
  	CONSTRAINT fk_user_account
    	FOREIGN KEY(account_id) 
	  	REFERENCES accounts(account_id)
);

CREATE TABLE IF NOT EXISTS orders (
	order_id VARCHAR PRIMARY KEY NOT NULL,
	account_id VARCHAR NOT NULL,
	service_id VARCHAR NOT NULL,
	cost FLOAT NOT NULL,
	status VARCHAR NOT NULL,
  	CONSTRAINT fk_order_account
    	FOREIGN KEY(account_id) 
	  	REFERENCES accounts(account_id)
);