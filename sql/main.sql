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

SELECT round(avg(age)) AS round_age
FROM users
WHERE country_code = 'US';




-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT * FROM users;

SELECT username FROM users WHERE is_admin == true;

sqlQuery := fmt.Sprintf(`
INSERT INTO users(name, age, country_code)
VALUES ('%s', %v, '%s');
`, user.Name, user.Age, user.CountryCode)
