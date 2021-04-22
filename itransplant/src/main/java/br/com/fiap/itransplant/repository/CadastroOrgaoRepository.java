package br.com.fiap.itransplant.repository;

import br.com.fiap.itransplant.model.CadastroOrgao;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface CadastroOrgaoRepository extends JpaRepository<CadastroOrgao, Integer> {

    @Query(value = "SELECT * FROM tb_cadt_orgao AS T WHERE id_cadastro_orgao NOT IN ( SELECT T2.fk_id_cadastro_orgao FROM tb_transplante AS T2 WHERE T.id_cadastro_orgao = T2.fk_id_cadastro_orgao )\n" +
            "order by dt_cadastro asc", nativeQuery = true)
    List<CadastroOrgao> getOrgaoCadastradoNaoUsados();
}

