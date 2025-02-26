ALTER TABLE people
RENAME TO users;

ALTER TABLE users
RENAME COLUMN tag TO username;

ALTER TABLE users
ADD COLUMN password TEXT;

-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

PRAGMA TABLE_INFO('users');

INSERT INTO users(id, name, age, country_code, username, password, is_admin)
VALUES (1, 'David', 34, 'US', 'DavidDev', 'insertPractice', false);

INSERT INTO users(id, name, age, country_code, username, password, is_admin)
VALUES (2, 'Samantha', 29, 'BR', 'Sammy93', 'addingRecords!', false);

INSERT INTO users(name, age, country_code, username, password, is_admin)
VALUES ('Lance', 20, 'US', 'LanChr', 'bootdevisbest', false);

INSERT INTO users(name, age, country_code, username, password, is_admin)
VALUES ('Tiffany', 28, 'US', 'Tifferoon', 'autoincrement', true);

SELECT name, age
FROM users
WHERE age BETWEEN 18 and 30;

SELECT DISTINCT country_code
    FROM users;

DELETE FROM users
    WHERE username = "Sammy93";

SELECT *
    FROM users
    WHERE country_code = 'CA'
    AND age < 18;

SELECT COUNT(*) AS junior_count
    FROM users
    WHERE (country_code = 'US' OR country_code = 'CA')
    AND age < 18;

SELECT name, age, country_code
    FROM users
    WHERE country_code IN ('US', 'CA', 'MX');

SELECT * FROM users
WHERE name LIKE 'Al%';

SELECT * FROM users
    WHERE name LIKE 'Al___';

SELECT amount, note AS birthday_message
FROM transactions WHERE sender_id = 10;

SELECT * from users
    WHERE name LIKE 'or_%';

SELECT * from users
    WHERE name LIKE '__ing';

UPDATE users
SET is_admin = true
WHERE name = 'Lane';

SELECT COUNT(*) FROM users WHERE country_code = 'US';

SELECT *,
    IIF(was_successful = true, 'No action required', 'Perform an audit') AS audit
    FROM transactions;

UPDATE users
  SET country_code = 'US'
  WHERE country_code = 'USA';

-- challenges --
SELECT *, (age > 55 or country_code = 'CA') AS discount_eligible
  FROM users;

SELECT * FROM products
    WHERE product_name LIKE '%berry%'
    LIMIT 5;

SELECT * FROM transactions
  WHERE amount BETWEEN 10 AND 80
  ORDER BY amount DESC;

SELECT * FROM transactions
WHERE amount BETWEEN 10 AND 80
ORDER BY amount DESC
LIMIT 4;

SELECT COUNT(*)
FROM transactions
WHERE user_id = 6 AND was_successful = 1;

SELECT SUM(amount)
FROM transactions 
WHERE user_id = 9 AND was_successful = 1;


SELECT name, username
  FROM users
  Where (username LIKE '%cashpal%' or username LIKE '%support%') and is_admin = false;

SELECT * FROM transactions
    WHERE note LIKE '%lunch%'
    LIMIT 5;

SELECT user_id, max(amount)
FROM transactions
WHERE was_successful = true
AND (
  (user_id = 4 AND sender_id is NOT NULL)
  OR recipient_id = 4
);

SELECT min(age)
FROM users
WHERE country_code = 'US';

SELECT user_id, sum(amount) AS balance
FROM transactions
WHERE was_successful = 1
GROUP BY user_id;

SELECT sender_id, sum(amount) AS balance
FROM transactions
WHERE note LIKE '%lunch%' AND NOT sender_id IS NULL AND was_successful = 1
GROUP BY sender_id
HAVING balance > 20
ORDER BY balance;

SELECT country_code, round(avg(age)) AS average_age
FROM users
GROUP BY country_code;

SELECT recipient_id, count(*) AS transactions_received
FROM transactions
WHERE was_successful = 1 AND recipient_id IS NOT NULL
GROUP BY recipient_id
ORDER BY transactions_received DESC
LIMIT 2;

SELECT *
FROM transactions
WHERE user_id IN (
    SELECT id
    FROM users
    WHERE name LIKE 'David'
);

SELECT *
FROM users
WHERE id IN (
    SELECT id
    FROM users
    WHERE floor(age_in_days / 365) > 40
);

SELECT *
FROM users
WHERE is_admin = 0 AND id IN (
    SELECT sender_id
    FROM transactions
    WHERE note LIKE '%invoice%' OR note LIKE '%tax%'
);

CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  is_admin BOOLEAN
);

CREATE TABLE countries (
  id INTEGER PRIMARY KEY,
  country_code TEXT NOT NULL,
  name TEXT NOT NULL,
  user_id INTEGER,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  age INTEGER NOT NULL,
  username TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  is_admin BOOLEAN
);

CREATE TABLE countries (
  id INTEGER PRIMARY KEY,
  country_code TEXT,
  name TEXT
);

CREATE TABLE users_countries (
  country_id INTEGER,
  user_id INTEGER,
  UNIQUE(country_id, user_id)
);


SELECT round(avg(age)) AS round_age
FROM users
WHERE country_code = 'US';

CREATE TABLE companies (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  num_employees INTEGER NOT NULL
);

-- Don't touch between these comments --

INSERT INTO companies(name, num_employees)
  VALUES ('Pfizer', 10000);
INSERT INTO companies(name, num_employees)
  VALUES ('WorldBanc', 80);
INSERT INTO companies(name, num_employees)
  VALUES ('Fantasy Quest', 30);
INSERT INTO companies(name, num_employees)
  VALUES ('Walmart', 1000);

-- Don't touch between these comments --

SELECT *,
  CASE 
    WHEN num_employees >= 100 THEN 'Large'
    ELSE 'Small'
  END AS size
FROM companies;




-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT * FROM users;

SELECT username FROM users WHERE is_admin == true;

sqlQuery := fmt.Sprintf(`
INSERT INTO users(name, age, country_code)
VALUES ('%s', %v, '%s');
`, user.Name, user.Age, user.CountryCode)

CREATE TABLE banks (
  id INTEGER PRIMARY KEY,
  name TEXT,
  routing_number INTEGER
);

CREATE TABLE users_banks (
  user_id INTEGER,
  bank_id INTEGER,
  UNIQUE (user_id, bank_id)
);

SELECT *
FROM users
INNER JOIN countries
ON countries.country_code = users.country_code;
