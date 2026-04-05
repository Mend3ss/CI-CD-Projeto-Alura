package database

import (
	"log"

	//O "os" importa funcionalidades do sistema operacional
	"os"
	
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	
	//Aqui definimos as variaveis que usaremos
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")
	
	//Agora a stringDeConexao ira puxar os valores acima
	stringDeConexao := "host=" + endereco + " user=" + usuario + " password=" + senha + " dbname=" + nomeBanco + " port=" + portaBanco + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao se conectar com o banco de dados")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}
