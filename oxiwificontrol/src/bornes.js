import * as React from "react";
import { List, Edit, Create, TabbedForm,FormTab,TextInput, Datagrid, TextField, EditButton,PasswordInput } from 'react-admin';

export const BorneList = props => (
    <List {...props} >
        <Datagrid rowClick="edit">
            <TextField source="nom" sortable={false} />
            <TextField source="description" sortable={false} /> 
            <TextField source="groupe" sortable={false} />
            <TextField source="etat" sortable={false} />
            <TextField source="lastrefresh" sortable={false} />
            <EditButton />
        </Datagrid>
    </List>
);

export const BorneEdit = props => (
    <Edit {...props} >
    <TabbedForm>
      <FormTab
          label="base"
      >
      <TextInput source="nom" />
      <TextInput source="description" />
      </FormTab>
    
        <FormTab 
        label="détail" 
        >
        <TextInput source="device" />
        <TextInput source="adresse" />
        <TextInput source="groupe" />
        <TextInput source="modele" />
        <TextInput source="username" />
        <PasswordInput source="password" />
        <PasswordInput source="enablepassword" />
        <TextInput source="interface" />
        <TextInput source="etat" />
        <TextInput source="lastrefresh" />
        <TextInput source="utilisateurs" />
      </FormTab>
    </TabbedForm>
  </Edit>
);

export const BorneCreate = props => (
    <Create {...props} >
    <TabbedForm>
      <FormTab
          label="base"
      >
      <TextInput source="nom" />
      <TextInput source="description" />
      </FormTab>
        <FormTab
        label="details"
        >
        <TextInput source="device" />
        <TextInput source="adresse" />
        <TextInput source="groupe" />
        <TextInput source="modele" />
        <TextInput source="username" />
        <PasswordInput source="password" />
        <PasswordInput source="enablepassword" />
        <TextInput source="interface" />
        <TextInput source="etat" />
        <TextInput source="lastrefresh" />
        <TextInput source="utilisateurs" />
      </FormTab>
    </TabbedForm>
  </Create>
);
