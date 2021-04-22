package br.com.fiap.itransplant.controller;

import br.com.fiap.itransplant.exception.ResourceNotFoundException;
import br.com.fiap.itransplant.model.Endereco;
import br.com.fiap.itransplant.model.Orgao;
import br.com.fiap.itransplant.repository.EnderecoRepository;
import br.com.fiap.itransplant.repository.OrgaoRepository;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("endereco")
public class EnderecoController {

    private EnderecoRepository repository;

    public EnderecoController(EnderecoRepository repository) {
        this.repository = repository;
    }


    @GetMapping
    public List<Endereco> listar(){
        return repository.findAll();
    }

    @GetMapping("{id}")
    public Endereco buscar(@PathVariable int id) {
        return repository.findById(id).orElseThrow(()->new ResourceNotFoundException());
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Endereco cadastrar(@RequestBody Endereco endereco){
        return repository.save(endereco);
    }

    @PutMapping("{id}")
    public Endereco atualizar(@PathVariable int id, @RequestBody Endereco endereco){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        endereco.setCodigoEndereco(id);
        return repository.save(endereco);
    }

    @DeleteMapping("{id}")
    public void remover(@PathVariable int id){
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        repository.deleteById(id);
    }
}
