CREATE TABLE payments (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          reservation_id UUID ,
                          amount DECIMAL(10, 2) NOT NULL,
                          payment_method VARCHAR(50) NOT NULL,
                          payment_status VARCHAR(20) CHECK (payment_status IN ('pending', 'completed', 'failed')) NOT NULL,
                          created_at timestamp default  current_timestamp,
                          deleted_at timestamp default  current_timestamp,
                          updated_at timestamp
);

INSERT INTO payments (reservation_id, amount, payment_method, payment_status, created_at, deleted_at, updated_at) VALUES
                                                                                                                      ('b870b7ec-bf70-44f0-9b1b-2906272c7c91', 50.00, 'credit_card', 'completed', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('d8a5b1a8-cfd5-4569-883b-52f5b6f7b8f7', 75.50, 'paypal', 'pending', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('6a497f8b-9d72-43a6-bb91-b92a8a7c97f3', 40.75, 'credit_card', 'failed', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('3f9c889a-8946-4d5b-8d2a-9d6847a09a9e', 100.00, 'bank_transfer', 'completed', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('72b23e65-fd25-4ff0-97d5-80a628f7d0e3', 20.00, 'paypal', 'completed', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('e3e05614-6f92-4bb5-ae02-529d7e9e6a3a', 60.25, 'credit_card', 'pending', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('4d6f5c94-5f37-4673-87de-85f58a9e59c3', 85.00, 'bank_transfer', 'failed', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('1c4f58b6-2c3e-4bbf-83f4-25a0c0a1e4e2', 55.75, 'paypal', 'completed', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('a12d4f5e-0e9c-42b8-847a-3c54b7f6a4c2', 30.00, 'credit_card', 'pending', current_timestamp, NULL, current_timestamp),
                                                                                                                      ('f7a6b3c8-7f2a-4e9a-8a1d-75a8b7f5a4d3', 90.00, 'bank_transfer', 'completed', current_timestamp, NULL, current_timestamp);
