package br.com.fiap.itransplant.repository;

import br.com.fiap.itransplant.model.Endereco;
import org.springframework.data.jpa.repository.JpaRepository;

public interface EnderecoRepository extends JpaRepository<Endereco, Integer> {
}
