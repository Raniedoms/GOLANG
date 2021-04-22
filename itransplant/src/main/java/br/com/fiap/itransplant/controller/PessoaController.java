package br.com.fiap.itransplant.controller;

import br.com.fiap.itransplant.exception.ResourceNotFoundException;
import br.com.fiap.itransplant.model.Pessoa;
import br.com.fiap.itransplant.repository.PessoaRepository;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("pessoa")
public class PessoaController {

    private PessoaRepository repository;

    public PessoaController(PessoaRepository repository) {
        this.repository = repository;
    }

    @GetMapping
    public List<Pessoa> listar(){
        return repository.findAll();
    }

    @GetMapping("/disponiveis")
    public List<Pessoa> pessoasDisponiveis() {
        return repository.listarPessoasDisponiveis();
    }

    @GetMapping("{id}")
    public Pessoa buscar(@PathVariable int id) {
        return repository.findById(id).orElseThrow(()->new ResourceNotFoundException());
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Pessoa cadastrar(@RequestBody Pessoa pessoa){
        pessoa.setAtivo("S");
        return repository.save(pessoa);
    }

    @PutMapping("{id}")
    public Pessoa atualizar(@PathVariable int id, @RequestBody Pessoa pessoa){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        pessoa.setId(id);
        return repository.save(pessoa);
    }

    @DeleteMapping("{id}")
    public void remover(@PathVariable int id){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        repository.deleteById(id);
    }
}
