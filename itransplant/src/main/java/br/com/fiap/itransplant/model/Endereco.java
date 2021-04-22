package br.com.fiap.itransplant.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.List;

@Getter
@Setter
@Table(name = "tb_endereco")
@Entity
@SequenceGenerator(name="endereco",sequenceName = "SQ_TB_ENDERECO", allocationSize = 1)
public class Endereco {

    @Id
    @Column(name = "cod_endereco")
    @GeneratedValue(generator = "endereco", strategy = GenerationType.SEQUENCE)
    private Integer codigoEndereco;

    @Column(name = "nm_rua", length = 50)
    private String rua;

    @Column(name = "nm_cep", length = 15)
    private String cep;

    @Column(name = "numero")
    private Integer numero;

    @Column(name = "nm_estado", length = 50)
    private String estado;

    @Column(name = "complemento", length = 50)
    private String complemento;

}
