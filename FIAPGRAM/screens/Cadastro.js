import React,
{
    useState
} from 'react';

import {
    Button,
    SafeAreaView,
    StyleSheet,
    Text,
    View
} from 'react-native';

import InputEmail from '../components/Input/InputEmail';
import InputUsername from '../components/Input/InputUsername';
import InputSenha from '../components/Input/InputSenha';

import { insert, read } from '../DB';

const Cadastro = (props) => {

    let [username, setUsername] = useState('');
    let [email, setEmail] = useState('');
    let [password, setPassword] = useState('');
    let [confirmPassword, setConfirmPassword] = useState('');
    let [errors, setErrors] = useState([]);

    const validar = () => {
        const newErrors = [];

        if ( username.trim().length < 6 ) {
            newErrors.push('Usuário não pode ter menos que 06 caracteres!');
        }

        const regexEmail = /^[a-z0-9.]+@[a-z0-9]+\.[a-z]+\.([a-z]+)?$/i;
        if ( ! regexEmail.test(email) ) {
            newErrors.push('E-mail informado não é válido!');
        }

        const regexPassword = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{6,}$/;
        if ( !regexPassword.test(password) ) {
            newErrors.push('A senha precisa ter: maiúsculas, minúsculas e números!');
        }

        if ( password !== confirmPassword ) {
            newErrors.push('Senhas não são idênticas!');
        }

        setErrors( newErrors );

        return newErrors.length === 0 ;
    }

    const handleSalvar = () => {
        if ( validar() ) {
            read(username, (errors, data) => {
                if ( errors ) {
                    alert('Ocorreu algum erro ao buscar no banco de dados!');
                } else if ( data === null ) {
                    insert(username, { password, email }, (err) => {
                        if ( err ) {
                            alert('Erro ao inserir no banco de dados!');
                        } else {
                            props.navigation.pop();
                        }
                    })
                } else {
                    alert('Usuário já existente no banco de dados!');
                }
            });
        }
    }

    return (
        <SafeAreaView style={ styles.container }>
            <InputUsername 
                onChangeText={ (txt) => setUsername(txt) }
                value={ username } />

            <InputEmail 
                onChangeText={ (txt) => setEmail(txt) }
                value={ email } />

            <InputSenha
                onChangeText={ (txt) => setPassword(txt) }
                value={ password } />

            <InputSenha 
                confirm 
                onChangeText={ (txt) => setConfirmPassword(txt) }
                value={ confirmPassword } />

            <Button 
                color="#ed145b" 
                onPress={ () => handleSalvar() }
                title="Salvar" />

            { errors.length > 0 &&
                <View style={ styles.containerErrors }>
                    {errors.map( (erro, index) => (
                        <Text key={index} style={ styles.error }>{erro}</Text> 
                    ))}
                </View>
            }

        </SafeAreaView>
    );
};

export default Cadastro;

const styles = StyleSheet.create({
    container : {
        padding : 16
    },
    containerErrors : {
        backgroundColor : '#DDD',
        marginTop : 8,
        padding : 16
    },
    error : {
        marginBottom : 8
    }
});