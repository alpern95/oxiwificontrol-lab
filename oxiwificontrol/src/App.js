import React from 'react';
import { Admin } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import {UserList} from "./components/users";

//connect the data provider to the REST endpoint
const dataProvider = jsonServerProvider('http://192.168.1.32');

function App() {
 return (
     <Admin dataProvider={dataProvider} >
         <Resource name="users" list={UserList}/>
     </Admin>
 );
}

export default App;
