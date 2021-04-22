package br.com.fiap.itransplant.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.Date;

@Getter
@Setter
@Table(name = "tb_cadt_orgao")
@Entity
@SequenceGenerator(name="cadastroOrgao",sequenceName = "SQ_TB_CADASTRO_ORGAO", allocationSize = 1)
public class CadastroOrgao {

    @Id
    @Column(name = "id_cadastro_orgao")
    @GeneratedValue(generator = "cadastroOrgao", strategy = GenerationType.SEQUENCE)
    private Integer id;

    @ManyToOne
    @JoinColumn(name = "fk_id_orgao")
    private Orgao orgao;

    @Column(name = "tp_sanguineo", length = 3)
    private String tipoSanguineo;

    @Column(name = "observacao_orgao", length = 100)
    private String observacao;

    @Column(name = "dt_cadastro")
    @Temporal(TemporalType.TIMESTAMP)
    private Date dataCadastro;

}
