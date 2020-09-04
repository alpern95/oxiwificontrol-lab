import * as React from "react";
import { List, Edit, Create, SimpleForm, TextInput, Datagrid, TextField, EditButton,PasswordInput } from 'react-admin';

export const UserList = props => (
    <List {...props} >
        <Datagrid rowClick="edit">
            <TextField source="username" sortable={false} />
            <TextField source="password" sortable={false} /> 
            <TextField source="email" sortable={false} />
            <TextField source="role" sortable={false} />
            <TextField source="token" sortable={false} />
            <EditButton />
        </Datagrid>
    </List>
);

export const UserEdit = props => (
    <Edit {...props} >
      <SimpleForm >
        <TextInput source="username" />
        <PasswordInput source="password" />
        <TextInput source="email" />
        <TextInput source="role" />
        </SimpleForm>
    </Edit>
);

export const UserCreate = props => (
    <Create {...props} >
        <SimpleForm >
        <TextInput source="username" />
        <PasswordInput source="password" />
        <TextInput source="email" />
        <TextInput source="role" />
        </SimpleForm>
    </Create>
);
