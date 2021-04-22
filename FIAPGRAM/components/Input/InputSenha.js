import React from 'react';

import { Input } from 'react-native-elements';

import { FontAwesomeIcon } from '@fortawesome/react-native-fontawesome';
import { faLock as fasLock } from '@fortawesome/free-solid-svg-icons';

const InputSenha = (props) => {

    const placeholder = ( props.confirm ) 
                            ? 'Confirme a sua senha:'
                            : 'Digite a sua senha:'

    return (
        <Input 
            leftIcon={
                <FontAwesomeIcon 
                    color="#333"
                    icon={ fasLock }
                    size={ 24 }/>
            }
            placeholder={ placeholder }
            secureTextEntry
            {...props} />
    );
};

export default InputSenha;