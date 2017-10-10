--creo el usuario
--CREATE USER gestion WITH PASSWORD 'gestion';
--GRANT ALL PRIVILEGES ON DATABASE banwire_gestion TO gestion;


--creo la estructura de BD
CREATE OR REPLACE FUNCTION actualizaBD() RETURNS VOID AS
$$
BEGIN

	--------   Para notificaciones   --------
	-- Catalogo de etados de notificacion
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_estado_notificacion')
	THEN
        CREATE TABLE ctl_estado_notificacion(
            id_estado_notificacion	VARCHAR(36) NOT NULL,
            descripcion 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlEdoNotificacion PRIMARY KEY(id_estado_notificacion)
        );
	END IF;

	
	
	--------   Para generales   --------
	-- Catalogo de paises
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_pais')
	THEN
        CREATE TABLE ctl_pais(
            id_pais					VARCHAR(36) NOT NULL,
            nombre		 			VARCHAR(25) NOT NULL,
            CONSTRAINT PK_CtlPais PRIMARY KEY(id_pais)
        );
	END IF;

	
	
	--------   Para cliente   --------
	-- Catalogo de tipos de tarjetas
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_tipo_tarjeta')
	THEN
        CREATE TABLE ctl_tipo_tarjeta(
            id_tipo_tarjeta			VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlTipTarjeta PRIMARY KEY(id_tipo_tarjeta)
        );
	END IF;

	-- Catalogo de emisores de tarjetas
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_emisor_tarjeta')
	THEN
        CREATE TABLE ctl_emisor_tarjeta(
            id_emisor_tarjeta		VARCHAR(36) NOT NULL,
            nombre		 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlEmisor PRIMARY KEY(id_emisor_tarjeta)
        );
	END IF;

	-- Catalogo de estados de tarjetas
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_estado_tarjeta')
	THEN
        CREATE TABLE ctl_estado_tarjeta(
            id_estado_tarjeta		VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlEdoTarjeta PRIMARY KEY(id_estado_tarjeta)
        );
	END IF;

	-- tarjetas
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='tarjeta')
	THEN
        CREATE TABLE tarjeta(
            id_tarjeta				VARCHAR(36) NOT NULL,
            codigo_autorizacion		VARCHAR(25),
            tarejtahabiente			VARCHAR(80) NOT NULL,
            id_tipo_tarjeta			VARCHAR(36) NOT NULL,
            id_emisor_tarjeta		VARCHAR(36) NOT NULL,
            mail_usuario			VARCHAR(80),
            digitos					CHAR(16) NOT NULL,
            bin						VARCHAR(25) NOT NULL,
            vigencia				TIMESTAMP NOT NULL,
            marca					VARCHAR(25) NOT NULL,
            id_pais					VARCHAR(36) NOT NULL,
            ultimo_cargo			NUMERIC(11,3),
            cvv						CHAR(3) NOT NULL,
            token					VARCHAR(25),
			id_estado_tarjeta		VARCHAR(36) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
            creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_Tarjeta PRIMARY KEY(id_tarjeta)
        );

		ALTER TABLE tarjeta ADD CONSTRAINT FK_TarjetaTipo FOREIGN KEY(id_tipo_tarjeta)
            REFERENCES ctl_tipo_tarjeta(id_tipo_tarjeta);

		ALTER TABLE tarjeta ADD CONSTRAINT FK_TarjetaEmisor FOREIGN KEY(id_emisor_tarjeta)
            REFERENCES ctl_emisor_tarjeta(id_emisor_tarjeta);

		ALTER TABLE tarjeta ADD CONSTRAINT FK_TarjetaPais FOREIGN KEY(id_pais)
            REFERENCES ctl_pais(id_pais);

		ALTER TABLE tarjeta ADD CONSTRAINT FK_TarjetaEstado FOREIGN KEY(id_estado_tarjeta)
            REFERENCES ctl_estado_tarjeta(id_estado_tarjeta);
	END IF;

	
	
	--------   Para comercio   --------
	-- Catalogo de estados de comercio
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_estado_comercio')
	THEN
        CREATE TABLE ctl_estado_comercio(
            id_estado_comercio		VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlEdoComercio PRIMARY KEY(id_estado_comercio)
        );
	END IF;
	
	-- Catalogo de giros
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_giro')
	THEN
        CREATE TABLE ctl_giro(
            id_giro					VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlGiro PRIMARY KEY(id_giro)
        );
	END IF;
	
	-- Terminales
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='terminal')
	THEN
        CREATE TABLE terminal(
            id_terminal				VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_Terminal PRIMARY KEY(id_terminal)
        );
	END IF;
	
	-- Comercios
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='comercio')
	THEN
        CREATE TABLE comercio(
            id_comercio				VARCHAR(36) NOT NULL,
            razon_social 			VARCHAR(120) NOT NULL,
			rfc 					VARCHAR(13) NOT NULL,
			agente_comercial		VARCHAR(80),
			id_giro					VARCHAR(36) NOT NULL,
			id_estado_comercio		VARCHAR(36) NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_Comercio PRIMARY KEY(id_comercio)
        );

		ALTER TABLE comercio ADD CONSTRAINT FK_ComercioGiro FOREIGN KEY(id_giro)
            REFERENCES ctl_giro(id_giro);

		ALTER TABLE comercio ADD CONSTRAINT FK_ComercioEstado FOREIGN KEY(id_estado_comercio)
            REFERENCES ctl_estado_comercio(id_estado_comercio);
	END IF;

	
	
	--------   Para transaccion   --------
	-- Catalogo de tipos de iso
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_tipo_iso')
	THEN
        CREATE TABLE ctl_tipo_iso(
            id_tipo_iso				VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlTipoIso PRIMARY KEY(id_tipo_iso)
        );
	END IF;

	-- Catalogo de tipos de movimiento
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_tipo_movimiento')
	THEN
        CREATE TABLE ctl_tipo_movimiento(
            id_tipo_movimiento		VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlTipoMovimiento PRIMARY KEY(id_tipo_movimiento)
        );
	END IF;

	-- Catalogo de tipos de movimiento
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_tipo_servicio')
	THEN
        CREATE TABLE ctl_tipo_servicio(
            id_tipo_servicio		VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlTipoServicio PRIMARY KEY(id_tipo_servicio)
        );
	END IF;

	-- Catalogo de tipos de transaccion
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_tipo_transaccion')
	THEN
        CREATE TABLE ctl_tipo_transaccion(
            id_tipo_transaccion		VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlTipoTransaccion PRIMARY KEY(id_tipo_transaccion)
        );
	END IF;

	-- Catalogo de estados de transaccion
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='ctl_estado_transaccion')
	THEN
        CREATE TABLE ctl_estado_transaccion(
            id_estado_transaccion	VARCHAR(36) NOT NULL,
            descripcion	 			VARCHAR(25) NOT NULL,
			activo					BOOLEAN NOT NULL,
			creador					VARCHAR(36) NOT NULL,
			creacion				TIMESTAMP NOT NULL,
			modificador				VARCHAR(36),
			modificacion			TIMESTAMP,
            CONSTRAINT PK_CtlEdoTransaccion PRIMARY KEY(id_estado_transaccion)
        );
	END IF;

	-- Transacciones
	IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE tablename='transaccion')
	THEN
        CREATE TABLE transaccion(
            id_transaccion			VARCHAR(36) NOT NULL,
            autorizacion 			VARCHAR(25) NOT NULL,
			monto					NUMERIC(11,3) NOT NULL,
			id_estado_transaccion	VARCHAR(36) NOT NULL,
			concepto	 			VARCHAR(25) NOT NULL,
			referencia	 			VARCHAR(25) NOT NULL,
			id_terminal				VARCHAR(36) NOT NULL,
			id_tipo_transaccion		VARCHAR(36) NOT NULL,
			id_tipo_servicio		VARCHAR(36),
			id_comercio				VARCHAR(36) NOT NULL,
			id_tarjeta				VARCHAR(36) NOT NULL,
			id_tipo_movimiento		VARCHAR(36) NOT NULL,
			id_tipo_iso				VARCHAR(36),
			ticket					VARCHAR(20),
			comercial_asignado		VARCHAR(80),
			fecha					TIMESTAMP NOT NULL,
            CONSTRAINT PK_Transaccion PRIMARY KEY(id_transaccion)
        );

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionEdo FOREIGN KEY(id_estado_transaccion)
            REFERENCES ctl_estado_transaccion(id_estado_transaccion);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionTerminal FOREIGN KEY(id_terminal)
            REFERENCES terminal(id_terminal);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionTipo FOREIGN KEY(id_tipo_transaccion)
            REFERENCES ctl_tipo_transaccion(id_tipo_transaccion);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionTipoServicio FOREIGN KEY(id_tipo_servicio)
            REFERENCES ctl_tipo_servicio(id_tipo_servicio);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionComercio FOREIGN KEY(id_comercio)
            REFERENCES comercio(id_comercio);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionTarjeta FOREIGN KEY(id_tarjeta)
            REFERENCES tarjeta(id_tarjeta);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionTipoMovi FOREIGN KEY(id_tipo_movimiento)
            REFERENCES ctl_tipo_movimiento(id_tipo_movimiento);

		ALTER TABLE transaccion ADD CONSTRAINT FK_TransaccionTipoIso FOREIGN KEY(id_tipo_iso)
            REFERENCES ctl_tipo_iso(id_tipo_iso);
	END IF;

END;
$$
LANGUAGE 'plpgsql';

SELECT actualizaBD();
DROP FUNCTION actualizaBD();