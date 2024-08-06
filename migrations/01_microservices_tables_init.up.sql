-- INSERT INTO users(name, userId) VALUES ('name','userId where create some uuid.newV4()')

-- SELECT * from users

-- Insert into products(name, description, userId) values ('test 1', 'ddnsdjnvaksmvlkdmvl','users(id)')

-- select * from products

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS products CASCADE;


CREATE TABLE users 
(
  id UUID PRIMARY KEY   DEFAULT uuid_generate_v4(),
  name VARCHAR(250),
  surname VARCHAR(250),
  phone VARCHAR(250),
  userId VARCHAR(500) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products 
(
  id UUID PRIMARY KEY    DEFAULT uuid_generate_v4(),
  name VARCHAR(250),
  description VARCHAR(250),
  userId UUID REFERENCES users(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);