create table aluno(
	aluno_id serial not null,
	nome varchar(255),
	curso_id int,
	data_ingresso varchar,
	primary key (aluno_id)
);

create table professor(
	professor_id serial not null,
	nome varchar,
	departamento_id int,
	primary key (professor_id)
);

create table curso(
	curso_id serial not null,
	nome varchar,
	departamento_id int,
	primary key (curso_id)
);

create table departamento(
	departamento_id serial not null,
	nome varchar,
	chefe_id int,
	primary key (departamento_id)
);

create table disciplina(
	disciplina_id serial not null,
	nome varchar,
	curso_id int,
	professor_id int,
	primary key (disciplina_id)
);

ALTER TABLE aluno
ADD CONSTRAINT curso_fk FOREIGN KEY (curso_id)
REFERENCES curso(curso_id)
ON DELETE CASCADE;

ALTER TABLE professor
ADD CONSTRAINT departamento_fk FOREIGN KEY (departamento_id)
REFERENCES departamento(departamento_id)
ON DELETE CASCADE;

ALTER TABLE curso
ADD CONSTRAINT departamento_fk FOREIGN KEY (departamento_id)
REFERENCES departamento(departamento_id)
ON DELETE CASCADE;

ALTER TABLE departamento
ADD CONSTRAINT chefe_fk FOREIGN KEY (chefe_id)
REFERENCES professor(professor_id)
ON DELETE CASCADE;

ALTER TABLE disciplina
ADD CONSTRAINT curso_fk FOREIGN KEY (curso_id)
REFERENCES curso(curso_id)
ON DELETE CASCADE;

ALTER TABLE disciplina
ADD CONSTRAINT professor_fk FOREIGN KEY (professor_id)
REFERENCES professor(professor_id)
ON DELETE CASCADE;

