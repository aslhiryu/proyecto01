--creo el usuario
--CREATE USER gestion WITH PASSWORD 'gestion';
--GRANT ALL PRIVILEGES ON DATABASE banwire_gestion TO gestion;


--creo la estructura de BD
CREATE OR REPLACE FUNCTION actualizaBD() RETURNS VOID AS
$$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='transaccion')
	THEN
        CREATE TABLE transaccion(
            id_transaccion      VARCHAR(36) NOT NULL,
            num_autorizacion    INT NOT NULL,
            monto               NUMERIC(11,3) NOT NULL,
            fecha               TIMESTAMP NOT NULL,
            id_terminal         VARCHAR(36) NOT NULL,
            id_servicio         VARCHAR(36) NOT NULL,
            id_suscripcion      VARCHAR(36) NOT NULL,
            CONSTRAINT PK_transaccion PRIMARY KEY(id_transaccion)
        );
	END IF;

	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='suscripcion')
	THEN
        CREATE TABLE suscripcion(
            id_suscripcion      VARCHAR(36) NOT NULL,
            id_plan             VARCHAR(36) NOT NULL,
            id_tarjeta          VARCHAR(36) NOT NULL,
            CONSTRAINT PK_suscripcion PRIMARY KEY(id_suscripcion)
        );
	END IF;

	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='plan')
	THEN
        CREATE TABLE plan(
            id_plan             VARCHAR(36) NOT NULL,
            nombre              VARCHAR(36) NOT NULL,
            id_comercio         VARCHAR(36) NOT NULL,
            CONSTRAINT PK_plan PRIMARY KEY(id_plan)
        );
	END IF;
END;
$$
LANGUAGE 'plpgsql';

SELECT actualizaBD();
DROP FUNCTION actualizaBD();