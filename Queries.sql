-- 1.

select d.disciplina_id, d.nome, h.semestre, h.ano, h.nota_final 
from disciplina as d
right join historico_escolar as h
	on d.disciplina_id = h.disciplina_id
where aluno_id = 970783349016330241 --id do aluno a ser buscado

-- 2.

select * from disciplina_ministrada
where professor_id = 970783375471345665 --id do professor a ser buscado

-- 3. 

select a.nome, a.aluno_id, a.curso_id, mc.disciplina_id,
CASE
WHEN he.nota_final < 5 THEN null
ELSE he.nota_final
END AS nota_final, he.semestre,he.ano
from aluno AS a
join historico_escolar AS he ON a.aluno_id = he.aluno_id
join matriz_curricular AS mc ON a.curso_id = mc.curso_id AND mc.disciplina_id = he.disciplina_id
where he.semestre = 1 --semestre a ser buscado
AND he.ano = 2012 --ano a ser buscado
AND NOT EXISTS (
SELECT 1
FROM matriz_curricular AS mc2
LEFT JOIN historico_escolar AS he2 ON mc2.disciplina_id = he2.disciplina_id AND he2.aluno_id = a.aluno_id
where mc2.curso_id = a.curso_id
AND he2.nota_final IS NULL
)
ORDER BY
    a.aluno_id,
    mc.disciplina_id;

	
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
