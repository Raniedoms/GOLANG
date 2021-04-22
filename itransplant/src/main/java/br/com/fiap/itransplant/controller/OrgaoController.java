package br.com.fiap.itransplant.controller;


import br.com.fiap.itransplant.exception.ResourceNotFoundException;
import br.com.fiap.itransplant.model.Orgao;
import br.com.fiap.itransplant.repository.OrgaoRepository;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("orgao")
public class OrgaoController {

    private OrgaoRepository repository;

    public OrgaoController(OrgaoRepository repository) {
        this.repository = repository;
    }


    @GetMapping
    public List<Orgao> listar(){
        return repository.findAll();
    }

    @GetMapping("{id}")
    public Orgao buscar(@PathVariable int id) {
        return repository.findById(id).orElseThrow(()->new ResourceNotFoundException());
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Orgao cadastrar(@RequestBody Orgao orgao){
        return repository.save(orgao);
    }

    @PutMapping("{id}")
    public Orgao atualizar(@PathVariable int id, @RequestBody Orgao orgao){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        orgao.setId(id);
        return repository.save(orgao);
    }

    @DeleteMapping("{id}")
    public void remover(@PathVariable int id){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        repository.deleteById(id);
    }
}
