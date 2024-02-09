-- Create ENUM type if it doesn't exist
CREATE TYPE IF NOT EXISTS Erole AS ENUM ('Admin', 'User', 'Guru');

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    nis SERIAL PRIMARY KEY,
    nama VARCHAR(50) NOT NULL,
    passphrase VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    gender VARCHAR(15) NOT NULL,
    agama VARCHAR(30) NOT NULL,
    profilepicture VARCHAR(100),
    status VARCHAR(255) DEFAULT 'pending',
    role Erole DEFAULT 'User',
    kelas INT REFERENCES kelas(id),
    mapel INT REFERENCES mapel(id)  
);

