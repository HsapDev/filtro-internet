package main

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
)

func main() {
	dominio := "xvideos.com"
	if len(os.Args) > 1 {
		dominio = os.Args[1]
	}

	fmt.Printf("🔍 Testando consulta para: %s no nosso servidor local (127.0.0.1:8053)...\n", dominio)

	mensagem := new(dns.Msg)
	mensagem.SetQuestion(dns.Fqdn(dominio), dns.TypeA)

	cliente := new(dns.Client)
	resposta, rtt, err := cliente.Exchange(mensagem, "127.0.0.1:8053")

	if err != nil {
		fmt.Printf("❌ Erro de conexão: %v\n", err)
		return
	}

	fmt.Printf("⚡ Resposta recebida em %v!\n", rtt)
	for _, respostaItem := range resposta.Answer {
		if registroA, ok := respostaItem.(*dns.A); ok {
			ip := registroA.A.String()
			if ip == "0.0.0.0" {
				fmt.Printf("🛑 BLOQUEIO CONFIRMADO: %s -> IP: %s\n", dominio, ip)
			} else {
				fmt.Printf("🌐 SITE LIBERADO: %s -> IP: %s\n", dominio, ip)
			}
		}
	}
}
