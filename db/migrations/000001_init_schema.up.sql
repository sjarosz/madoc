CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(20) NOT NULL, 
    fName       VARCHAR(50) NOT NULL, 
    lName       VARCHAR(50) NOT NULL, 
    utype       INTEGER NOT NULL,
    created     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS username_idx ON users(username);
CREATE UNIQUE INDEX IF NOT EXISTS id_idx ON users(id);

CREATE TABLE IF NOT EXISTS appointments (
    apptId      SERIAL PRIMARY KEY,
    startTime   TIMESTAMP NOT NULL,
    endTime     TIMESTAMP NOT NULL,
    patient     VARCHAR(20) NOT NULL,
    status      INTEGER NOT NULL,
    createdBy   VARCHAR(20) NOT NULL,
    created     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users
        FOREIGN KEY(patient)
            REFERENCES users(username)
);

CREATE UNIQUE INDEX IF NOT EXISTS appt_idx ON appointments(apptId);

CREATE TABLE IF NOT EXISTS healthrecords (
    healthRecordId      SERIAL PRIMARY KEY,
    apptId              INTEGER,
    description         VARCHAR(200) NOT NULL,
    patient             VARCHAR(20) NOT NULL,
    createdBy           VARCHAR(20) NOT NULL,
    created             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_users
        FOREIGN KEY(patient)
            REFERENCES users(username),
    CONSTRAINT fk_appointments
        FOREIGN KEY(apptId)
            REFERENCES appointments(apptId)

);

CREATE UNIQUE INDEX IF NOT EXISTS health_record_idx ON healthrecords(healthRecordId);