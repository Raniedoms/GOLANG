package br.com.fiap.itransplant.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.sql.Time;
import java.sql.Timestamp;
import java.time.LocalDate;
import java.util.Calendar;
import java.util.Date;
import java.util.List;

@Getter
@Setter
@Table(name = "tb_pessoa")
@Entity
@SequenceGenerator(name="pessoa",sequenceName = "SQ_TB_PESSOA", allocationSize = 1)
public class Pessoa {

    @Id
    @Column(name = "id_pessoa")
    @GeneratedValue(generator = "pessoa", strategy = GenerationType.SEQUENCE)
    private Integer id;

    @Column(name = "cd_cpf", length = 11)
    private String cpf;

    @ManyToOne
    @JoinColumn(name = "fk_cod_endereco")
    private Endereco endereco;

    @Column(name = "nm_pessoa", length = 100)
    private String nome;

    @Column(name = "dt_nascimento")
    @Temporal(TemporalType.DATE)
    private Date dataNascimento;

    @Column(name = "tp_sanguineo", length = 3)
    private String tipoSanguineo;

    @Column(name = "num_altura")
    private float altura;

    @Column(name = "ds_observacao", length = 300)
    private String observacao;

    @Column(name = "dt_entrada")
    @Temporal(TemporalType.TIMESTAMP)
    private Date dataEntrada;

    @OneToOne
    @JoinColumn(name = "fk_id_orgao")
    private Orgao orgao;

    @Column(name = "fl_ativo" , length = 1)
    private String ativo;
}
