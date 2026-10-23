package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/miekg/dns"
)

var listaBloqueados map[string]bool

func carregarSitesSafados(caminho string) (map[string]bool, error) {
	arquivo, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	listaSites := make(map[string]bool)
	leitor := bufio.NewScanner(arquivo)
	for leitor.Scan() {
		url := strings.TrimSpace(leitor.Text())
		if url != "" {
			listaSites[url] = true
		}
	}

	if err := leitor.Err(); err != nil {
		return nil, err
	}
	return listaSites, nil
}

func atendenteDeSites(w dns.ResponseWriter, r *dns.Msg) {
	resposta := new(dns.Msg)
	resposta.SetReply(r)
	resposta.Authoritative = true
	resposta.RecursionAvailable = true

	for _, q := range r.Question {
		dominioLimpo := strings.TrimSuffix(q.Name, ".")
		fmt.Printf("📥 Recebido pedido para: %s (Tipo: %d)\n", q.Name, q.Qtype)

		if listaBloqueados[dominioLimpo] {
			fmt.Printf("🚫 BLOQUEADO: %s -> Respondendo 0.0.0.0\n", dominioLimpo)
			registro, err := dns.NewRR(fmt.Sprintf("%s 300 IN A 0.0.0.0", q.Name))
			if err == nil {
				resposta.Answer = append(resposta.Answer, registro)
			}
			w.WriteMsg(resposta)
			return
		} else {
			fmt.Printf("✅ LIBERADO: %s -> Consultando Cloudflare (1.1.1.1)\n", dominioLimpo)
			cliente := new(dns.Client)
			respReal, _, err := cliente.Exchange(r, "1.1.1.1:53")
			if err == nil {
				w.WriteMsg(respReal)
				return
			} else {
				fmt.Printf("Erro ao consultar upstream: %v\n", err)
			}
		}
	}
	w.WriteMsg(resposta)
}

func iniciarServidor() {
	dns.HandleFunc(".", atendenteDeSites)

	servidor := &dns.Server{Addr: "127.0.0.1:8053", Net: "udp"}

	fmt.Println("Central ouvindo na porta 127.0.0.1:8053 UDP...")

	err := servidor.ListenAndServe()
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}

func main() {
	var err error
	listaBloqueados, err = carregarSitesSafados("blocklist.txt")
	if err != nil {
		fmt.Println("Erro ao carregar arquivo:", err)
		return
	}
	fmt.Printf("Sucesso: foram carregados %d sites na memória\n", len(listaBloqueados))

	iniciarServidor()
}
