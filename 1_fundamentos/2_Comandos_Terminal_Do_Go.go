package main

import "fmt"

func main() {
	fmt.Print("Comandos e para que servem!")
	fmt.Println("\n-------------------------------------")

	fmt.Println("go -->")
	fmt.Println("  Comando principal do Go. Exibe ajuda, subcomandos e controla todo o toolchain.")

	fmt.Println("\ngo help get -->")
	fmt.Println("  Mostra ajuda detalhada sobre o comando 'go get' (instalação de dependências).")

	fmt.Println("\ngo version -->")
	fmt.Println("  Exibe a versão atual do Go instalada no sistema.")

	fmt.Println("\ngodoc -http=:6060 -->")
	fmt.Println("  Inicia um servidor local com a documentação completa da linguagem Go.")
	fmt.Println("  Acesse no navegador: http://localhost:6060")

	fmt.Println("\ngo env -->")
	fmt.Println("  Lista variáveis de ambiente usadas pelo Go, como GOPATH, GOROOT e GOBIN.")

	fmt.Println("\ngo doc cmd/vet -->")
	fmt.Println("  Exibe a documentação do comando 'go vet', usado para análise estática do código.")

	fmt.Println("\ngo vet comandos.go -->")
	fmt.Println("  Analisa o código procurando padrões suspeitos e possíveis bugs não detectados pela compilação.")

	fmt.Println("\ngo build comandos.go -->")
	fmt.Println("  Compila o arquivo e gera um binário executável na pasta atual.")

	fmt.Println("\n./comandos -->")
	fmt.Println("  Executa o binário gerado pelo 'go build' (Linux/Mac).")

	fmt.Println("\ngo run comandos.go -->")
	fmt.Println("  Compila e executa o código imediatamente sem gerar binário.")

	fmt.Println("\nls ~/go/src/github.com -->")
	fmt.Println("  Lista os pacotes instalados no workspace GOPATH (Mac/Linux).")

	fmt.Println("\ndir ~/go/src/github.com -->")
	fmt.Println("  Lista os pacotes instalados no GOPATH (Windows).")

	fmt.Println("\ngo get -u github.com/go-sql-driver/mysql -->")
	fmt.Println("  Instala ou atualiza o driver MySQL no projeto (Go Modules).")

	// -------------------------------
	// Comandos COMPLEMENTARES (importantes)
	// -------------------------------

	fmt.Println("\ngo mod init <nome> -->")
	fmt.Println("  Cria um novo módulo Go no diretório atual (necessário em projetos Go modernos).")

	fmt.Println("\ngo mod tidy -->")
	fmt.Println("  Ajusta dependências: remove as não usadas e baixa as que faltam.")

	fmt.Println("\ngo mod download -->")
	fmt.Println("  Baixa todas as dependências listadas no go.mod.")

	fmt.Println("\ngo list ./... -->")
	fmt.Println("  Lista todos os pacotes do módulo atual, útil para automação e ferramentas.")

	fmt.Println("\ngo install <pacote>@latest -->")
	fmt.Println("  Compila e instala ferramentas/binários escritos em Go dentro do GOBIN.")

	fmt.Println("\ngo test -->")
	fmt.Println("  Executa testes automatizados do pacote atual.")

	fmt.Println("\ngo test ./... -->")
	fmt.Println("  Executa testes em todos os pacotes do projeto.")

	fmt.Println("\ngo fmt ./... -->")
	fmt.Println("  Formata automaticamente todos os arquivos Go no padrão oficial.")

	fmt.Println("\ngo clean -->")
	fmt.Println("  Remove arquivos temporários, cache e restos de builds.")

	fmt.Println("\ngo fix -->")
	fmt.Println("  Atualiza código antigo para sintaxe nova do Go.")

	fmt.Println("\ngo tool pprof -->")
	fmt.Println("  Ferramenta de análise de performance para CPU/memória (avançado).")

	fmt.Println("\n-------------------------------------")
	fmt.Println("Todos os comandos mais importantes do Go foram exibidos com sucesso!")

}
