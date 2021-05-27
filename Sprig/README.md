Esse repositório contem funções do GO Template

**Links para ajuda:**
- https://golang.org/pkg/text/template/
- https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet


Usamos a biblioteca de funções spring no adaptador tmplGO para adicionar algumas funções úteis. Muitos dos seus problemas podem ser resolvidos usando algumas dessas funções :
- https://masterminds.github.io/sprig/


**EXEMPLOS**

**Regex**

regexFindAll

Usando a função regex para filtrar apenas caracteres numéricos. Usaremos a função regexFindAll.
Existem outras funções regex que você pode usar: https://masterminds.github.io/sprig/strings.html
A função regex retorna uma matriz. Então, podemos usar a função de junção para juntá-los.

code: 
```
{
	"regex": "{{regexFindAll "[0-9]" .commercialCode -1}}",
	"join":"{{(regexFindAll "[0-9]" .commercialCode -1) | join ""}}"
}
```
input:
```
{
	"commercialCode": "ABC123"
}
```
output:
```
{
	"regex": "[1 2 3]",
	"join": "123"
}
```
**regexMatch**

regexMatch retorna um booleano.

O exemplo abaixo mostra como verificar se um número de telefone possui o código de país 55 (DDI)

code:
```
{
	"regex": {{regexMatch "^5{2}" .numero}}	
}
```
input:
```
{
	"numero":"5511980266122"
}
```
output:
```
{
  "regex": true
}
```
**Replace**

code:
```
{{.PathParameters.campo | replace "-" ""}}
```
input:
```
{
	"PathParameters":{
		"campo": "2020-08-19"
	}
}
```
output:
```
20200819
```

**Criação de suas funções**

Definir uma função chamada "país" para traduzir o nome de um país pelo seu código ISO.
Podemos passar essa função no final de todo o template

code:

```
{"country":"{{template "country" .country}}"}

{{- define "country"}}
    {{- $out := "NONE"}}
    {{- if eq . "Argentina"}}{{- $out = "AR"}}
    {{- else if eq . "Brasil"}}{{- $out = "BR"}}
    {{- end}}
        {{- $out}}
{{- end}}
```
input:
```
{
    "country": "Brasil"
}
```
output:
```
{
  "country": "BR"
}
```

**FUNÇÕES COM DATAS**
Soma os dias em alguma data

"2006-01-02" é uma máscara de data. - https://www.pauladamsmith.com/blog/2011/05/go_time.html

code:
```
{{$date := toDate "2006-01-02" .date}}

{
  "date": "{{$date | date_modify "+720h"}}"
}
```
input:
```
{
  "date":"2019-07-27"
}
```
output:
```
{
  "date": "2019-08-26 00:00:00 +0000 UTC"
}
```

**Date to Epoch**

code:
```
{"ativacaoData": "{{toDate "2006-01-02" .item | unixEpoch}}"}
```
input:
```
{"item": "1991-03-12"}
```
Output:
```
{
  "ativacaoData": "668746800"
}
```
**Date "YYYY-MM-DD" para "DD/MM/YYYY"**

code:
```
{{toDate "2006-01-02" .date | date "02/01/2006"}}
```
input:
```
{
    "date":"2020-12-25"
}
```
Output:
```
"25/12/2020"
```
**Sysdate no formato DD/MM/YYYY HH24:MI:SS**


now - pega a data atual, no formato 02/01/2006

code:
```
{{now | date "02/01/2006 15:04:05"}}
```
output:
```
"19/09/2020 17:53:25"
```

**Date YYYY-MM-DDTHH24:Mi:SS.SSSZ to YYYYMMDDHHMiSS**

Exemplo: converta 2020-07-31T23:59:59.999Z para 20200731235959

code:
```
{{toDate "2006-01-02T15:04:05.000Z" .PathParameters.initialDate | date "20060102150405"}}
```
input:
```
{
	"PathParameters": {
		"finalDate": "2020-07-31T23:59:59.999Z",
	}
}
```
output:
```
20200731235959
```
**Date YYYY-MM-DDTHH24:Mi:SS.SSSZ to DDMMAAAAHHMISSFFFFFF**

Exemplo: converta 2020-08-17T19:49:59.139Z para 17082020194959000000
atribui uma variavel $date e já imputei as transformações e valores dela

code:
```
{{- $date := toDate "2006-01-02T15:04:05.000Z" .Body.date | date "02012006150405000000" }}

{{$date}}
```
input:
```
{
  "Body":{
    "date":"2020-08-17T19:49:59.139Z"
  }
}
```
output:
```
17082020194959000000
```
**Pegando o tamanho da String e verificando se é menor ou igual a 11 caracteres**

*$tamanho* é a variável que irá pegar a quantidade de caracteres dentro da tag mapeada usando a função: len

code:
```
{{$tamanho := len .quantidade}}

{{- if le $tamanho 11}}
    {
        "validate":"true", 
        "tamanho":"{{toJson $tamanho}}"
    }
{{else}}
    {
        "validate":"false",
        "tamanho":"{{toJson $tamanho}}"
    }
{{end}}
```
input:
```
{
    "quantidade":"12345678901"
}
```
Output:
```
{
  "validate": "true",
  "tamanho": "11"
}
```

**Makelist e Range**
*Makelist* - ele verifica se a estrutura esta vindo dentro de um array ou não. Caso não venha um array, ele transforma em um para que o range consiga fazer a função de lista
*Range* - só funcionará quando vier um array de objetos
*list* - ele inputará todos os valores adicionados no append em uma lista
*toJson* - imprimindo em estrutura Json

code:
```
{{$ClientList := list}}
{{$listaClientes := makeList .Cliente.Clientes}}
    {{range $index, $usuario := $listaClientes}}
        {{$c := initTag}}    
        {{ addTag $c "chave" $usuario.key}}
        {{ addTag $c "valor" $usuario.value}}
        {{$ClientList = append $ClientList $c}} 
    {{end}}
{{end}}


{{toJson $ClientList}}
```
input:
```
{
	"Cliente": {
		"Clientes": [{
			"key": "A1",
			"value": "1"
		}, {
			"key": "A2",
			"value": "2"
		}, {
			"key": "A3",
			"value": "3"
		}, {
			"key": "A4",
			"value": "4"
		}]
	}
}
```
output:
```
[{
	"key": "A1",
	"value": "1"
}, {
	"key": "A2",
	"value": "2"
}, {
	"key": "A3",
	"value": "3"
}, {
	"key": "A4",
	"value": "4"
}]
```
**Contando os elementos de uma lista**

code:
```
{{$variavel := len .preferencia | int }}

{{if gt $variavel 10 }}
    {
        "quantidade": {{$variavel}},
        "validacao": "true"
    }
{{else}}
    {
        "quantidade": {{$variavel}},
        "validacao": "false"
    }
{{end}}
```
input:
```
{
	"preferencia": [
		"id",
		"cpf",
		"tel",
		"endereco",
		"numero"
	]
}
```
output:
```
{
	"quantidade": 5,
	"validacao": "false"
}

```
**Valor entre datas**

code:
```
{{- $dateNow := now | date "2006-01-02" }}
{{- $date := "2020-12-31"}}

{
"dateNow": "{{$dateNow}}",
"date": "{{$date}}",
{{$period := getPeriod "2006-01-02"  $dateNow $date }}
"period": "{{$period}} 

dias para o fim do ano (⊙_⊙)"
}
```
output:
```
{
  "dateNow": "2020-09-29",
  "date": "2020-12-31",
  "period": "93 dias para o fim do ano (⊙_⊙)"
}
```
**Comparando valores com GT - Gr**

* eq - Returns the boolean truth of arg1 == arg2
* ne - Returns the boolean truth of arg1 != arg2
* lt - Returns the boolean truth of arg1 < arg2
* le - Returns the boolean truth of arg1 <= arg2
* gt -	Returns the boolean truth of arg1 > arg2
* ge -	Returns the boolean truth of arg1 >= arg2

Verifica se um elemento é maior do que outro. Ex: arg1> arg2.

code:
```
{{$variable := 11}}

{{if gt $variable 10 }}
    {
        "requestValidate": "true"
    }
{{else}}
    {
        "requestValidate": "false"
    }
{{end}}
```
output:
```
{
  "requestValidate": "true"
}

```
**Adicionando 30 dias no dia atual (SysDate)**

code:
```
{{$dateNow := now | date "2006-01-02T15:04:05.000"}}

{{$date := toDate "2006-01-02T15:04:05.000" $dateNow}}
{{$date = $date | date_modify "+720h"}}


{
    "data atual": "{{$dateNow}}",
    "data mais trinta dias": "{{$date | date "2006-01-02T15:04:05.000Z"}}"
}
```

**Verificando se o último elemento é número, sendo que pode vir : 91410120456F8**

code:
```
{{$num:= len .numero }} 
{{$num1:=  sub $num 1 }} 
{{$var := substr ($num1 | int) ($num | int) .numero }}
{{ $var1 := regexMatch "[0-9]" $var }}
{{ $var1 }}
```
input:
```
{
    "numero":"7323322F2"
}
```
output:
```
true
```

**Calcular o valor retornado + SYSDATE. No formato: YYYY-MM-DDTHH24:Mi:SS.SSSZ**

code 1:
```
{{$dateNow := now | date "2006-01-02T15:04:05.000"}}
{{$expDate := 30}}
{{$daysHours := 24}}

{{$result := mul $expDate $daysHours | int64}}
{{$result2 := (cat "+" $result "h" | nospace) }}
{{$date := toDate "2006-01-02T15:04:05.000" $dateNow}}
{{$date = $date | date_modify $result2 }}

{{toJson $date}}
```
code 2:
```
{{$expDate := 30}}

{{$dateNow := now | date "2006-01-02T15:04:05.000"}}
{{$date := toDate "2006-01-02T15:04:05.000" $dateNow}}
{{$hours := (mul ($expDate | int ) 24 ) | toString}}
{{$expDate = $date | date_modify (printf "+%sh" $hours) | date "2006-01-02T15:04:05.000Z"}}

{{toJson $expDate}}
```