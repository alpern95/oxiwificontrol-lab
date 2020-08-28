//  in src/App.js
import * as React from "react";
import { Admin, Resource } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import { fetchUtils } from 'react-admin';
import authProvider from './authProvider';
import Dashboard from './Dashboard';
import { BorneList, BorneEdit, BorneCreate} from './bornes';
import { GroupeList} from './groupes';

//import MyLayout from './MyLayout';

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' });
    }
    //options.headers.set('X-Total-Count');

    const token = localStorage.getItem('token');
    //const role = localStorage.getItem('permissions');
    options.headers.set('Authorization', `Bearer ${token}`);
    return fetchUtils.fetchJson(url, options);
};


const dataProvider = simpleRestProvider('http://192.168.112.10:3000/api/v1', httpClient);

const App = () => (
    <Admin 
        dashboard={Dashboard}
        authProvider={authProvider}
        dataProvider={dataProvider}
    >
  {permissions => [
    permissions === 'admin'
    ? <Resource 
            name="borne"
            list={BorneList}
            edit={BorneEdit}
            create={BorneCreate}
      />
      : null,
    permissions !== 'admin'
    ? <Resource
            name="groupe"
            list={GroupeList}
      />
      : null,
  ]}
    </Admin>
)

export default App;
