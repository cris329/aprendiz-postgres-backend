-- Recrear tabla aprendiz_postgres en PostgreSQL (ejecutar en DataGrip)

DROP TABLE IF EXISTS aprendiz_postgres CASCADE;

CREATE TABLE IF NOT EXISTS aprendiz_postgres (
    id SERIAL PRIMARY KEY,
    numero_identificacion VARCHAR(255) NOT NULL,
    nombre VARCHAR(255) NOT NULL,
    apellido VARCHAR(255) NOT NULL,
    genero VARCHAR(255) NOT NULL,
    tipo_sangre VARCHAR(255) NOT NULL,
    telefono VARCHAR(255) NOT NULL,
    programa VARCHAR(255) NOT NULL,
    ficha VARCHAR(255) NOT NULL,
    regional VARCHAR(255) NOT NULL
);
