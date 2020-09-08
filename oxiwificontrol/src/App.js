//  in src/App.js
import * as React from "react";
import { Admin, Login, Resource } from 'react-admin';
import simpleRestProvider from 'ra-data-simple-rest';
import { fetchUtils } from 'react-admin';
import authProvider from './authProvider';
import Dashboard from './Dashboard';
import { BorneList, BorneEdit, BorneCreate} from './bornes';
import { UserList, UserEdit, UserCreate} from'./users';
import { GroupeList} from './groupes';

//import MyLayout from './MyLayout';

const MyLoginPage = () => <Login backgroundImage="/background.jpg" />;

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

const dataProvider = simpleRestProvider('http://192.168.112.10:8081/api/v1', httpClient);
//const dataProviser = simpleRestProviser(REAC_APP_OXIWIFICONTROLADDR,httpClient);

const App = () => (
    <Admin loginPage={MyLoginPage} 
        dashboard={Dashboard}
        authProvider={authProvider}
        dataProvider={dataProvider}
    >
  {permissions => [
    permissions === 'admin'
    ?  <Resource 
            name="borne"
            list={BorneList}
            edit={BorneEdit}
            create={BorneCreate}
       />   
       : null,
    permissions === 'admin'
    ? <Resource
            name="users"
            list={UserList}
            edit={UserEdit}
            create={UserCreate}
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
