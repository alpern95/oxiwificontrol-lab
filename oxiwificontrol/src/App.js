//  in src/App.js
import * as React from "react";
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import { fetchUtils } from 'react-admin';
import authProvider from './authProvider';
import Dashboard from './Dashboard';
import { BorneList, BorneEdit, BorneCreate} from './bornes';

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
    <Admin dashboard={Dashboard} authProvider={authProvider} dataProvider={dataProvider}>
    <Resource name="borne" list={BorneList} edit={BorneEdit} create={BorneCreate} />
    </Admin>
)

export default App;
