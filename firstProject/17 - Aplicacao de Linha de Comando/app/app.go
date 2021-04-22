package app

import (
	"log"
	"fmt"
	"net"
	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicacao de linha de comando pronta para ser executada.

func Gerar() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicacao de Linha de Comando"
	app.Usage = "Busca IPs e Noes de Servidor na internet"

	//command e um struct
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPS de enderecos na internet",
			//Flags - parametro apra que esse comando funcione
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}

	return app
} 

func buscarIps(c *cli.Context){
	host := c.String("host")

	//net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips{
		fmt.Println(ip)
	}
}