package br.com.fiap.itransplant.repository;

import br.com.fiap.itransplant.model.Pessoa;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface PessoaRepository extends JpaRepository<Pessoa, Integer> {

    @Query(value = "SELECT * FROM tb_pessoa WHERE fl_ativo = 'S' ORDER BY dt_entrada asc ", nativeQuery = true)
    List<Pessoa> listarPessoasDisponiveis();

}
