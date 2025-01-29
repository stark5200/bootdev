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

DELETE FROM users
    WHERE username = "Sammy93";



-- TEST SUITE, DON'T TOUCH BELOW THIS LINE --

SELECT * FROM users;

SELECT username FROM users WHERE is_admin == true;

sqlQuery := fmt.Sprintf(`
INSERT INTO users(name, age, country_code)
VALUES ('%s', %v, '%s');
`, user.Name, user.Age, user.CountryCode)
