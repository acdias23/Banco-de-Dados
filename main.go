package main

import (
	"database/sql"
	"log"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"github.com/jaswdr/faker"
	_ "github.com/lib/pq"
)

type Aluno struct {
	Name string
	Date string
}

var cursos = []string{
	"Administracao",
	"Ciencia da Computacao",
	"Engenharia Civil",
	"Engenharia Eletrica",
}

var cursos_id = [4]int{}
var departamentos_id = [5]int{}
var disciplinas_adm_id = [40]int{}
var disciplinas_cc_id = [40]int{}
var disciplinas_engc_id = [40]int{}
var disciplinas_enge_id = [40]int{}
var alunos_id = [20]int{}
var professores_id = [15]int{}

func pickRandom(arr []int, n int) int {
	idx := rand.IntN(n)
	return arr[idx]
}

func randYear() int {
	min := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2020, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int64N(delta) + min
	return time.Unix(sec, 0).Year()
}

func main() {
	connStr := "" //Coloque aqui sua string de conexão
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	m := make(map[string][]string)

	m["Administracao"] = []string{
		"Matematica aplicada a administracao",
		"Fundamentos da administracao",
		"Estudos em macroeconomia",
		"Sociologia I",
		"Sociologia II",
		"Linguagens e generos textuais",
		"Etica nas organizacoes",
		"Ensino social cristao",
		"Calculo basico I",
		"Calculo basico II",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia I",
		"Filosofia II",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Gramatica",
		"Estudos em microeconomia",
		"Sustentabilidade",
		"Probabilidade estatistica",
		"Estudos do MEI",
		"Gestao de negocios",
		"Gestao de pessoas",
		"Gesta Financeira",
		"Matematica Financeira",
		"Gestao de recusos humanos",
		"PowerBI",
		"Excel Basico",
		"Excel intermediario",
		"Excel avancado",
		"Redacao",
		"Comunicacao para empresas",
		"Estrategia de marketing",
		"Direito Tributario",
		"Logistica",
		"Implementacao de negocios",
	}

	m["Ciencia da Computacao"] = []string{
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Sociologia",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Circuitos eletricos",
		"Ecologia e sustentabilidade",
		"Sistemas embarcados",
		"Redes de computadores",
		"Calculo numerico",
		"Termodinamica",
		"Cinematica e dinamica",
		"Algebra Linear",
		"Automatos",
		"Redes moveis",
		"Algoritmos",
		"Sistema Operacional",
		"Computacao grafica",
		"Compiladores",
		"Empreendedorismo",
		"Inovacao",
		"Teste de software",
		"IA",
		"Orientacao a objetos",
		"Desenvolvimento de jogos",
		"Robotica",
		"TCC I",
		"TCC II",
	}

	m["Engenharia Civil"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"Arquitetura e representacao grafica",
		"Mecanica geral",
		"Topografia",
		"Eletricidade geral",
		"Instalacoes eletricas",
		"Economia",
		"Geotecnia I",
		"Geotecnia II",
		"Transportes I",
		"Transportes II",
		"Optativa I",
		"Optativa II",
		"Custos",
		"Tecnologia das construcoes",
		"Metodos estatisticos",
		"Termodinamica",
		"Gestao de obras",
		"Planejamento de obras",
	}

	m["Engenharia Eletrica"] = []string{
		"Ensino social cristao",
		"Calculo basico",
		"Modelos organizacionais",
		"Contabilidade financeira",
		"Estudos em microeconomia",
		"Filosofia",
		"Linguagem em comunicacao organizacional",
		"Calculo diferencial e integral I",
		"Fundamentos de algoritmos",
		"Desenvolvimento web",
		"Comunicacao e expressao",
		"Calculo diferencial e integral II",
		"Calculo diferencial e integral III",
		"Calculo diferencial e integral IV",
		"Geometria analitica",
		"Teoria dos grafos",
		"Desenvolvimento de algoritmos",
		"Fisica",
		"Desenho tecnico",
		"Praticas de inovacao",
		"Eletronica geral",
		"Quimica",
		"sinais e Sistemas",
		"Praticas de inovacao I",
		"Praticas de inovacao II",
		"Praticas de inovacao III",
		"Circuitos integrados",
		"Mecanica Geral",
		"Gestao organizacional",
		"TCC 1",
		"TCC 2",
		"Redes",
		"Mecanica dos solidos",
		"Mecanica dos fluidos",
		"Sistemas eletricos",
		"Seguranca do trabalho",
		"Economia",
		"Conversao de energia I",
		"Conversao de energia II",
		"Conversao de energia III",
	}

	departamentos := [5]string{
		"Matematica",
		"Ciencia da Computacao",
		"Quimica",
		"Fisica",
		"Engenharia",
	}

	fakeName := faker.New().Person()
	fakeDate := faker.New().Time()

	var pk int
	for i := 0; i < 20; i++ {
		pk = insertAluno(db, Aluno{fakeName.Name(), fakeDate.UnixDate(time.Now())})
		alunos_id[i] = pk
	}

	for i := 0; i < 15; i++ {
		pk = insertProfessor(db, fakeName.Name())
		professores_id[i] = pk
	}

	for i, curso := range cursos {
		pk = insertCurso(db, curso)
		cursos_id[i] = pk
	}

	for i, departamento := range departamentos {
		pk = insertDeparatmento(db, departamento)
		departamentos_id[i] = pk
	}

	// for i, disciplina := range disciplinas {
	// 	pk = insertDisciplina(db, disciplina)
	// 	disciplinas_id[i] = pk
	// }

	// ADM

	for i, disc := range m["Administracao"] {
		pk = insertDisciplina(db, disc)
		disciplinas_adm_id[i] = pk
	}

	for i, disc := range m["Ciencia da Computacao"] {
		pk = insertDisciplina(db, disc)
		disciplinas_cc_id[i] = pk
	}

	for i, disc := range m["Engenharia Civil"] {
		pk = insertDisciplina(db, disc)
		disciplinas_engc_id[i] = pk
	}

	for i, disc := range m["Engenharia Eletrica"] {
		pk = insertDisciplina(db, disc)
		disciplinas_enge_id[i] = pk
	}

	atualizarFKAluno(db)
	atualizarFKCurso(db)
	atualizarFKDepartamento(db)
	atualizarFKDisciplina(db, disciplinas_adm_id[:], 0)
	atualizarFKDisciplina(db, disciplinas_cc_id[:], 1)
	atualizarFKDisciplina(db, disciplinas_engc_id[:], 2)
	atualizarFKDisciplina(db, disciplinas_enge_id[:], 3)
	atualizarFKProfessor(db)

	// Preencher Tabelas

	for i, disc := range disciplinas_adm_id {
		query := `INSERT INTO disciplina_ministrada (professor_id, disciplina_id, semestre, ano)
			  VALUES ($1, $2, $3, $4)`

		professor_id := pickRandom(professores_id[:], len(professores_id))
		semestre := int(i/5) + 1
		ano := randYear()

		db.Exec(query, professor_id, disc, semestre, ano)
	}

	for i, disc := range disciplinas_cc_id {
		query := `INSERT INTO disciplina_ministrada (professor_id, disciplina_id, semestre, ano)
			  VALUES ($1, $2, $3, $4)`

		professor_id := pickRandom(professores_id[:], len(professores_id))
		semestre := int(i/5) + 1
		ano := randYear()

		db.Exec(query, professor_id, disc, semestre, ano)
	}

	for i, disc := range disciplinas_engc_id {
		query := `INSERT INTO disciplina_ministrada (professor_id, disciplina_id, semestre, ano)
			  VALUES ($1, $2, $3, $4)`

		professor_id := pickRandom(professores_id[:], len(professores_id))
		semestre := int(i/5) + 1
		ano := randYear()

		db.Exec(query, professor_id, disc, semestre, ano)
	}

	for i, disc := range disciplinas_enge_id {
		query := `INSERT INTO disciplina_ministrada (professor_id, disciplina_id, semestre, ano)
			  VALUES ($1, $2, $3, $4)`

		professor_id := pickRandom(professores_id[:], len(professores_id))
		semestre := int(i/5) + 1
		ano := randYear()

		db.Exec(query, professor_id, disc, semestre, ano)
	}

	for i := 0; i < 6; i++ {
		query := `INSERT INTO grupo_tcc (aluno1_id, aluno2_id, aluno3_id, orientador)
			  VALUES ($1, $2, $3, $4)`

		aluno1 := pickRandom(alunos_id[:], len(alunos_id))
		aluno2 := pickRandom(alunos_id[:], len(alunos_id))
		aluno3 := pickRandom(alunos_id[:], len(alunos_id))
		orientador := pickRandom(professores_id[:], len(professores_id))

		db.Exec(query, aluno1, aluno2, aluno3, orientador)
	}

	// Gerar historico para cada aluno

	for _, aluno := range alunos_id {
		query := `INSERT INTO historico_escolar (aluno_id, disciplina_id, semestre, ano, nota_final)
			  VALUES ($1, $2, $3, $4, $5)`
		// preencher disciplina para os quatro anos
		var dataEntrada string
		var curso int
		db.QueryRow("SELECT data_ingresso FROM aluno WHERE aluno_id = ($1)", aluno).Scan(&dataEntrada)
		db.QueryRow("SELECT curso_id FROM aluno WHERE aluno_id = ($1)", aluno).Scan(&curso)

		date := strings.Split(dataEntrada, " ")
		ano, _ := strconv.ParseInt(date[len(date)-1], 10, 64)

		var d [40]int
		switch curso {
		case cursos_id[0]:
			d = disciplinas_adm_id
		case cursos_id[1]:
			d = disciplinas_cc_id
		case cursos_id[2]:
			d = disciplinas_engc_id
		case cursos_id[3]:
			d = disciplinas_enge_id
		}

		for i, disc := range d {
			if _, err := db.Exec(
				query, aluno, disc,
				int(i/5)+1, ano+int64(i/10), rand.Float32()*10); err != nil {
				log.Fatal(err)
			}
		}
	}

	for _, curso := range cursos_id {
		query := `INSERT INTO matriz_curricular (curso_id, disciplina_id, semestre)
			  VALUES ($1, $2, $3)`

		var d [40]int
		switch curso {
		case cursos_id[0]:
			d = disciplinas_adm_id
		case cursos_id[1]:
			d = disciplinas_cc_id
		case cursos_id[2]:
			d = disciplinas_engc_id
		case cursos_id[3]:
			d = disciplinas_enge_id
		}
		for i, disc := range d {
			if _, err := db.Exec(
				query, curso, disc, int(i/5)+1); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func insertProfessor(db *sql.DB, nome string) int {
	query := `INSERT INTO professor (nome)
          VALUES ($1) RETURNING professor_id`

	var pk int
	err := db.QueryRow(query, nome).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func atualizarFKProfessor(db *sql.DB) {
	query := `update professor 
			  set departamento_id = ($1)
			  where professor_id = ($2);`

	for _, id := range professores_id {
		db.Exec(query, pickRandom(departamentos_id[:], len(departamentos_id)), id)
	}
}

func insertAluno(db *sql.DB, aluno Aluno) int {
	query := `INSERT INTO aluno (nome, data_ingresso)
          VALUES ($1, $2) RETURNING aluno_id`

	var pk int
	err := db.QueryRow(query, aluno.Name, aluno.Date).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func atualizarFKAluno(db *sql.DB) {
	query := `update aluno 
			  set curso_id = ($1)
			  where aluno_id = ($2);`

	for _, id := range alunos_id {
		db.Exec(query, pickRandom(cursos_id[:], len(cursos_id)), id)
	}
}

func insertCurso(db *sql.DB, nome string) int {
	query := `INSERT INTO curso (nome)
          VALUES ($1) RETURNING curso_id`

	var pk int
	err := db.QueryRow(query, nome).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func atualizarFKCurso(db *sql.DB) {
	query := `update curso 
			  set departamento_id = ($1)
			  where curso_id = ($2);`

	for _, id := range cursos_id {
		db.Exec(query, pickRandom(departamentos_id[:], len(departamentos_id)), id)
	}
}

func insertDeparatmento(db *sql.DB, nome string) int {
	query := `INSERT INTO departamento (nome)
          VALUES ($1) RETURNING departamento_id`

	var pk int
	err := db.QueryRow(query, nome).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func atualizarFKDepartamento(db *sql.DB) {
	query := `update departamento 
			  set chefe_id = ($1)
			  where departamento_id = ($2);`

	for _, id := range departamentos_id {
		db.QueryRow(query, pickRandom(professores_id[:], len(professores_id)), id)
	}
}

func insertDisciplina(db *sql.DB, nome string) int {
	query := `INSERT INTO disciplina (nome)
          VALUES ($1) RETURNING disciplina_id`

	var pk int
	err := db.QueryRow(query, nome).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func atualizarFKDisciplina(db *sql.DB, arr []int, idx int) {
	query := `update disciplina 
			  set curso_id = ($1), professor_id = ($2)
			  where disciplina_id = ($3);`

	for _, id := range arr {
		db.QueryRow(query, cursos_id[idx],
			pickRandom(professores_id[:], len(professores_id)), id)
	}
}
