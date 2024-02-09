CREATE TABLE mapel(
  id_Mapel SERIAL PRIMARY KEY,
  nama_mapel VARCHAR(50) NOT NULL,
  guru INT REFERENCES users(nis)
);
