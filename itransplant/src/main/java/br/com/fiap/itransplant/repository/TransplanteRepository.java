package br.com.fiap.itransplant.repository;

import br.com.fiap.itransplant.model.CadastroOrgao;
import br.com.fiap.itransplant.model.Transplante;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface TransplanteRepository extends JpaRepository<Transplante, Integer> {

    @Transactional
    @Modifying
    @Query(value = "UPDATE tb_pessoa SET fl_ativo = 'N' WHERE id_pessoa = :id ",
            nativeQuery = true)
    void inativarPessoa(@Param("id") int id);
}
