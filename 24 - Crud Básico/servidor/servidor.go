package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type usuario struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario cria um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para Struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao inserir usuário"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro obter o ID inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso, id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuários"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro no encode"))
		return
	}
}

// BuscarUsuario busca um usuário no banco
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o param para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro escanear usuário"))
			return
		}
	}

	if usuario.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado!"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
			w.Write([]byte("Erro no encode"))
			return
		}
	}
}

// AtualizarUsuario Altera os dados de um usuário no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o param para inteiro"))
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o param para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao preparar stetament"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
