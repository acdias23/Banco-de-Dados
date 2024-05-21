create table historico_escolar (
	aluno_id int,
	disciplina_id int,
	semestre int,
	ano int,
	nota_final float,
	foreign key (aluno_id) references aluno(aluno_id),
	foreign key (disciplina_id) references disciplina(disciplina_id)
)

create table disciplina_ministrada (
	professor_id int,
	disciplina_id int,
	semestre int,
	ano int,
	foreign key (professor_id) references professor(professor_id),
	foreign key (disciplina_id) references disciplina(disciplina_id)
)

create table matriz_curricular (
	curso_id int,
	disciplina_id int,
	semestre int,
	foreign key (curso_id) references curso(curso_id),
	foreign key (disciplina_id) references disciplina(disciplina_id)
)

-- o grupo é obrigatório ser composto por exatamente 3 alunos
create table grupo_tcc (
	tcc_id serial not null PRIMARY KEY,
	aluno1_id int,
	aluno2_id int,
	aluno3_id int,
	orientador int,
	foreign key (aluno1_id) references aluno(aluno_id),
	foreign key (aluno2_id) references aluno(aluno_id),
	foreign key (aluno3_id) references aluno(aluno_id),
	foreign key (orientador) references professor(professor_id)
)

delete from matriz_curricular;
delete from historico_escolar; 
delete from grupo_tcc;
delete from disciplina_ministrada;
delete from disciplina;
delete from departamento;
delete from curso;
delete from aluno;
delete from professor
