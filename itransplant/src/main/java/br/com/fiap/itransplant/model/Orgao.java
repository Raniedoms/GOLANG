package br.com.fiap.itransplant.model;

import lombok.Getter;
import lombok.Setter;
import javax.persistence.*;

@Getter
@Setter
@Table(name = "tb_param_orgao")
@Entity
@SequenceGenerator(name="orgao",sequenceName = "SQ_TB_ORGAO", allocationSize = 1)
public class Orgao {

    @Id
    @Column(name = "id_orgao")
    @GeneratedValue(generator = "orgao", strategy = GenerationType.SEQUENCE)
    private Integer id;

    @Column(name = "nm_orgao", length = 100)
    private String nome;
}
