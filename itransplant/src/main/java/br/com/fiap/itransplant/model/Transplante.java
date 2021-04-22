package br.com.fiap.itransplant.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.Date;

@Getter
@Setter
@Table(name = "tb_transplante")
@Entity
@SequenceGenerator(name="transplante",sequenceName = "SQ_TB_TRANSPLANTE", allocationSize = 1)
public class Transplante {

    @Id
    @Column(name = "id_transplante")
    @GeneratedValue(generator = "transplante", strategy = GenerationType.SEQUENCE)
    private Integer id;

    @ManyToOne
    @JoinColumn(name = "fk_id_cadastro_orgao")
    private CadastroOrgao cadastroOrgao;

    @ManyToOne
    @JoinColumn(name = "fk_id_pessoa")
    private Pessoa pessoa;

    @Column(name = "transplante_efetivo", length = 1)
    private String transplanteEfetivo;

    @Column(name = "dt_transplante")
    @Temporal(TemporalType.TIMESTAMP)
    private Date dataTransplante;

}
