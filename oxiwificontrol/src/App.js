//  in src/App.js
import * as React from "react";
import { Admin, Resource } from 'react-admin';

import jsonServerProvider from 'ra-data-json-server';

import { fetchUtils } from 'react-admin';
import authProvider from './authProvider';
//import authProvider from "./providers/authProvider";
//import {
//	   FirebaseAuthProvider,
//	   } from 'react-admin-firebase';

//import { ListGuesser } from 'react-admin';
import { BorneList,BorneCreate } from './bornes';
import { UserList,UserCreate } from './users';
import { GroupeList ,GroupeCreate} from './groupes';

const fetchJson = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    // add your own headers here
    //options.headers.set('X-Custom-Header', 'users');
    return fetchUtils.fetchJson(url, options);
}

const dataProvider = jsonServerProvider('http://192.168.1.32:3000/api/v1',fetchJson);
const App = () => (
    <Admin authProvider={authProvider} dataProvider={dataProvider}>
    <Resource name="user" list={UserList} create={UserCreate} />
    <Resource name="borne" list={BorneList} create={BorneCreate} />
    <Resource name="groupe" list={GroupeList} create={GroupeCreate} />
    </Admin>
)
export default App;