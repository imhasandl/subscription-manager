CREATE TABLE subscriptions (
   id UUID PRIMARY KEY,
   service_name VARCHAR(256) NOT NULL,
   price_rub INT,
   user_id UUID NOT NULL,
   start_date DATE NOT NULL,
   end_date DATE NULL
);