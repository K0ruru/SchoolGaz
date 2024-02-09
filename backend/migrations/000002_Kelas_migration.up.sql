CREATE TABLE IF NOT EXISTS kelas (
  id_kelas SERIAL PRIMARY KEY,
  siswa INT REFERENCES users(nis),
  walas INT REFERENCES users(nis),
  nama_kelas VARCHAR(25) NOT NULL
);

