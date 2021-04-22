import React, 
{
    useCallback,
    useState
} from 'react';

import {
    Button,
    SafeAreaView,
    StyleSheet,
    Text,
    TouchableOpacity,
    View
} from 'react-native';

import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { faCamera as fasCamera } from '@fortawesome/free-solid-svg-icons';

import InputUsername from '../components/Input/InputUsername';
import InputSenha from '../components/Input/InputSenha';

import { read } from '../DB';

const Login = (props) => {

    let [username, setUsername] = useState('');
    let [password, setPassword] = useState('');

    const handleLogin = () => {
        if ( username.trim().length > 0 && password.length > 0 ) {
            read(username, (errors, data) => {
                if ( errors ) {
                    alert('Não foi possível buscar no banco de dados!');
                } else if ( data === null ) {
                    alert('Usuário não encontrado!');
                } else {
                    const json = JSON.parse(data);
                    if ( json.password === password ) {
                        props.navigation.navigate('Main');
                    } else {
                        alert('Senha inválida!');
                    }
                }
            });

            return;
        }

        alert('Informe os campos corretamente!');
    }

    const handleCadastro = useCallback(() => {
        props.navigation.navigate('Cadastro');
    }, []);

    return (
        <SafeAreaView style={ styles.container }>
            <View style={ styles.logotipo }>
                <FontAwesomeIcon
                    color="#ed145b" 
                    icon={ fasCamera }
                    size={ 128 } />
                <Text style={ styles.logotipoTexto }>FIAPGram 2</Text>
            </View>
            

            <InputUsername
                onChangeText={ (txt) => setUsername(txt) }
                value={ username } />

            <InputSenha 
                onChangeText={ (txt) => setPassword(txt) }
                value={ password } />

            <Button 
                color="#ed145b" 
                onPress={ () => handleLogin() }
                title="Login" />

            <TouchableOpacity 
                onPress={ useCallback(() => handleCadastro(), []) }
                style={ styles.btnCadastro }>
                <Text>Cadastre-se</Text>
            </TouchableOpacity>
        </SafeAreaView>
    );
};

export default Login;

const styles = StyleSheet.create({
    btnCadastro : {
        alignItems : 'center',
        marginTop : 8
    },
    container : {
        flex : 1,
        justifyContent : 'center',
        padding : 16
    },
    logotipo : {
        alignItems : 'center'
    },
    logotipoTexto : {
        fontSize : 18,
        fontWeight : 'bold'
    }
});