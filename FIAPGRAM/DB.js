import AsyncStorage from '@react-native-community/async-storage';

const insert = async (key, value, callback = null) => {
    try {
        const jsonValue = JSON.stringify(value);
        return await AsyncStorage.setItem(key, jsonValue, callback);
    } catch ( e ) {
        throw new Error('Não foi possível salvar no banco de dados!');
    }
}

const read = async (key, callback = null) => {
    try {
        return await AsyncStorage.getItem(key, callback);
    } catch ( e ) {
        throw new Error('Não foi possível ler do banco de dados!');
    }
}


export { insert, read };