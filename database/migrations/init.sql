CREATE TABLE subscriptions (
   id UUID PRIMARY KEY,
   service_name VARCHAR(255) NOT NULL,
   price_rub INT NOT NULL,
   user_id UUID NOT NULL,
   start_date VARCHAR(255) NOT NULL,
   end_date VARCHAR(255)
);