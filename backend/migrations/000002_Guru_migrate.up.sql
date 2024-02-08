DROP TABLE IF EXISTS guru;


CREATE TABLE guru (
    nik SERIAL PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    passphrase VARCHAR(100) NOT NULL,
    email VARCHAR(25) NOT NULL,
    gender VARCHAR(15) NOT NULL,
    agama VARCHAR(30) NOT NULL,
  status statusG DEFAULT 'pending'
);
