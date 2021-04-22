package br.com.fiap.itransplant.controller;

import br.com.fiap.itransplant.exception.ResourceNotFoundException;
import br.com.fiap.itransplant.model.Transplante;
import br.com.fiap.itransplant.repository.TransplanteRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Slf4j
@RestController
@RequestMapping("transplante")
public class TransplanteController {

    private TransplanteRepository repository;

    public TransplanteController(TransplanteRepository repository) {
        this.repository = repository;
    }

    @GetMapping
    public List<Transplante> listar() {
        return repository.findAll();
    }

    @GetMapping("{id}")
    public Transplante buscar(@PathVariable int id) {
        return repository.findById(id).orElseThrow(() -> new ResourceNotFoundException());
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public Transplante cadastrar(@RequestBody Transplante transplante) {
        try {
            this.repository.inativarPessoa(transplante.getPessoa().getId());
            log.info("Pessoa Inativada ");
        } catch (Exception e) {
            log.error("Erro ao inativar pessoa", e.getMessage());
        }
        return repository.save(transplante);
    }

    @PutMapping("{id}")
    public Transplante atualizar(@PathVariable int id, @RequestBody Transplante transplante) {
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        transplante.setId(id);
        return repository.save(transplante);
    }

    @DeleteMapping("{id}")
    public void remover(@PathVariable int id) {
        repository.findById(id).orElseThrow(ResourceNotFoundException::new);
        repository.deleteById(id);
    }
}
