package br.com.fiap.itransplant.controller;

import br.com.fiap.itransplant.exception.ResourceNotFoundException;
import br.com.fiap.itransplant.model.CadastroOrgao;
import br.com.fiap.itransplant.model.Orgao;
import br.com.fiap.itransplant.repository.CadastroOrgaoRepository;
import br.com.fiap.itransplant.repository.OrgaoRepository;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("cadastroOrgao")
public class CadastroOrgaoController {

    private CadastroOrgaoRepository repository;

    public CadastroOrgaoController(CadastroOrgaoRepository repository) {
        this.repository = repository;
    }


    @GetMapping("/orgao")
    public List<CadastroOrgao> listarOrgaoParaTransplante() {
        return repository.getOrgaoCadastradoNaoUsados();
    }

    @GetMapping
    public List<CadastroOrgao> listar(){
        return repository.findAll();
    }

    @GetMapping("{id}")
    public CadastroOrgao buscar(@PathVariable int id) {
        return repository.findById(id).orElseThrow(()->new ResourceNotFoundException());
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public CadastroOrgao cadastrar(@RequestBody CadastroOrgao cadastroOrgao){
        return repository.save(cadastroOrgao);
    }

    @PutMapping("{id}")
    public CadastroOrgao atualizar(@PathVariable int id, @RequestBody CadastroOrgao cadastroOrgao){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        cadastroOrgao.setId(id);
        return repository.save(cadastroOrgao);
    }

    @DeleteMapping("{id}")
    public void remover(@PathVariable int id){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        repository.deleteById(id);
    }
}
