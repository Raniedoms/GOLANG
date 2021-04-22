import React from 'react';

import {
    SafeAreaView
} from 'react-native';

import { getPosts } from '../services/PostService';

const Main = (props) => {

    getPosts().then((response) => {
        console.log(response.data);
    })

    return (
        <SafeAreaView>

        </SafeAreaView>
    );
};

export default Main;