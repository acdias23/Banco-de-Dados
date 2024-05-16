-- 1.

select d.disciplina_id, d.nome, h.semestre, h.ano, h.nota_final 
from disciplina as d
right join historico_escolar as h
	on d.disciplina_id = h.disciplina_id
where aluno_id = 968549803004461057 --id do aluno a ser buscado

-- 2.

select * from disciplina_ministrada
where professor_id = 968549813298003969

-- 3. 

select a.nome,a.aluno_id, mc.curso_id, mc.disciplina_id
from aluno as a
left join matriz_curricular mc
	on a.curso_id = mc.curso_id 
left join historico_escolar he
	on he.disciplina_id = mc.disciplina_id
	and a.aluno_id = he.aluno_id 
where 
	he.ano = 2012
	and he.semestre = 1
	
-- 4.

select p.professor_id, p.nome, d.nome
from professor as p
right join departamento as d 
	on p.professor_id = d.chefe_id

-- 5.
	
select gt.tcc_id, a.nome, p.nome 
from aluno as a
right join grupo_tcc as gt
	on aluno_id = gt.aluno1_id
	or aluno_id = gt.aluno2_id
	or aluno_id = gt.aluno3_id
right join professor as p
	on p.professor_id = gt.orientador 
where a.nome is not null
